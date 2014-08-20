App.ProjectRoute = Ember.Route.extend({
	model: function(params){
    console.log(params.project_id);
    return App.Project.find(params.project_id);
	}
});