App.ProjectlinesIndexController = Ember.ArrayController.extend({
  needs: ["ProjectIndex"],
  newItemText : '',
  newItemStatus: 'NOTE',
  statuses: ["NOTE", "OPEN", "DONE"],
  actions: {
    addItem: function() {
      var newItem = { Status: this.get('newItemStatus'), Text: this.get('newItemText'), ProjectID:  this.parentController.get('ID')};
      // check for blank - disable save button
      var controller = this;
      App.Project.saveline(newItem).then( function(data){
        console.log(data.line);
        controller.get('model').addObject(data.line);
        controller.set('newItemText', '');
        controller.set('newItemStatus', 'NOTE');
      });
    }
    
    // To do - watch new item first characters for "n " for note, "c " for open checkbox, "x " for done checkbox, "r " for running checkbox
  } 
  
});