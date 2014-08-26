App.DayController = Ember.ObjectController.extend({
  needs: ["dayIndex"],
  newItemText : '',
  newItemStatus: 'OPEN',
  statuses: ["NOTE", "OPEN"],
  searchChangeStatus: function() {
        if (this.get('newItemText') == '/') {
          this.set('newItemStatus','NOTE');
          this.set('newItemText', '');
        }
    if (this.get('newItemText') == '.') {
          this.set('newItemStatus','OPEN');
          this.set('newItemText', '');
        }
    }.observes("newItemText"),
  isNote: function() {
    return this.get('newItemStatus') == 'NOTE';
  }.property('newItemStatus'),
  isOpen: function() {
    return this.get('newItemStatus') == 'OPEN';
  }.property('newItemStatus'),
  actions: {
    addItem: function() {
      var newItem = { 
        Day: parseInt(this.get("model.id")),
        Status: this.get('newItemStatus'), 
        Text: this.get('newItemText'), 
        ProjectID: "",
        ProjectItemID: ""
      };
      // check for blank - disable save button
      var controller = this;
      App.DayItem.save(newItem).then( function(data){
        console.log(data.dayitem);
        var model = App.DayItem.create(data.dayitem); 
        controller.get('controllers.dayIndex.model').addObject(model);
        controller.set('newItemText', '');
        controller.set('newItemStatus', 'OPEN');
      });
    },
    toggleStatus: function() {  
      if(this.get('newItemStatus') == 'OPEN') {
        this.set('newItemStatus', 'NOTE');
      } else if(this.get('newItemStatus') == 'NOTE') {
        this.set('newItemStatus', 'OPEN');  
      }
    }
  } 
});