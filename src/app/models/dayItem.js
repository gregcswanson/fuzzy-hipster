App.DayItem = Ember.Object.extend({
    Text: ''
});

App.DayItem.reopenClass({
  findForDay: function(day) {
       return Em.$.ajax({
              url: '/api/1/day/' + day,
              type: 'GET',
              dataType: 'json',
              beforeSend: setHeader
       }).then(function(response) { 
                 var items = [];
               response.dayitems.forEach(function(data){
                var model = App.DayItem.create(data); 
                items.addObject(model);
               })
                return items;
              }
            );
  	},
  save: function(dayitem) {
      var data =  JSON.stringify({ "dayitem" : dayitem });
      return Em.$.ajax({
        url: '/api/1/dayitem',
        type: 'POST',
        data: data,
        dataType: 'json',
        beforeSend: setHeader
      }).fail(function( jqXHR, textStatus, errorThrown ){
          alert(jqXHR.responseText); 
      });
  }
});