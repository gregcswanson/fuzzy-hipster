App.ListsRoute = Ember.Route.extend({
  model: function() {
    //return App.ListHeader.findAll();
    return this.store.findAll('list');
  }
});