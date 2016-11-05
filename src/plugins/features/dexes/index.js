'use strict';

const Controller = require('./controller');

exports.register = (server, options, next) => {

  server.route([{
    method: 'GET',
    path: '/users/{username}/dexes/{slug}',
    config: {
      handler: (request, reply) => reply(Controller.retrieve(request.params))
    }
  }]);

  next();

};

exports.register.attributes = {
  name: 'dexes'
};
