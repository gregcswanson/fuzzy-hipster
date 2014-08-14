App.ItemsController = Ember.ArrayController.extend({
  actions: {
    toggleIsDone: function(item) {
      item.set('isDone', !item.get('isDone'));
      item.save();
    }
  }
});