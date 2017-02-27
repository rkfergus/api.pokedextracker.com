'use strict';

const Joi = require('joi');

module.exports = Joi.object({
  dex: Joi.number().integer().required()
});
