'use strict';

exports.up = function (knex, Promise) {
  return knex.raw('ALTER TABLE captures ALTER COLUMN user_id DROP NOT NULL');
};

exports.down = function (knex, Promise) {
  return knex.raw('ALTER TABLE captures ALTER COLUMN user_id SET NOT NULL');
};
