'use strict';

exports.up = function (knex, Promise) {
  return knex.schema.createTable('dexes', (table) => {
    table.increments('id').index();
    table.integer('user_id').references('id').inTable('users').onDelete('CASCADE').notNullable().index();
    table.text('title').notNullable();
    table.text('slug').notNullable().index();
    table.boolean('shiny').notNullable();
    table.integer('generation').notNullable();
    table.timestamp('date_created').notNullable().defaultTo(knex.fn.now());
    table.timestamp('date_modified').notNullable().defaultTo(knex.fn.now());

    table.unique(['user_id', 'slug']);
  });
};

exports.down = function (knex, Promise) {
  return knex.schema.dropTable('dexes');
};
