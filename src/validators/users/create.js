'use strict';

const Joi = require('joi');

module.exports = Joi.object().keys({
  username: Joi.string().token().max(20).trim().required(),
  password: Joi.string().min(8).max(72).required(),
  friend_code: Joi.string().regex(/^\d{4}-\d{4}-\d{4}$/).empty(['', null]),
  referrer: Joi.string().empty(['', null]),
  title: Joi.string().max(300).trim().default('Living Dex'),
  shiny: Joi.boolean().default(false),
  generation: Joi.number().integer().min(6).max(7).default(6),
  region: Joi.string().valid(['national', 'alola']).default('national')
})
.options({
  language: {
    string: { regex: { base: 'must be a 12-digit number' } }
  }
});
