App.ModalController = Ember.ObjectController.extend({
  title: 'Modal Test',
  newText: 'new',
  newStatus: 'NOTE',
  dependency: function() {
    this.set('newText', this.get('model.Text'));
    this.set('newStatus', this.get('model.Status'));

  }.observes('model'),
  isNote: function() {
    return this.get('newStatus') == 'NOTE';
  }.property('newStatus')
  , isOpen: function() {
    return this.get('newStatus') == 'OPEN';
  }.property('newStatus')
  , isDone: function() {
    return this.get('newStatus') == 'DONE';
  }.property('newStatus')
  , isDiscarded: function() {
    return this.get('newStatus') == 'DISCARDED';
  }.property('newStatus'),
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
    },
    setDone: function() {  
        this.set('newStatus', 'DONE');
    },
    setOpen: function() {  
        this.set('newStatus', 'OPEN');
    },
    setNote: function() {  
        this.set('newStatus', 'NOTE');
    },
    setDiscarded: function() {  
        this.set('newStatus', 'DISCARDED');
    }                                            
  }
});