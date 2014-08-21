App.ModalController = Ember.ObjectController.extend({
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