'use strict';

const Bookshelf = require('../libraries/bookshelf');
const Knex      = require('../libraries/knex');

const TOTALS = {
  6: 721,
  7: 802
};

module.exports = Bookshelf.model('Dex', Bookshelf.Model.extend({
  tableName: 'dexes',
  hasTimestamps: ['date_created', 'date_modified'],
  caught () {
    return Knex('captures').count().where('dex_id', this.get('id'))
    .then((res) => parseInt(res[0].count));
  },
  virtuals: {
    total () {
      return TOTALS[this.get('generation')];
    }
  },
  serialize () {
    return this.caught()
    .then((caught) => {
      return {
        id: this.get('id'),
        user_id: this.get('user_id'),
        title: this.get('title'),
        slug: this.get('slug'),
        shiny: this.get('shiny'),
        generation: this.get('generation'),
        caught,
        total: this.get('total'),
        date_created: this.get('date_created'),
        date_modified: this.get('date_modified')
      };
    });
  }
}));
