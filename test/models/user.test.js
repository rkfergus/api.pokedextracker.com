'use strict';

const Dex  = require('../../src/models/dex');
const User = require('../../src/models/user');

const user = Factory.build('user');

const dex = Factory.build('dex', { user_id: user.id });

describe('user model', () => {

  describe('virtuals', () => {

    describe('jwt_summary', () => {

      it('only includes the fields needed for the JWT', () => {
        expect(User.forge(user).get('jwt_summary')).to.have.all.keys([
          'id',
          'username',
          'friend_code',
          'date_created',
          'date_modified'
        ]);
      });

    });

  });

  describe('serialize', () => {

    it('returns the correct fields', () => {
      const model = User.forge(user);
      model.relations.dexes = [Dex.forge(dex), Dex.forge(dex)];
      const json = model.serialize();

      expect(json).to.have.all.keys([
        'id',
        'username',
        'friend_code',
        'dexes',
        'date_created',
        'date_modified'
      ]);
      expect(json.dexes).to.have.length(2);
    });

  });

});
