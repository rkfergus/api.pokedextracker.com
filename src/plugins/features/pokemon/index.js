'use strict';

const Controller      = require('./controller');
const ParamsValidator = require('../../../validators/params');

exports.register = (server, options, next) => {

  server.route([{
    method: 'GET',
    path: '/pokemon',
    config: {
      handler: (request, reply) => reply(Controller.list())
    }
  }, {
    method: 'GET',
    path: '/pokemon/{id}',
    config: {
      handler: (request, reply) => reply(Controller.retrieve(request.params.id)),
      validate: { params: ParamsValidator }
    }
  }]);

  next();

};

exports.register.attributes = {
  name: 'pokemon'
};
