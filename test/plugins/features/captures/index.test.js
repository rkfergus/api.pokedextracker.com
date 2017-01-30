'use strict';

const Bluebird = require('bluebird');
const JWT      = require('jsonwebtoken');

const Config = require('../../../../config');
const Knex   = require('../../../../src/libraries/knex');
const Server = require('../../../../src/server');

const firstPokemon  = Factory.build('pokemon', { id: 1, national_id: 1 });
const secondPokemon = Factory.build('pokemon', { id: 2, national_id: 2 });

const user = Factory.build('user');

const dex = Factory.build('dex', { user_id: user.id });

const firstCapture = Factory.build('capture', { pokemon_id: firstPokemon.id, user_id: user.id, dex_id: dex.id });

const auth = `Bearer ${JWT.sign(user, Config.JWT_SECRET)}`;

describe('captures integration', () => {

  beforeEach(() => {
    return Bluebird.all([
      Knex('pokemon').insert([firstPokemon, secondPokemon]),
      Knex('users').insert(user)
    ])
    .then(() => Knex('dexes').insert(dex))
    .then(() => Knex('captures').insert(firstCapture));
  });

  describe('list', () => {

    it('returns a collection of captures', () => {
      return Server.inject({
        method: 'GET',
        url: `/captures?user=${user.id}`
      })
      .then((res) => {
        expect(res.statusCode).to.eql(200);
        expect(res.result).to.be.an.instanceof(Array);
      });
    });

  });

  describe('create', () => {

    it('saves a capture', () => {
      return Server.inject({
        method: 'POST',
        url: '/captures',
        headers: { authorization: auth },
        payload: { pokemon: secondPokemon.id, dex: dex.id }
      })
      .then((res) => {
        expect(res.statusCode).to.eql(200);
        expect(res.result[0].pokemon.id).to.eql(secondPokemon.id);
      });
    });

    it('requires authentication', () => {
      return Server.inject({
        method: 'POST',
        url: '/captures',
        payload: { pokemon: secondPokemon.id, dex: dex.id }
      })
      .then((res) => {
        expect(res.statusCode).to.eql(401);
      });
    });

  });

  describe('delete', () => {

    it('deletes a capture', () => {
      return Server.inject({
        method: 'DELETE',
        url: '/captures',
        headers: { authorization: auth },
        payload: { pokemon: firstPokemon.id, dex: dex.id }
      })
      .then((res) => {
        expect(res.statusCode).to.eql(200);
        expect(res.result.deleted).to.be.true;
      });
    });

    it('requires authentication', () => {
      return Server.inject({
        method: 'DELETE',
        url: '/captures',
        payload: { pokemon: firstPokemon.id, dex: dex.id }
      })
      .then((res) => {
        expect(res.statusCode).to.eql(401);
      });
    });

  });

});
