'use strict';

exports.up = function (knex, Promise) {
  return knex.schema.table('pokemon', (table) => {
    table.integer('alola_id');
  });
};

exports.down = function (knex, Promise) {
  return knex.schema.table('pokemon', (table) => {
    table.dropColumn('alola_id');
  });
};
