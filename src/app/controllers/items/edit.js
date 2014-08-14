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