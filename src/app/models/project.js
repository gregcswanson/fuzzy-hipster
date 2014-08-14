App.Project = Ember.Object.extend({
    title: ''
});

App.Project.reopenClass({
	findAll: function() {
      return $.ajax({
              url: '/api/1/projects',
              type: 'GET',
              dataType: 'json',
              success: function(response) { 
                console.log(response);
                //return response.projects.map(function (child) {
                //  return App.Project.create(child.data);
                //});
              },
              error: function() { alert('no'); },
              beforeSend: setHeader
            });
  	}
});