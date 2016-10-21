'use strict';

exports.up = function (knex, Promise) {
  return knex('users').select('users.id').leftOuterJoin('dexes', 'users.id', 'dexes.user_id').whereNull('dexes.id')
  .map((user) => {
    return {
      user_id: user.id,
      title: 'Living Dex',
      slug: 'living-dex',
      shiny: false,
      generation: 6,
      date_created: new Date(),
      date_modified: new Date()
    };
  })
  .then((dexes) => knex('dexes').insert(dexes));
};

exports.down = function (knex, Promise) {
  return knex('dexes').where('title', 'Living Dex').delete();
};
