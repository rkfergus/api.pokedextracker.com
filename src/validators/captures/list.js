'use strict';

const Joi = require('joi');

module.exports = Joi.object({
  user: Joi.number().integer().required()
});
