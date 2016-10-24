'use strict';

const Joi = require('joi');

module.exports = Joi.object({
  dex: Joi.number().integer(),
  user: Joi.number().integer()
})
.xor(['dex', 'user']);
