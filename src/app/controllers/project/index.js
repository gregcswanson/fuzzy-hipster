App.ProjectIndexController = Ember.ObjectController.extend({
  
  statuses: ["NOTE", "OPEN", "DONE"],
  actions: {
    save: function() { 
      this.transitionToRoute('projects');
    },
    addItem: function() {
      var newItem = { Status: this.get('newItemStatus'), Text: this.get('newItemText'), ProjectID: this.get('ID')};
      // check for blank - disable save button
      console.log(newItem);
      var controller = this;
      //App.Project.saveline(newItem).then( function(data){
        this.get('controllers.projectlinesindex').push(newItem);
        //alert('saved');
      //});
    },
    saveItem: function() { 
      console.log('saveItem clicked');
    }
    
    // To do - watch new item first characters for "n " for note, "c " for open checkbox, "x " for done checkbox, "r " for running checkbox
  } 
});