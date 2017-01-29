'use strict';

exports.up = function (knex, Promise) {
  return knex.schema.table('pokemon', (table) => {
    table.specificType('id', 'serial').unique().index();
    table.text('form');
    table.text('box');
    table.integer('national_order');
  })
  .then(() => knex.raw('UPDATE pokemon SET id = id + (SELECT MAX(id) FROM pokemon)'))
  .then(() => knex.raw('UPDATE pokemon SET id = national_id, national_order = national_id'))
  .then(() => knex.raw('ALTER TABLE captures ADD CONSTRAINT captures_pokemon_id_foreign2 FOREIGN KEY (pokemon_id) REFERENCES pokemon(id) ON DELETE CASCADE NOT VALID'))
  .then(() => knex.raw('ALTER TABLE captures DROP CONSTRAINT captures_pokemon_id_foreign'))
  .then(() => knex.raw('ALTER TABLE captures RENAME CONSTRAINT captures_pokemon_id_foreign2 TO captures_pokemon_id_foreign'))
  .then(() => knex.raw('ALTER TABLE captures VALIDATE CONSTRAINT captures_pokemon_id_foreign'))
  .then(() => knex.raw('ALTER TABLE evolutions ADD CONSTRAINT evolutions_evolved_pokemon_id_foreign2 FOREIGN KEY (evolved_pokemon_id) REFERENCES pokemon(id) ON DELETE CASCADE NOT VALID'))
  .then(() => knex.raw('ALTER TABLE evolutions DROP CONSTRAINT evolutions_evolved_pokemon_id_foreign'))
  .then(() => knex.raw('ALTER TABLE evolutions RENAME CONSTRAINT evolutions_evolved_pokemon_id_foreign2 TO evolutions_evolved_pokemon_id_foreign'))
  .then(() => knex.raw('ALTER TABLE evolutions VALIDATE CONSTRAINT evolutions_evolved_pokemon_id_foreign'))
  .then(() => knex.raw('ALTER TABLE evolutions ADD CONSTRAINT evolutions_evolving_pokemon_id_foreign2 FOREIGN KEY (evolving_pokemon_id) REFERENCES pokemon(id) ON DELETE CASCADE NOT VALID'))
  .then(() => knex.raw('ALTER TABLE evolutions DROP CONSTRAINT evolutions_evolving_pokemon_id_foreign'))
  .then(() => knex.raw('ALTER TABLE evolutions RENAME CONSTRAINT evolutions_evolving_pokemon_id_foreign2 TO evolutions_evolving_pokemon_id_foreign'))
  .then(() => knex.raw('ALTER TABLE evolutions VALIDATE CONSTRAINT evolutions_evolving_pokemon_id_foreign'))
  .then(() => knex.raw(`
    BEGIN;
    ALTER TABLE pokemon DROP CONSTRAINT pokemon_pkey;
    ALTER TABLE pokemon ADD PRIMARY KEY (id);
    COMMIT;
  `))
  .then(() => knex.raw('ALTER TABLE pokemon DROP CONSTRAINT IF EXISTS pokemon_national_id_unique'))
  .then(() => knex.raw('ALTER TABLE pokemon ALTER COLUMN national_order SET NOT NULL'));
};

exports.down = function (knex, Promise) {
  return knex.raw('ALTER TABLE pokemon ADD CONSTRAINT pokemon_national_id_unique UNIQUE (national_id)')
  .then(() => knex.raw('ALTER TABLE captures ADD CONSTRAINT captures_pokemon_id_foreign2 FOREIGN KEY (pokemon_id) REFERENCES pokemon(national_id) ON DELETE CASCADE NOT VALID'))
  .then(() => knex.raw('ALTER TABLE captures DROP CONSTRAINT captures_pokemon_id_foreign'))
  .then(() => knex.raw('ALTER TABLE captures RENAME CONSTRAINT captures_pokemon_id_foreign2 TO captures_pokemon_id_foreign'))
  .then(() => knex.raw('ALTER TABLE evolutions ADD CONSTRAINT evolutions_evolved_pokemon_id_foreign2 FOREIGN KEY (evolved_pokemon_id) REFERENCES pokemon(national_id) ON DELETE CASCADE NOT VALID'))
  .then(() => knex.raw('ALTER TABLE evolutions DROP CONSTRAINT evolutions_evolved_pokemon_id_foreign'))
  .then(() => knex.raw('ALTER TABLE evolutions RENAME CONSTRAINT evolutions_evolved_pokemon_id_foreign2 TO evolutions_evolved_pokemon_id_foreign'))
  .then(() => knex.raw('ALTER TABLE evolutions ADD CONSTRAINT evolutions_evolving_pokemon_id_foreign2 FOREIGN KEY (evolving_pokemon_id) REFERENCES pokemon(national_id) ON DELETE CASCADE NOT VALID'))
  .then(() => knex.raw('ALTER TABLE evolutions DROP CONSTRAINT evolutions_evolving_pokemon_id_foreign'))
  .then(() => knex.raw('ALTER TABLE evolutions RENAME CONSTRAINT evolutions_evolving_pokemon_id_foreign2 TO evolutions_evolving_pokemon_id_foreign'))
  .then(() => knex.raw(`
    BEGIN;
    ALTER TABLE pokemon DROP CONSTRAINT pokemon_pkey;
    ALTER TABLE pokemon ADD PRIMARY KEY (national_id);
    COMMIT;
  `))
  .then(() => {
    return knex.schema.table('pokemon', (table) => {
      table.dropColumn('id');
      table.dropColumn('form');
      table.dropColumn('box');
      table.dropColumn('national_order');
    });
  });
};

exports.config = { transaction: false };
