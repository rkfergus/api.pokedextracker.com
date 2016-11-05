'use strict';

const Knex   = require('../../../../src/libraries/knex');
const Server = require('../../../../src/server');

const user = Factory.build('user');

const dex = Factory.build('dex', { user_id: user.id });

describe('dexes integration', () => {

  describe('retrieve', () => {

    beforeEach(() => {
      return Knex('users').insert(user)
      .then(() => Knex('dexes').insert(dex));
    });

    it('returns a dex from the username and slug', () => {
      return Server.inject({
        method: 'GET',
        url: `/users/${user.username}/dexes/${dex.slug}`
      })
      .then((res) => {
        expect(res.statusCode).to.eql(200);
      });
    });

  });

});
