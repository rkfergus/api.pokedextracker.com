'use strict';

exports.up = function (knex, Promise) {
  return knex.schema.table('pokemon', (table) => {
    table.integer('generation');
    table.text('sun_location');
    table.text('moon_location');
  })
  .then(() => {
    return Promise.all([
      knex('pokemon').whereBetween('national_id', [1, 151]).update({ generation: 1 }),
      knex('pokemon').whereBetween('national_id', [152, 251]).update({ generation: 2 }),
      knex('pokemon').whereBetween('national_id', [252, 386]).update({ generation: 3 }),
      knex('pokemon').whereBetween('national_id', [387, 493]).update({ generation: 4 }),
      knex('pokemon').whereBetween('national_id', [494, 649]).update({ generation: 5 }),
      knex('pokemon').whereBetween('national_id', [650, 721]).update({ generation: 6 })
    ]);
  })
  .then(() => knex.raw('ALTER TABLE pokemon ALTER COLUMN generation SET NOT NULL'));
};

exports.down = function (knex, Promise) {
  return knex.schema.table('pokemon', (table) => {
    table.dropColumn('generation');
    table.dropColumn('sun_location');
    table.dropColumn('moon_location');
  });
};
