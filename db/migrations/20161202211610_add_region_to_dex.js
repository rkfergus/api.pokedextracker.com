'use strict';

exports.up = function (knex, Promise) {
  return knex.schema.table('dexes', (table) => {
    table.enum('region', ['national', 'kalos', 'hoenn', 'alola']);
  });
};

exports.down = function (knex, Promise) {
  return knex.schema.table('dexes', (table) => {
    table.dropColumn('region');
  });
};
