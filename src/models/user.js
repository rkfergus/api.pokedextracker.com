'use strict';

const Bookshelf = require('../libraries/bookshelf');

module.exports = Bookshelf.model('User', Bookshelf.Model.extend({
  tableName: 'users',
  hasTimestamps: ['date_created', 'date_modified'],
  dexes () {
    return this.hasMany('Dex');
  },
  virtuals: {
    jwt_summary () {
      return {
        id: this.get('id'),
        username: this.get('username'),
        friend_code: this.get('friend_code'),
        date_created: this.get('date_created'),
        date_modified: this.get('date_modified')
      };
    }
  },
  serialize () {
    return {
      id: this.get('id'),
      username: this.get('username'),
      friend_code: this.get('friend_code'),
      dexes: this.related('dexes').map((dex) => dex.serialize()),
      date_created: this.get('date_created'),
      date_modified: this.get('date_modified')
    };
  }
}, {
  RELATED: [{ dexes: (qb) => qb.orderBy('date_created', 'DESC') }]
}));
