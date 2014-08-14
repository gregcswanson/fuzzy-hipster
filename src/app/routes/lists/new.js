App.ListsNewRoute = Ember.Route.extend({
  model: function() {
    return {'title': 'new title', 'description': ''};
  }
});