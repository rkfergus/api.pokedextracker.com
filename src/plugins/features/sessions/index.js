'use strict';

const Controller              = require('./controller');
const SessionsCreateValidator = require('../../../validators/sessions/create');

exports.register = (server, options, next) => {

  server.route([{
    method: 'POST',
    path: '/sessions',
    config: {
      handler: (request, reply) => reply(Controller.create(request.payload, request)),
      validate: { payload: SessionsCreateValidator }
    }
  }]);

  next();

};

exports.register.attributes = {
  name: 'sessions'
};
