'use strict';

exports.up = function (knex, Promise) {
  return knex('dexes').where('generation', 6).update({ region: 'national' })
  .then(() => knex('dexes').where('generation', 7).update({ region: 'alola' }))
  .then(() => knex.raw('ALTER TABLE dexes ALTER COLUMN region SET NOT NULL'));
};

exports.down = function (knex, Promise) {
  return knex.raw('ALTER TABLE dexes ALTER COLUMN region DROP NOT NULL');
};
