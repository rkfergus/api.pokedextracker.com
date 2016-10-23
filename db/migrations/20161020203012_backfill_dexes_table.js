'use strict';

const LIMIT = 1000;

function batch (knex) {
  return knex('users')
    .select('users.id')
    .leftOuterJoin('dexes', 'users.id', 'dexes.user_id')
    .whereNull('dexes.id')
    .limit(LIMIT)
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
  .tap((dexes) => knex('dexes').insert(dexes))
  .then((dexes) => {
    return dexes.length === LIMIT ? batch(knex) : null;
  });
}

exports.up = function (knex, Promise) {
  return batch(knex);
};

exports.down = function (knex, Promise) {
  return knex('dexes').where('title', 'Living Dex').delete();
};

exports.config = {
  transaction: false
};
