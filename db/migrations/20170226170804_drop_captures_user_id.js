'use strict';

exports.up = function (knex, Promise) {
  return knex.schema.table('captures', (table) => {
    table.dropColumn('user_id');
  });
};

exports.down = function (knex, Promise) {
  return knex.schema.table('captures', (table) => {
    table.integer('user_id').references('id').inTable('users').onDelete('CASCADE').index();
  });
};
