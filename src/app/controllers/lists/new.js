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