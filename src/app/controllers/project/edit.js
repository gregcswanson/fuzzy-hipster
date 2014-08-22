App.ProjectEditController = Ember.ObjectController.extend({
  newTitle: '',
  dependency: function() {
    this.set('newTitle', this.get('model.Title'));

  }.observes('model'),
  actions: {
    close: function() {
      return this.send('closeModal');
    },
    save: function() {
      var controller = this;
      var model = this.get('model');
      model.Title = this.get('newTitle');
      App.Project.update(model).then(function(data){
        controller.send('closeModal');
        controller.send("itemsChanged");
      });
    }
  }
});