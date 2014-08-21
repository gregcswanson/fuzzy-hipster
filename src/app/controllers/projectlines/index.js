App.ProjectlinesIndexController = Ember.ArrayController.extend({
  needs: ["ProjectIndex"],
  newItemText : '',
  newItemStatus: 'OPEN',
  statuses: ["NOTE", "OPEN"],
  searchChangeStatus: function() {
        if (this.get('newItemText') == '/') {
          this.set('newItemStatus','NOTE');
          this.set('newItemText', '');
        }
    if (this.get('newItemText') == '.') {
          this.set('newItemStatus','OPEN');
          this.set('newItemText', '');
        }
    }.observes("newItemText"),
  isNote: function() {
    return this.get('newItemStatus') == 'NOTE';
  }.property('newItemStatus'),
  isOpen: function() {
    return this.get('newItemStatus') == 'OPEN';
  }.property('newItemStatus'),
  actions: {
    addItem: function() {
      var newItem = { Status: this.get('newItemStatus'), Text: this.get('newItemText'), ProjectID:  this.parentController.get('ID')};
      // check for blank - disable save button
      var controller = this;
      App.Project.saveline(newItem).then( function(data){
        console.log(data.line);
        controller.get('model').addObject(data.line);
        controller.set('newItemText', '');
        controller.set('newItemStatus', 'OPEN');
      });
    },
    toggleStatus: function() {  
      if(this.get('newItemStatus') == 'OPEN') {
        this.set('newItemStatus', 'NOTE');
      } else if(this.get('newItemStatus') == 'NOTE') {
        this.set('newItemStatus', 'OPEN');  
      }
    }
  } 
  
});