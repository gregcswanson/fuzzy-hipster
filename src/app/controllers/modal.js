App.ModalController = Ember.ObjectController.extend({
  title: 'Modal Test',
  actions: {
    close: function() {
      return this.send('closeModal');
    },
    refresh: function() {
      this.send('closeModal');
      this.send("itemsChanged");
    },
    save: function() {
      alert('save');
    }
  }
});