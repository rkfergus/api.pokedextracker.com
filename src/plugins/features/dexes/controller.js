'use strict';

const Bluebird = require('bluebird');
const Slug     = require('slug');

const Capture = require('../../../models/capture');
const Dex     = require('../../../models/dex');
const Errors  = require('../../../libraries/errors');
const Knex    = require('../../../libraries/knex');

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

    let captures;

    if (payload.generation || payload.region) {
      captures = new Capture().query((qb) => {
        qb.where('dex_id', dex.get('id'));
        qb.whereIn('pokemon_id', function () {
          this.select('id').from('pokemon');
          if (payload.generation) {
            this.where('generation', '>', payload.generation);
          }
          if (payload.region) {
            this.orWhereNull(`${payload.region}_id`);
          }
        });
      });
    }

    return Knex.transaction((transacting) => {
      return Bluebird.all([
        dex.save(payload, { patch: true, transacting }),
        captures && captures.destroy({ transacting })
      ]);
    });
  })
  .spread((dex) => dex.refresh())
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

    return Knex.transaction((transacting) => {
      return Dex.where({ user_id: auth.id }).count({ transacting })
      .then((count) => {
        if (parseInt(count) === 1) {
          throw new Errors.AtLeastOneDex();
        }

        return new Dex().where({ user_id: auth.id, slug: params.slug }).destroy({ require: true, transacting });
      });
    });
  })
  .then(() => ({ deleted: true }))
  .catch(Dex.NoRowsDeletedError, () => {
    throw new Errors.NotFound('dex');
  });
};
