'use strict';

const Bluebird = require('bluebird');
const Slug     = require('slug');

const Dex    = require('../../../models/dex');
const Errors = require('../../../libraries/errors');

exports.retrieve = function (params) {
  return new Dex().query((qb) => {
    qb.innerJoin('users', 'dexes.user_id', 'users.id');
    qb.where({ username: params.username, slug: params.slug });
  }).fetch({ require: true })
  .catch(Dex.NotFoundError, () => {
    throw new Errors.NotFound('dex');
  });
};

exports.create = function (params, payload, auth) {
  return Bluebird.resolve()
  .then(() => {
    if (params.username !== auth.username) {
      throw new Errors.ForbiddenAction('creating a dex for this user');
    }

    payload.user_id = auth.id;
    payload.slug = Slug(payload.title, { lower: true });

    return new Dex().where({ user_id: auth.id, slug: payload.slug }).fetch();
  })
  .then((existing) => {
    if (existing) {
      throw new Errors.ExistingDex();
    }

    return new Dex().save(payload);
  })
  .then((dex) => dex.refresh())
  .catch(Errors.DuplicateKey, () => {
    throw new Errors.ExistingDex();
  });
};

exports.update = function (params, payload, auth) {
  return Bluebird.resolve()
  .then(() => {
    if (params.username !== auth.username) {
      throw new Errors.ForbiddenAction('updating a dex for this user');
    }

    return new Dex().where({ user_id: auth.id, slug: params.slug }).fetch({ require: true });
  })
  .then((dex) => {
    if (payload.title) {
      payload.slug = Slug(payload.title, { lower: true });
    }

    return dex.save(payload, { patch: true });
  })
  .then((dex) => dex.refresh())
  .catch(Dex.NotFoundError, () => {
    throw new Errors.NotFound('dex');
  })
  .catch(Errors.DuplicateKey, () => {
    throw new Errors.ExistingDex();
  });
};

exports.delete = function (params, auth) {
  return Bluebird.resolve()
  .then(() => {
    if (params.username !== auth.username) {
      throw new Errors.ForbiddenAction('deleting a dex for this user');
    }

    return new Dex().where({ user_id: auth.id, slug: params.slug }).fetch({ require: true });
  })
  .then((dex) => dex.destroy())
  .then(() => ({ deleted: true }))
  .catch(Dex.NotFoundError, () => {
    throw new Errors.NotFound('dex');
  });
};
