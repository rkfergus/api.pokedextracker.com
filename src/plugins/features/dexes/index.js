'use strict';

const Controller           = require('./controller');
const DexesCreateValidator = require('../../../validators/dexes/create');
const DexesUpdateValidator = require('../../../validators/dexes/update');

exports.register = (server, options, next) => {

  server.route([{
    method: 'GET',
    path: '/users/{username}/dexes/{slug}',
    config: {
      handler: (request, reply) => reply(Controller.retrieve(request.params))
    }
  }, {
    method: 'POST',
    path: '/users/{username}/dexes',
    config: {
      auth: 'token',
      handler: (request, reply) => reply(Controller.create(request.params, request.payload, request.auth.credentials)),
      validate: { payload: DexesCreateValidator }
    }
  }, {
    method: 'POST',
    path: '/users/{username}/dexes/{slug}',
    config: {
      auth: 'token',
      handler: (request, reply) => reply(Controller.update(request.params, request.payload, request.auth.credentials)),
      validate: { payload: DexesUpdateValidator }
    }
  }, {
    method: 'DELETE',
    path: '/users/{username}/dexes/{slug}',
    config: {
      auth: 'token',
      handler: (request, reply) => reply(Controller.delete(request.params, request.auth.credentials))
    }
  }]);

  next();

};

exports.register.attributes = {
  name: 'dexes'
};
