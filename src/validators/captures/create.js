'use strict';

const Joi = require('joi');

module.exports = Joi.object({
  dex: Joi.number().integer(),
  pokemon: Joi.array().items(Joi.number().integer()).single().required()
});
