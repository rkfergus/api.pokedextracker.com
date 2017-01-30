'use strict';

const Joi = require('joi');

module.exports = Joi.object({
  generation: Joi.number().integer(),
  region: Joi.string().valid(['national', 'alola'])
});
