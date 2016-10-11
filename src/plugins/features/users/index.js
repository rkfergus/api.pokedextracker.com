'use strict';

const Controller          = require('./controller');
const UserCreateValidator = require('../../../validators/user/create');
const UserListValidator   = require('../../../validators/user/list');
const UserUpdateValidator = require('../../../validators/user/update');

exports.register = (server, options, next) => {

  server.route([{
    method: 'GET',
    path: '/users',
    config: {
      handler: (request, reply) => reply(Controller.list(request.query)),
      validate: { query: UserListValidator }
    }
  }, {
    method: 'GET',
    path: '/users/{id}',
    config: {
      handler: (request, reply) => reply(Controller.retrieve(request.params.id))
    }
  }, {
    method: 'POST',
    path: '/users',
    config: {
      handler: (request, reply) => reply(Controller.create(request.payload, request)),
      validate: { payload: UserCreateValidator }
    }
  }, {
    method: 'POST',
    path: '/users/{id}',
    config: {
      auth: 'token',
      handler: (request, reply) => reply(Controller.update(request.params.id, request.payload, request.auth.credentials)),
      validate: { payload: UserUpdateValidator }
    }
  }]);

  next();

};

exports.register.attributes = {
  name: 'users'
};
