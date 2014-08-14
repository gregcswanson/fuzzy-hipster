App.Settings = Ember.Object.extend({ 
  token : '', 
  username: ''
});

App.Settings.reopenClass({
  find: function() {
    return $.getJSON("/api/1/gettoken",
        function(response) {
          $.ajaxSetup({
            beforeSend: function(xhr) {
              xhr.setRequestHeader('Authorization-Token', response.token);
            }
          });
          App.Token = App.Settings.create(response);
          return App.Token;
        }
      ).fail(function(){ location.reload(); });
  }
});

function setHeader(xhr) {
   xhr.setRequestHeader('Authorization-Token', App.get('app_settings_token'));
 }
