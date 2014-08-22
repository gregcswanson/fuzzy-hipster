App.ModalController = Ember.ObjectController.extend({
  title: 'Modal Test',
  newText: 'new',
  newStatus: 'NOTE',
  dependency: function() {
    this.set('newText', this.get('model.Text'));
    this.set('newStatus', this.get('model.Status'));

  }.observes('model'),
  actions: {
    close: function() {
      return this.send('closeModal');
    },
    refresh: function() {
      this.send('closeModal');
      this.send("itemsChanged");
    },
    save: function() {
      var controller = this;
      var model = this.get('model');
      model.Text = this.get('newText');
      model.Status = this.get('newStatus');
      App.Project.updateline(model).then( function(data){
        controller.send('closeModal');
        controller.send("itemsChanged");
      });
    },
    delete: function() {
      var controller = this;
      var model = this.get('model');
      App.Project.deleteline(model).then( function(data){
        controller.send('closeModal');
        controller.send("itemsChanged");
      });
    }
  }
});