'use strict';

const stones = ['fire', 'water', 'thunder', 'leaf', 'moon', 'sun', 'shiny', 'dusk', 'dawn', 'ice'];

exports.up = function (knex, Promise) {
  const array = stones.map((stone) => `'${stone}'::text`).join(', ');

  return knex.raw('ALTER TABLE evolutions DROP CONSTRAINT evolutions_stone_check')
  .then(() => knex.raw(`ALTER TABLE evolutions ADD CONSTRAINT evolutions_stone_check CHECK (stone = ANY (ARRAY[${array}])) NOT VALID`))
  .then(() => knex.raw('ALTER TABLE evolutions VALIDATE CONSTRAINT evolutions_stone_check'));
};

exports.down = function (knex, Promise) {
  const array = stones.filter((stone) => stone !== 'ice').map((stone) => `'${stone}'::text`).join(', ');

  return knex.raw('ALTER TABLE evolutions DROP CONSTRAINT evolutions_stone_check')
  .then(() => knex.raw(`ALTER TABLE evolutions ADD CONSTRAINT evolutions_stone_check CHECK (stone = ANY (ARRAY[${array}])) NOT VALID`))
  .then(() => knex.raw('ALTER TABLE evolutions VALIDATE CONSTRAINT evolutions_stone_check'));
};
