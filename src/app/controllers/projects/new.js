App.ProjectsNewController = Ember.ObjectController.extend({
	actions: {
    	createProject: function() { 
			
        var newProject = { "title": this.get('title')};
      var self = this;
        console.log(newProject);
        App.Project.save(newProject).then(function(data) {
          console.log('saved');
          console.log(data);
          self.transitionToRoute('projects');
        });
    	}
  	} 
});