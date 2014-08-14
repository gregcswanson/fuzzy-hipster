var App = Ember.Application.create({
  LOG_TRANSITIONS: true,
  app_settings_token: '',
  ready: function() {
    Ember.debug("App ready!");
  }
});
