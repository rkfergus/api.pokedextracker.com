'use strict';

const Joi = require('joi');

module.exports = Joi.object({
  pokemon: Joi.array().items(Joi.number().integer()).single().required()
});
