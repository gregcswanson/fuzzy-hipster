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