'use strict';

const Joi = require('joi');

module.exports = Joi.object().keys({
  title: Joi.string().max(300).trim(),
  shiny: Joi.boolean(),
  generation: Joi.number().integer().min(6).max(7)
});
