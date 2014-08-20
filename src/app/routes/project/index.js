App.ProjectIndexRoute = Ember.Route.extend({
	model: function(params){
    return this.modelFor("project"); //get the parent model
	}
});