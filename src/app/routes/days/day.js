App.DaysDayRoute = Ember.Route.extend({
	model: function(params){
    console.log(params.day_id);
    return App.DayItem.findForDay(params.day_id);
	}
});