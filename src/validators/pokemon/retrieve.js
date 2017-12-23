'use strict';

const Joi = require('joi');

module.exports = Joi.object({
  generation: Joi.number().integer(),
  game_family: Joi.string().max(50).trim(),
  region: Joi.string().valid(['national', 'alola']),
  regional: Joi.when('game_family', {
    is: Joi.required(),
    then: Joi.boolean(),
    otherwise: Joi.forbidden()
  })
});
