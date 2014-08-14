App.IndexController = Ember.ObjectController.extend({
  actions: {
    getToken: function() { 
           $.ajax({
              url: '/api/1/checktoken',
              type: 'GET',
              dataType: 'json',
              success: function(response) { alert(response.token); },
              error: function() { alert('no'); }
            });
    }
  } 
});