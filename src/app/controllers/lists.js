App.ListsController = Ember.ArrayController.extend({
  sortProperties: ['title'],
  listsCount: Ember.computed.alias('length')
});