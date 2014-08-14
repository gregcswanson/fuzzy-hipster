App.ProjectsIndexRoute = Ember.Route.extend({
	model: function(){
		return App.Project.findAll();
	}
});