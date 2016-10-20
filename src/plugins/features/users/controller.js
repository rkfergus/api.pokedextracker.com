'use strict';

const Bluebird = require('bluebird');
const Bcrypt   = Bluebird.promisifyAll(require('bcrypt'));
const Slug     = require('slug');

const Config = require('../../../../config');
const Dex    = require('../../../models/dex');
const Errors = require('../../../libraries/errors');
const JWT    = require('../../../libraries/jwt');
const Knex   = require('../../../libraries/knex');
const User   = require('../../../models/user');

exports.list = function (query) {
  return new User().query((qb) => {
    qb.orderBy('id', 'DESC');
    qb.limit(query.limit);
    qb.offset(query.offset);
  })
  .fetchAll();
};

exports.retrieve = function (username) {
  return new User().where('username', username).fetch({ require: true })
  .catch(User.NotFoundError, () => {
    throw new Errors.NotFound('user');
  });
};

exports.create = function (payload, request) {
  return Bcrypt.hashAsync(payload.password, Config.SALT_ROUNDS)
  .then((hash) => {
    const xff = request.headers['x-forwarded-for'];
    const ip = xff ? xff.split(',')[0].trim() : request.info.remoteAddress;

    payload.password = hash;
    payload.last_ip = ip;

    return new User().where('username', payload.username).fetch();
  })
  .then((existing) => {
    if (existing) {
      throw new Errors.ExistingUsername();
    }

    const title = 'Living Dex';

    return Knex.transaction((transacting) => {
      return new User().save(payload, { transacting })
      .tap((user) => {
        return new Dex().save({
          user_id: user.id,
          title,
          slug: Slug(title, { lower: true })
        }, { transacting });
      });
    });
  })
  .then((user) => user.refresh())
  .then((user) => JWT.sign(user))
  .catch(Errors.DuplicateKey, () => {
    throw new Errors.ExistingUsername();
  });
};

exports.update = function (username, payload, auth) {
  return Bluebird.resolve()
  .then(() => {
    if (payload.password) {
      return Bcrypt.hashAsync(payload.password, Config.SALT_ROUNDS)
      .then((hash) => payload.password = hash);
    }
  })
  .then(() => new User({ id: auth.id }).where('username', username).save(payload))
  .then((user) => user.refresh())
  .then((user) => JWT.sign(user))
  .catch(User.NoRowsUpdatedError, () => {
    throw new Errors.ForbiddenAction('updating this user');
  });
};
