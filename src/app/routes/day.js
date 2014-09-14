App.DayRoute = Ember.Route.extend({
	model: function(params){
    // convert the parameter to a new date
    if (params.day_id == '0'){
      var d = new Date();
      var curr_date = d.getDate();
      var curr_month = d.getMonth() + 1;
      var curr_year = d.getFullYear();
   
      var day_id = curr_year.toString();
      if(curr_month < 10) {
        day_id = day_id + '0';
      }
      day_id = day_id + curr_month;
      if (curr_date < 10) {
        day_id = day_id + '0';
      }
      day_id = day_id + curr_date;
      this.transitionTo('day', day_id);
    }
    
    return { id: params.day_id };
	}
});