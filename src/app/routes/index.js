App.IndexRoute = Ember.Route.extend({
  model: function() {
    return App.Settings.find();
    //return this.store.findAll('list');
  }
});