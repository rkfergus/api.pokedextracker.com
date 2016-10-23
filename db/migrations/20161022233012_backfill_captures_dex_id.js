'use strict';

const LIMIT = 10000;

function upBatch (knex) {
  return knex.raw(`
    UPDATE captures
    SET
      dex_id = dexes.id
    FROM users
    INNER JOIN dexes ON users.id = dexes.user_id
    WHERE
      captures.user_id = users.id
      AND (captures.pokemon_id, captures.user_id) IN (SELECT pokemon_id, user_id FROM captures WHERE dex_id IS NULL LIMIT ${LIMIT})
  `)
  .then((response) => {
    return response.rowCount === LIMIT ? upBatch(knex) : null;
  });
}

function downBatch (knex) {
  return knex.raw(`
    UPDATE captures
    SET
      dex_id = null
    WHERE
      (captures.pokemon_id, captures.user_id) IN (SELECT pokemon_id, user_id FROM captures WHERE dex_id IS NOT NULL LIMIT ${LIMIT})
  `)
  .then((response) => {
    return response.rowCount === LIMIT ? downBatch(knex) : null;
  });
}

exports.up = function (knex, Promise) {
  return upBatch(knex)
  .then(() => {
    return knex.raw(`
      BEGIN;
      ALTER TABLE captures DROP CONSTRAINT captures_pkey;
      ALTER TABLE captures ADD CONSTRAINT captures_pkey PRIMARY KEY (dex_id, pokemon_id);
      COMMIT;
    `);
  });
};

exports.down = function (knex, Promise) {
  return knex.raw(`
    BEGIN;
    ALTER TABLE captures DROP CONSTRAINT captures_pkey;
    ALTER TABLE captures ADD CONSTRAINT captures_pkey PRIMARY KEY (user_id, pokemon_id);
    ALTER TABLE captures ALTER COLUMN dex_id DROP NOT NULL;
    COMMIT;
  `)
  .then(() => downBatch(knex));
};

exports.config = {
  transaction: false
};
