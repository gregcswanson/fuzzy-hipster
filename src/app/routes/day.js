App.DayRoute = Ember.Route.extend({
	model: function(params){
    // convert the parameter to a new date
    
    return { id: params.day_id };
	}
});