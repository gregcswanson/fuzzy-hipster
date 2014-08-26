App.DayIndexRoute = Ember.Route.extend({
	model: function(params){
    var day = this.modelFor("day");
    console.log(day.id);
    return App.DayItem.findForDay(day.id);
	}
});