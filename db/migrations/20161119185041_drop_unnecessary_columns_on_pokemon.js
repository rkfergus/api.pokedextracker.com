'use strict';

const TYPES = [
  'normal',
  'fighting',
  'flying',
  'poison',
  'ground',
  'rock',
  'bug',
  'ghost',
  'steel',
  'fire',
  'water',
  'grass',
  'electric',
  'psychic',
  'ice',
  'dragon',
  'dark',
  'fairy'
];

exports.up = function (knex, Promise) {
  return knex.schema.table('pokemon', (table) => {
    table.dropColumn('regionless');
    table.dropColumn('type1');
    table.dropColumn('type2');
    table.dropColumn('icon_url');
    table.dropColumn('avatar_url');
  })
  .then(() => {
    return knex.raw(`
      ALTER TABLE pokemon
        ALTER COLUMN x_location DROP NOT NULL,
        ALTER COLUMN y_location DROP NOT NULL,
        ALTER COLUMN or_location DROP NOT NULL,
        ALTER COLUMN as_location DROP NOT NULL
    `);
  });
};

exports.down = function (knex, Promise) {
  return knex.schema.table('pokemon', (table) => {
    table.boolean('regionless').notNullable().defaultTo(false);
    table.enum('type1', TYPES).notNullable();
    table.enum('type2', TYPES);
    table.string('icon_url').notNullable();
    table.string('avatar_url').notNullable();
  })
  .then(() => {
    return knex.raw(`
      ALTER TABLE pokemon
        ALTER COLUMN x_location SET NOT NULL,
        ALTER COLUMN y_location SET NOT NULL,
        ALTER COLUMN or_location SET NOT NULL,
        ALTER COLUMN as_location SET NOT NULL
    `);
  });
};
