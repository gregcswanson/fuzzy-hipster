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