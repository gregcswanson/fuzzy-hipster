App.DaysDayItemController = Ember.ObjectController.extend({
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
      var model = this.get('model');
      //App.Project.updateline(model).then( function(data){
      //  console.log('saved');
      //});
    } else if(this.get('Status') == 'DONE') {
      this.set('Status', 'DISCARDED');   
      var model = this.get('model');
      //App.Project.updateline(model).then( function(data){
      //  console.log('saved');
      //});
    } else {
      this.set('Status', 'OPEN');
      var model = this.get('model');
      //App.Project.updateline(model).then( function(data){
      //  console.log('saved');
      //});
    }
    }
  }
});