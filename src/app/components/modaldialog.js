App.ModalDialogComponent = Ember.Component.extend({
  didInsertElement: function () {
        var self = this;
    $('#myModal').modal({ keyboard: false, backdrop:'static' });
    },
  actions: {
    close: function() {
      return this.sendAction();
    }
  }
});