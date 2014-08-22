App.Project = Ember.Object.extend({
    Title: ''
});

App.Project.reopenClass({
  findAll: function() {
       return Em.$.ajax({
              url: '/api/1/projects',
              type: 'GET',
              dataType: 'json',
              beforeSend: setHeader
       }).then(function(response) { 
                console.log(response);
                var d = response.projects.map(function (child) {
                  return App.Project.create(child.data);
                });
                console.log(d);
                return response.projects;
              }
            );
  	},
  find: function(id) {
       return Em.$.ajax({
              url: '/api/1/projects/' + id,
              type: 'GET',
              dataType: 'json',
              beforeSend: setHeader
       }).then(function(response) { 
                console.log(response);
                var d = App.Project.create(response.project);
                console.log(d);
                return response.project;
              }
            );
  	},
  save: function(project) {
      var data =  JSON.stringify({ "project" : project });
      return Em.$.ajax({
        url: '/api/1/projects',
        type: 'POST',
        data: data,
        dataType: 'json',
        beforeSend: setHeader
      }).fail(function( jqXHR, textStatus, errorThrown ){
          alert(jqXHR.responseText); 
      });
  },
  update: function(project) {
      var data =  JSON.stringify({ "project" : project });
      return Em.$.ajax({
        url: '/api/1/projects/' + project.ID,
        type: 'PUT',
        data: data,
        dataType: 'json',
        beforeSend: setHeader
      }).fail(function( jqXHR, textStatus, errorThrown ){
          alert(jqXHR.responseText); 
      });
  },
  saveline: function(line) {
      var data =  JSON.stringify({ "line" : line });
      return Em.$.ajax({
        url: '/api/1/projects/' + line.ProjectID + '/lines',
        type: 'POST',
        data: data,
        dataType: 'json',
        beforeSend: setHeader
      }).fail(function( jqXHR, textStatus, errorThrown ){
          alert(jqXHR.responseText); 
      });
  },
  updateline: function(line) {
      var data =  JSON.stringify({ "line" : line });
      return Em.$.ajax({
        url: '/api/1/projects/' + line.ProjectID + '/lines/' + line.ID,
        type: 'PUT',
        data: data,
        dataType: 'json',
        beforeSend: setHeader
      }).fail(function( jqXHR, textStatus, errorThrown ){
          alert(jqXHR.responseText); 
      });
  },
  deleteline: function(line) {
      var data =  JSON.stringify({ "line" : line });
      return Em.$.ajax({
        url: '/api/1/projects/' + line.ProjectID + '/lines/' + line.ID,
        type: 'DELETE',
        data: data,
        dataType: 'json',
        beforeSend: setHeader
      }).fail(function( jqXHR, textStatus, errorThrown ){
          alert(jqXHR.responseText); 
      });
  }
});