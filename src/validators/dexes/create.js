'use strict';

const Joi = require('joi');

module.exports = Joi.object().keys({
  title: Joi.string().max(300).trim().required(),
  shiny: Joi.boolean().required(),
  generation: Joi.number().integer().min(1).max(6).required()
});
