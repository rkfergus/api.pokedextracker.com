'use strict';

const Config = require('../../../../config');
const Knex   = require('../../../../src/libraries/knex');
const Server = require('../../../../src/server');
const User   = require('../../../../src/models/user');

const user = Factory.build('user', { password: Config.PASSWORD_HASH });

describe('sessions integration', () => {

  describe('create', () => {

    beforeEach(() => {
      return Knex('users').insert(user);
    });

    it('creates a token', () => {
      return Server.inject({
        method: 'POST',
        url: '/sessions',
        payload: {
          username: user.username,
          password: Config.PASSWORD_VALUE
        }
      })
      .then((res) => {
        expect(res.statusCode).to.eql(200);
      });
    });

    it('saves an X-Forwarded-For IP address if present', () => {
      const ip = '123.123.123.123';

      return Server.inject({
        method: 'POST',
        url: '/sessions',
        headers: { 'X-Forwarded-For': ip },
        payload: {
          username: user.username,
          password: Config.PASSWORD_VALUE
        }
      })
      .then(() => new User().where('username', user.username).fetch())
      .then((u) => {
        expect(u.get('last_ip')).to.eql(ip);
      });
    });

    it('saves the first X-Forwarded-For IP address if multiple are present', () => {
      const ips = ['123.123.123.123', '124.124.124.124'];

      return Server.inject({
        method: 'POST',
        url: '/sessions',
        headers: { 'X-Forwarded-For': ips.join(',') },
        payload: {
          username: user.username,
          password: Config.PASSWORD_VALUE
        }
      })
      .then(() => new User().where('username', user.username).fetch())
      .then((u) => {
        expect(u.get('last_ip')).to.eql(ips[0]);
      });
    });

  });

});
