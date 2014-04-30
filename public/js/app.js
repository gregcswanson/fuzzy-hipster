var App = Ember.Application.create({
  LOG_TRANSITIONS: true
});

App.Router.map(function() {
  this.resource('lists', function() {
    this.route('new');
    this.resource('list', { path: 'list/:list_id' });
  });
  this.route('about', { path: '/about' });
});

App.IndexController = Ember.ArrayController.extend({
  listsCount: Ember.computed.alias('length')
});

App.ListsController = Ember.ArrayController.extend({
  sortProperties: ['title'],
  listsCount: Ember.computed.alias('length')
});

App.ListsNewController = Ember.ObjectController.extend({
  actions: {
    createList: function() { 
      var list = this.store.createRecord('list', { 
         title: this.get('title'), 
         description: this.get('description')
       });
      list.save();
      this.transitionToRoute('lists');
    }
  } 
});

App.IndexRoute = Ember.Route.extend({
  model: function() {
    return this.store.findAll('list');
  }
});

App.ListsRoute = Ember.Route.extend({
  model: function() {
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

App.ApplicationAdapter = DS.FixtureAdapter.extend();

App.List = DS.Model.extend({
  title: DS.attr('string'),
  current: DS.attr('number'),
  total: DS.attr('number'),
  description: DS.attr('string'),
  items: DS.hasMany('item', { async: true })
});

App.Item = DS.Model.extend({
  name: DS.attr('string'),
  description: DS.attr('string'),
  isDone: DS.attr('boolean')
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
    description: 'Flint is a hard, sedimentary cryptocrystalline form of the mineral quartz, categorized as a variety of chert.',
    isDone: true,
    list: 200
  },
  {
    id: 2,
    title: 'Kindling',
    description: 'Easily combustible small sticks or twigs used for starting a fire.',
    isDone: false,
    list: 200
  },
  {
    id: 3,
    title: 'Matches',
    description: 'One end is coated with a material that can be ignited by frictional heat generated by striking the match against a suitable surface.',
    isDone: true,
    list: 201
  },
  {
    id: 4,
    title: 'Bow Drill',
    description: 'The bow drill is an ancient tool. While it was usually used to make fire, it was also used for primitive woodworking and dentistry.',
    isDone: false,
    list: 201
  },
  {
    id: 5,
    title: 'Tinder',
    description: 'Tinder is easily combustible material used to ignite fires by rudimentary methods.',
    isDone: true,
    list: 201
  },
  {
    id: 6,
    title: 'Birch Bark Shaving',
    description: 'Fresh and easily combustable',
    isDone: true,
    list: 201
  }
];