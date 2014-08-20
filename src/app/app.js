var App = Ember.Application.create({
  LOG_TRANSITIONS: true,
  app_settings_token: '',
  ready: function() {
    Ember.debug("App ready!");
  }
});

//Ember.TEMPLATES['blog'] = Em.Handlebars.compile('<div>{{outlet}}</div>');


App.deferReadiness();
// Wait for all the javascript files to load.
$(document).ready(function(){
 
    // Set everything else up.
    $.getJSON("/api/1/gettoken",
        function(response) {
          if(response.token) {
            App.set('app_settings_token', response.token);         
          } else {
             App.set('app_settings_token', '');         
          }
           // Will start everything up.
          App.advanceReadiness();
        }
      );
  
});
