App.DayRoute = Ember.Route.extend({
	model: function(params){
    return { id: params.day_id };
	}
});