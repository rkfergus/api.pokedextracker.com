'use strict';

const Controller           = require('./controller');
const UsersCreateValidator = require('../../../validators/users/create');
const UsersListValidator   = require('../../../validators/users/list');
const UsersUpdateValidator = require('../../../validators/users/update');

exports.register = (server, options, next) => {

  server.route([{
    method: 'GET',
    path: '/users',
    config: {
      handler: (request, reply) => reply(Controller.list(request.query)),
      validate: { query: UsersListValidator }
    }
  }, {
    method: 'GET',
    path: '/users/{username}',
    config: {
      handler: (request, reply) => reply(Controller.retrieve(request.params.username))
    }
  }, {
    method: 'POST',
    path: '/users',
    config: {
      handler: (request, reply) => reply(Controller.create(request.payload, request)),
      validate: { payload: UsersCreateValidator }
    }
  }, {
    method: 'POST',
    path: '/users/{id}',
    config: {
      auth: 'token',
      handler: (request, reply) => reply(Controller.update(request.params.id, request.payload, request.auth.credentials)),
      validate: { payload: UsersUpdateValidator }
    }
  }]);

  next();

};

exports.register.attributes = {
  name: 'users'
};
