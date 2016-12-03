'use strict';

const Joi = require('joi');

module.exports = Joi.object().keys({
  title: Joi.string().max(300).trim().required(),
  shiny: Joi.boolean().required(),
  generation: Joi.number().integer().min(6).max(7).required(),
  region: Joi.string().valid(['national', 'alola']).required()
});
