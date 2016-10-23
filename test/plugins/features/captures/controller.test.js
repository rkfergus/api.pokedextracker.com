'use strict';

const Bluebird = require('bluebird');

const Capture    = require('../../../../src/models/capture');
const Controller = require('../../../../src/plugins/features/captures/controller');
const Errors     = require('../../../../src/libraries/errors');
const Knex       = require('../../../../src/libraries/knex');
const Pokemon    = require('../../../../src/models/pokemon');

const firstPokemon  = Factory.build('pokemon', { national_id: 1 });
const secondPokemon = Factory.build('pokemon', { national_id: 2 });

const user      = Factory.build('user');
const otherUser = Factory.build('user');

const dex      = Factory.build('dex', { user_id: user.id });
const otherDex = Factory.build('dex', { user_id: otherUser.id });

const firstCapture = Factory.build('capture', { pokemon_id: firstPokemon.national_id, user_id: user.id, dex_id: dex.id });
const otherCapture = Factory.build('capture', { pokemon_id: firstPokemon.national_id, user_id: otherUser.id, dex_id: otherDex.id });

describe('capture controller', () => {

  beforeEach(() => {
    return Bluebird.all([
      Knex('pokemon').insert([firstPokemon, secondPokemon]),
      Knex('users').insert([user, otherUser])
    ])
    .then(() => Knex('dexes').insert([dex, otherDex]))
    .then(() => Knex('captures').insert([firstCapture, otherCapture]));
  });

  describe('list', () => {

    it('returns a collection of captures filtered by user_id, filling in those that do not exist', () => {
      return new Pokemon().query((qb) => qb.orderBy('national_id')).fetchAll()
      .get('models')
      .then((pokemon) => Controller.list({ user: user.id }, pokemon))
      .map((capture) => capture.serialize())
      .then((captures) => {
        expect(captures).to.have.length(2);
        expect(captures[0].pokemon.national_id).to.eql(firstPokemon.national_id);
        expect(captures[0].user_id).to.eql(user.id);
        expect(captures[0].captured).to.be.true;
        expect(captures[1].pokemon.national_id).to.eql(secondPokemon.national_id);
        expect(captures[1].user_id).to.eql(user.id);
        expect(captures[1].captured).to.be.false;
      });
    });

    it('returns a collection of captures filtered by dex_id, filling in those that do not exist', () => {
      return new Pokemon().query((qb) => qb.orderBy('national_id')).fetchAll()
      .get('models')
      .then((pokemon) => Controller.list({ dex: dex.id }, pokemon))
      .map((capture) => capture.serialize())
      .then((captures) => {
        expect(captures).to.have.length(2);
        expect(captures[0].pokemon.national_id).to.eql(firstPokemon.national_id);
        expect(captures[0].dex_id).to.eql(dex.id);
        expect(captures[0].captured).to.be.true;
        expect(captures[1].pokemon.national_id).to.eql(secondPokemon.national_id);
        expect(captures[1].dex_id).to.eql(dex.id);
        expect(captures[1].captured).to.be.false;
      });
    });

    it('rejects for a user id that does not exist', () => {
      return Controller.list({ user: -1 })
      .catch((err) => err)
      .then((err) => {
        expect(err).to.be.an.instanceof(Errors.NotFound);
      });
    });

    it('rejects for a dex id that does not exist', () => {
      return Controller.list({ dex: -1 })
      .catch((err) => err)
      .then((err) => {
        expect(err).to.be.an.instanceof(Errors.NotFound);
      });
    });

  });

  describe('create', () => {

    it('creates a capture', () => {
      return Controller.create({ pokemon: [secondPokemon.national_id] }, { id: user.id })
      .then((captures) => {
        expect(captures).to.have.length(1);
        expect(captures.at(0).get('pokemon_id')).to.eql(secondPokemon.national_id);
        expect(captures.at(0).get('user_id')).to.eql(user.id);
        expect(captures.at(0).get('dex_id')).to.eql(dex.id);
        expect(captures.at(0).get('captured')).to.be.true;
      });
    });

    it('rejects with a bad pokemon id', () => {
      return Controller.create({ pokemon: [-1] }, { id: user.id })
      .catch((err) => err)
      .then((err) => {
        expect(err).to.be.an.instanceof(Errors.NotFound);
      });
    });

    it('does not err for duplicate captures', () => {
      return Controller.create({ pokemon: [firstPokemon.national_id] }, { id: user.id })
      .then((captures) => {
        expect(captures).to.have.length(1);
        expect(captures.at(0).get('pokemon_id')).to.eql(firstPokemon.national_id);
        expect(captures.at(0).get('user_id')).to.eql(user.id);
        expect(captures.at(0).get('captured')).to.be.true;
      });
    });

  });

  describe('delete', () => {

    it('deletes a capture', () => {
      return Controller.delete({ pokemon: [firstPokemon.national_id] }, { id: user.id })
      .then((res) => {
        expect(res.deleted).to.be.true;

        return new Capture().where({ pokemon_id: firstPokemon.national_id, user_id: user.id }).fetch();
      })
      .then((capture) => {
        expect(capture).to.be.null;
      });
    });

  });

});
