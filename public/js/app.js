var App = Ember.Application.create({
  LOG_TRANSITIONS: true
});

App.Router.map(function() {
  this.resource('lists', function() {
    this.route('new');
  });
  this.resource('list', { path: 'list/:list_id' }, function() {
    this.route("edit", { path: "/edit" });
    this.route("delete");
  });
  this.route('about', { path: '/about' });
});

App.IndexController = Ember.ObjectController.extend({
  //listsCount: Ember.computed.alias('length'),
  actions: {
    getToken: function() { 
      //$.getJSON("/api/1/gettoken", // ).then(
      //  function(response) {
      //    App.Settings.token = response.token;
          //alert(response.token);
          // use the token
           $.ajax({
              url: '/api/1/checktoken',
              type: 'GET',
              dataType: 'json',
              success: function(response) { alert(response.token); },
              error: function() { alert('no'); }//,
              //beforeSend: function (request)
              //{
              //    request.setRequestHeader("Authorization-Token", App.Settings.token); //response.token);
             // }
            });
      //  }
      //).fail(function(){ location.reload(); });
    }
  } 
});

App.ListsController = Ember.ArrayController.extend({
  sortProperties: ['title'],
  listsCount: Ember.computed.alias('length')
});

App.ItemsController = Ember.ArrayController.extend({
  actions: {
    toggleIsDone: function(item) {
      item.set('isDone', !item.get('isDone'));
      item.save();
    }
  }
});

App.ItemsEditController = Ember.ArrayController.extend({
  actions: {
    removeItem: function(item){
      if(item.get('isNew')) {
        item.deleteRecord();
        this.get('model').removeObject(item);
      } else {
        item.destroyRecord();
        this.get('model').removeObject(item);
      }
    }
  }
});

App.ListEditRoute = Ember.Route.extend({
  deactivate: function() {
    var model = this.controller.content;
    if (model.get('isDirty') && !model.get('isSaving')) {
      model.save();
    }
    this.controller.get('model.items').forEach(function(item){
      if((item.get('isNew') ||  item.get('isDirty')) && !item.get('isSaving')) {
        item.save();
      }
    });
  }
});

App.ListsNewController = Ember.ObjectController.extend({
  actions: {
    createList: function() { 
      var list = this.store.createRecord('list', { 
         title: this.get('title'), 
         description: this.get('description')
       });
      var self = this;
      list.save().then(function() {
        self.transitionToRoute('list.edit', list);
      });
    }
  } 
});

App.ListEditController = Ember.ObjectController.extend({
  actions: {
    save: function() { 
      //var model = this.get('model');
      //model.save();
      this.transitionToRoute('list.index');
    },
    addItem: function() {
      console.log('addItem clicked');
      var newItem = this.store.createRecord('item', { 
        name: '', 
        list: this.get('model'),
        isDone: false
       });
      this.get('model.items').addObject(newItem);
    }
  } 
});

App.ListDeleteController = Ember.ObjectController.extend({
  actions: {
    deleteList: function() { 
      var controller = this;
      var model = this.get('model').destroyRecord().then(function() {
        controller.transitionToRoute('lists');
       });
    }
  } 
});

App.IndexRoute = Ember.Route.extend({
  model: function() {
    return App.Settings.find();
    //return this.store.findAll('list');
  }
});

App.ListsRoute = Ember.Route.extend({
  model: function() {
    //return App.ListHeader.findAll();
    return this.store.findAll('list');
  }
});

App.ListsNewRoute = Ember.Route.extend({
  model: function() {
    return {'title': 'new title', 'description': ''};
  }
});

App.ListsIndexRoute = Ember.Route.extend({
  model: function(){
    return this.store.findAll('list');
  }
});

// Models
App.ListHeader = Ember.Object.extend({
    title: '',
    description: ''
});

App.ListHeader.reopenClass({

  findAll: function() {
    return $.getJSON("/api/1/lists").then(
      function(response) {
        return response.lists.map(function (child) {
          return App.ListHeader.create(child.data);
        });
      }
    );
  }

});

// Settings
App.Settings = Ember.Object.extend({ 
  token : '', 
  username: ''
});

App.Settings.reopenClass({
  find: function() {
    return $.getJSON("/api/1/gettoken",
        function(response) {
          $.ajaxSetup({
            beforeSend: function(xhr) {
              xhr.setRequestHeader('Authorization-Token', response.token);
            }
          });
          return App.Settings.create(response);
        }
      ).fail(function(){ location.reload(); });
  }
});

// Ember data

//App.ApplicationAdapter = DS.FixtureAdapter.extend();
DS.RESTAdapter.reopen({
  namespace: 'api/1'
});

App.ApplicationSerializer = DS.RESTSerializer.extend({
  serializeHasMany: function(record, json, relationship) {
    var key = relationship.key;
    var payloadKey = this.keyForRelationship ? this.keyForRelationship(key, 'hasMany') : key;
    var relationshipType = DS.RelationshipChange.determineRelationshipType(record.constructor, relationship);

    if (['manyToNone', 'manyToMany', 'manyToOne'].contains(relationshipType)) {
      json[payloadKey] = record.get(key).mapBy('id');
    }
  }
});

App.List = DS.Model.extend({
  title: DS.attr('string'),
  current: DS.attr('number'),
  total: DS.attr('number'),
  description: DS.attr('string'),
  items: DS.hasMany('item', { async: true })
});

App.Item = DS.Model.extend({
  name: DS.attr('string'),
  isDone: DS.attr('boolean'),
  list: DS.belongsTo('list')
});

App.List.FIXTURES = [
  {
    id: 200,
    title: 'List 1',
    description: 'The first list',
    current: 1,
    total: 2,
    items: [1, 2]
  },
  {
    id: 201,
    title: 'List 2',
    description: 'The second list',
    current: 3,
    total: 4,
    items: [3, 4, 5, 6]
  }
];

App.Item.FIXTURES = [
 {  
    id: 1,
    name: 'Flint',
    isDone: true,
    list: 200
  },
  {
    id: 2,
    name: 'Kindling',
    isDone: false,
    list: 200
  },
  {
    id: 3,
    name: 'Matches',
    isDone: true,
    list: 201
  },
  {
    id: 4,
    name: 'Bow Drill',
    isDone: false,
    list: 201
  },
  {
    id: 5,
    name: 'Tinder',
    isDone: true,
    list: 201
  },
  {
    id: 6,
    name: 'Birch Bark Shaving',
    isDone: true,
    list: 201
  }
];