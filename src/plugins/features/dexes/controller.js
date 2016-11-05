'use strict';

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
