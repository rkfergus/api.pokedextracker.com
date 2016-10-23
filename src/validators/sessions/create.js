'use strict';

const Joi = require('joi');

module.exports = Joi.object().keys({
  username: Joi.string().token().trim().required(),
  password: Joi.string().required()
});
