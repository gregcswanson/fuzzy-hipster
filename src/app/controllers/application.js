App.ApplicationController = Ember.Controller.extend( {
  isAuthenticated: false,  // assume we're not until proven wrong
  token: '',
  init: function () {
    var appController = this;    
    $.getJSON("/api/1/gettoken",
        function(response) {
          if(response.token) {
            appController.set('token', response.token);
            App.set('app_settings_token', response.token);
            appController.set('isAuthenticated', true);            
          } else {
            appController.set('isAuthenticated', false);  
          }
        }
      ).fail(function(){ location.reload(); });
  }
});

// http://stackoverflow.com/questions/22122570/ember-js-network-connectivity-check-at-startup-and-listener
