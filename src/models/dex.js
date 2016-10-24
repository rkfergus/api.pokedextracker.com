'use strict';

const Bookshelf = require('../libraries/bookshelf');

module.exports = Bookshelf.model('Dex', Bookshelf.Model.extend({
  tableName: 'dexes',
  hasTimestamps: ['date_created', 'date_modified'],
  serialize () {
    return {
      id: this.get('id'),
      user_id: this.get('user_id'),
      title: this.get('title'),
      slug: this.get('slug'),
      shiny: this.get('shiny'),
      generation: this.get('generation'),
      date_created: this.get('date_created'),
      date_modified: this.get('date_modified')
    };
  }
}));
