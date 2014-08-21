App.ProjectlinesItemController = Ember.ObjectController.extend({
  isNote: function() {
    return this.get('Status') == 'NOTE';
  }.property('Status')
  , isOpen: function() {
    return this.get('Status') == 'OPEN';
  }.property('Status')
  , isDone: function() {
    return this.get('Status') == 'DONE';
  }.property('Status')
  , isDiscarded: function() {
    return this.get('Status') == 'DISCARDED';
  }.property('Status')
  , actions: {
   toggleStatus: function() {  
    if(this.get('Status') == 'OPEN') {
      this.set('Status', 'DONE');
    } else if(this.get('Status') == 'DONE') {
      this.set('Status', 'DISCARDED');   
    } else {
      this.set('Status', 'OPEN');
    }
    }
  }
});