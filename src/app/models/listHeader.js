App.ListHeader = Ember.Object.extend({
    title: '',
    description: ''
});

App.ListHeader.reopenClass({

  findAll: function() {
    return $.getJSON("/api/1/lists").then(
      function(response) {
        return response.lists.map(function (child) {
          return App.ListHeader.create(child.data);
        });
      }
    );
  }

});