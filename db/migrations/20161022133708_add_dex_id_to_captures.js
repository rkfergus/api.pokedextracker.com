'use strict';

exports.up = function (knex, Promise) {
  return knex.schema.table('captures', (table) => {
    table.integer('dex_id').references('id').inTable('dexes').onDelete('CASCADE').index();
  });
};

exports.down = function (knex, Promise) {
  return knex.schema.table('captures', (table) => {
    table.dropColumn('dex_id');
  });
};
