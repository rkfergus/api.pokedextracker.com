'use strict';

const Sinon = require('sinon');

const Controller = require('../../../../src/plugins/features/users/controller');
const Dex        = require('../../../../src/models/dex');
const Errors     = require('../../../../src/libraries/errors');
const Knex       = require('../../../../src/libraries/knex');
const User       = require('../../../../src/models/user');

const firstUser      = Factory.build('user');
const secondUser     = Factory.build('user');
const friendCodeUser = Factory.build('user', { friend_code: '0000-0000-0000' });

describe('user controller', () => {

  describe('list', () => {

    beforeEach(() => {
      return Knex('users').insert([firstUser, secondUser]);
    });

    it('returns a collection of users ordered by id descending', () => {
      return Controller.list({ limit: 10, offset: 0 })
      .get('models')
      .map((user) => user.id)
      .then((users) => {
        expect(users).to.have.length(2);
        expect(users[0]).to.eql(secondUser.id);
        expect(users[1]).to.eql(firstUser.id);
      });
    });

    it('utilized a limit that is passed in', () => {
      return Controller.list({ limit: 1, offset: 0 })
      .get('models')
      .map((user) => user.id)
      .then((users) => {
        expect(users).to.have.length(1);
        expect(users[0]).to.eql(secondUser.id);
      });
    });

    it('utilized an offset that is passed in', () => {
      return Controller.list({ limit: 10, offset: 1 })
      .get('models')
      .map((user) => user.id)
      .then((users) => {
        expect(users).to.have.length(1);
        expect(users[0]).to.eql(firstUser.id);
      });
    });

  });

  describe('retrieve', () => {

    beforeEach(() => {
      return Knex('users').insert(firstUser);
    });

    it('returns an individual user from its username', () => {
      return Controller.retrieve(firstUser.username)
      .then((user) => {
        expect(user.id).to.eql(firstUser.id);
      });
    });

    it('rejects if the username does not exist', () => {
      return Controller.retrieve('bad_username')
      .catch((err) => err)
      .then((err) => {
        expect(err).to.be.an.instanceof(Errors.NotFound);
      });
    });

  });

  describe('create', () => {

    const request = { headers: {}, info: {} };

    it('saves a user with a hashed password', () => {
      const username = 'test';
      const password = 'test';

      return Controller.create({ username, password }, request)
      .then((session) => {
        expect(session.token).to.be.a('string');
      })
      .then(() => new User().where('username', username).fetch())
      .then((user) => {
        expect(user.get('password')).to.not.eql(password);
        expect(user.get('password')).to.have.length(60);
      });
    });

    it('saves last login date', () => {
      const username = 'test';

      return Controller.create({ username, password: 'test' }, request)
      .then(() => new User().where('username', username).fetch())
      .then((user) => {
        expect(user.get('last_login')).to.be.an.instanceof(Date);
      });
    });

    it('saves referrer', () => {
      const username = 'test';
      const referrer = 'http://test.com';

      return Controller.create({ username, password: 'test', referrer }, request)
      .then(() => new User().where('username', username).fetch())
      .then((user) => {
        expect(user.get('referrer')).to.eql(referrer);
      });
    });

    it('saves a default dex', () => {
      const username = 'test';

      return Controller.create({ username, password: 'test' }, request)
      .then(() => new User().where('username', username).fetch())
      .then((user) => new Dex().where('user_id', user.id).fetch())
      .then((dex) => {
        expect(dex.get('title')).to.eql('Living Dex');
        expect(dex.get('slug')).to.eql('living-dex');
      });
    });

    it('rejects if the username is already taken', () => {
      return Knex('users').insert(firstUser)
      .then(() => {
        return Controller.create({ username: firstUser.username, password: 'test' }, request);
      })
      .catch((err) => err)
      .then((err) => {
        expect(err).to.be.an.instanceof(Errors.ExistingUsername);
      });
    });

    it('rejects if the username is taken after the fetch', () => {
      Sinon.stub(User.prototype, 'save').throws(new Error('duplicate key value'));

      return Controller.create({ username: firstUser.username, password: 'test' }, request)
      .catch((err) => err)
      .then((err) => {
        expect(err).to.be.an.instanceof(Errors.ExistingUsername);
      })
      .finally(() => {
        User.prototype.save.restore();
      });
    });

  });

  describe('update', () => {

    beforeEach(() => {
      return Knex('users').insert([firstUser, secondUser, friendCodeUser]);
    });

    it('does not clear the password if it is not passed in', () => {
      return Controller.update(firstUser.username, { password: undefined }, { id: firstUser.id })
      .then(() => new User({ id: firstUser.id }).fetch())
      .then((user) => {
        expect(user.get('password')).to.exist;
      });
    });

    it('updates a user', () => {
      const friendCode = '4321-4321-4321';

      return Controller.update(firstUser.username, { friend_code: friendCode }, { id: firstUser.id })
      .then(() => new User({ id: firstUser.id }).fetch())
      .then((user) => {
        expect(user.get('friend_code')).to.eql(friendCode);
      });
    });

    it('rehashes password if it was passed in', () => {
      const password = 'test';

      return Controller.update(firstUser.username, { password }, { id: firstUser.id })
      .then(() => new User({ id: firstUser.id }).fetch())
      .then((user) => {
        expect(user.get('password')).to.not.eql(password);
        expect(user.get('password')).to.have.length(60);
      });
    });

    it('clears out the friend code if null is passed in', () => {
      return Controller.update(friendCodeUser.username, { friend_code: null }, { id: friendCodeUser.id })
      .then(() => new User({ id: friendCodeUser.id }).fetch())
      .then((user) => {
        expect(user.get('friend_code')).to.not.exist;
      });
    });

    it('returns a new session', () => {
      const friendCode = '4321-4321-4321';

      return Controller.update(firstUser.username, { friend_code: friendCode }, { id: firstUser.id })
      .then((session) => {
        expect(session.token).to.be.a('string');
      });
    });

    it('rejects if the username and auth id do not match', () => {
      return Controller.update(firstUser.username, {}, { id: secondUser.id })
      .catch((err) => err)
      .then((err) => {
        expect(err).to.be.an.instanceof(Errors.ForbiddenAction);
      });
    });

  });

});
