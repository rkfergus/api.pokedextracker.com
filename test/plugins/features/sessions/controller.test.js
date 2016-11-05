'use strict';

const Config     = require('../../../../config');
const Controller = require('../../../../src/plugins/features/sessions/controller');
const Errors     = require('../../../../src/libraries/errors');
const Knex       = require('../../../../src/libraries/knex');
const User       = require('../../../../src/models/user');

const user = Factory.build('user', { password: Config.PASSWORD_HASH });

describe('sessions controller', () => {

  describe('create', () => {

    const request = { headers: {}, info: {} };

    beforeEach(() => {
      return Knex('users').insert(user);
    });

    it('signs a token if the password is valid', () => {
      return Controller.create({ username: user.username, password: Config.PASSWORD_VALUE }, request)
      .then((session) => {
        expect(session.token).to.be.a('string');
      });
    });

    it('updates last login date', () => {
      return Controller.create({ username: user.username, password: Config.PASSWORD_VALUE }, request)
      .then(() => new User().where('username', user.username).fetch())
      .then((u) => {
        expect(u.get('last_login')).to.be.an.instanceof(Date);
      });
    });

    it('rejects if the user is not found', () => {
      return Controller.create({ username: 'bad_username', password: Config.PASSWORD_VALUE }, request)
      .catch((err) => err)
      .then((err) => {
        expect(err).to.be.an.instanceof(Errors.NotFound);
      });
    });

    it('rejects if the password is incorrect', () => {
      return Controller.create({ username: user.username, password: 'test' }, request)
      .catch((err) => err)
      .then((err) => {
        expect(err).to.be.an.instanceof(Errors.InvalidPassword);
      });
    });

  });

});
