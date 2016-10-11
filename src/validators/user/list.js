'use strict';

const Joi = require('joi');

const LIMIT_DEFAULT  = 10;
const LIMIT_MAX      = 100;
const LIMIT_MIN      = 0;
const OFFSET_DEFAULT = 0;
const OFFSET_MIN     = 0;

module.exports = Joi.object().keys({
  limit: Joi.number().integer().min(LIMIT_MIN).max(LIMIT_MAX).default(LIMIT_DEFAULT),
  offset: Joi.number().integer().min(OFFSET_MIN).default(OFFSET_DEFAULT)
});
