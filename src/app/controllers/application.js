App.ApplicationController = Ember.Controller.extend( {
  isAuthenticated: false,  // assume we're not until proven wrong
  init: function () {
    var appController = this;  
    if (App.get('app_settings_token') !== '') {
      appController.set('isAuthenticated', true);            
    } else {
      appController.set('isAuthenticated', false);  
    }
  }
});

// http://stackoverflow.com/questions/22122570/ember-js-network-connectivity-check-at-startup-and-listener
