App.ListEditRoute = Ember.Route.extend({
  deactivate: function() {
    var model = this.controller.content;
    if (model.get('isDirty') && !model.get('isSaving')) {
      model.save();
    }
    this.controller.get('model.items').forEach(function(item){
      if((item.get('isNew') ||  item.get('isDirty')) && !item.get('isSaving')) {
        item.save();
      }
    });
  }
});