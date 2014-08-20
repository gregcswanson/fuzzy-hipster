App.Router.map(function() {
	this.resource('lists', function() {
    	this.route('new');
	});
	this.resource('list', { path: 'list/:list_id' }, function() {
    	this.route("edit", { path: "/edit" });
		this.route("delete");
	});
	this.resource('projects', function() {
		this.route('new');
  });
  this.resource('project', { path: 'project/:project_id'}, function() {
    //this.route("edit", { path: "/edit" });
    this.route("delete");
  });
	this.route('about', { path: '/about' });
});