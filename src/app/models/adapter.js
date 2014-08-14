// not production ready

// Ember data

//App.ApplicationAdapter = DS.FixtureAdapter.extend();
DS.RESTAdapter.reopen({
  namespace: 'api/1'
});

App.ApplicationSerializer = DS.RESTSerializer.extend({
  serializeHasMany: function(record, json, relationship) {
    var key = relationship.key;
    var payloadKey = this.keyForRelationship ? this.keyForRelationship(key, 'hasMany') : key;
    var relationshipType = DS.RelationshipChange.determineRelationshipType(record.constructor, relationship);

    if (['manyToNone', 'manyToMany', 'manyToOne'].contains(relationshipType)) {
      json[payloadKey] = record.get(key).mapBy('id');
    }
  }
});

App.List = DS.Model.extend({
  title: DS.attr('string'),
  current: DS.attr('number'),
  total: DS.attr('number'),
  description: DS.attr('string'),
  items: DS.hasMany('item', { async: true })
});

App.Item = DS.Model.extend({
  name: DS.attr('string'),
  isDone: DS.attr('boolean'),
  list: DS.belongsTo('list')
});

App.List.FIXTURES = [
  {
    id: 200,
    title: 'List 1',
    description: 'The first list',
    current: 1,
    total: 2,
    items: [1, 2]
  },
  {
    id: 201,
    title: 'List 2',
    description: 'The second list',
    current: 3,
    total: 4,
    items: [3, 4, 5, 6]
  }
];

App.Item.FIXTURES = [
 {  
    id: 1,
    name: 'Flint',
    isDone: true,
    list: 200
  },
  {
    id: 2,
    name: 'Kindling',
    isDone: false,
    list: 200
  },
  {
    id: 3,
    name: 'Matches',
    isDone: true,
    list: 201
  },
  {
    id: 4,
    name: 'Bow Drill',
    isDone: false,
    list: 201
  },
  {
    id: 5,
    name: 'Tinder',
    isDone: true,
    list: 201
  },
  {
    id: 6,
    name: 'Birch Bark Shaving',
    isDone: true,
    list: 201
  }
];