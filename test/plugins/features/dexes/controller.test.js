'use strict';

const Controller = require('../../../../src/plugins/features/dexes/controller');
const Errors     = require('../../../../src/libraries/errors');
const Knex       = require('../../../../src/libraries/knex');

const firstUser = Factory.build('user');

const firstDex = Factory.build('dex', { user_id: firstUser.id });

describe('dexes controller', () => {

  describe('retrieve', () => {

    beforeEach(() => {
      return Knex('users').insert(firstUser)
      .then(() => Knex('dexes').insert(firstDex));
    });

    it('returns a dex from the username and slug', () => {
      return Controller.retrieve({ username: firstUser.username, slug: firstDex.slug })
      .then((dex) => {
        expect(dex.id).to.eql(firstDex.id);
      });
    });

    it('rejects if the username does not match with the slug', () => {
      return Controller.retrieve({ username: 'bad_username', slug: firstDex.slug })
      .catch((err) => err)
      .then((err) => {
        expect(err).to.be.an.instanceof(Errors.NotFound);
      });
    });

    it('rejects if the slug does not match with the username', () => {
      return Controller.retrieve({ username: firstUser.username, slug: 'bad-slug' })
      .catch((err) => err)
      .then((err) => {
        expect(err).to.be.an.instanceof(Errors.NotFound);
      });
    });

  });

});
