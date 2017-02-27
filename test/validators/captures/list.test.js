'use strict';

const Joi = require('joi');

const CapturesListValidator = require('../../../src/validators/captures/list');

describe('captures list validator', () => {

  describe('dex', () => {

    it('is required', () => {
      const data = {};
      const result = Joi.validate(data, CapturesListValidator);

      expect(result.error.details[0].path).to.eql('dex');
      expect(result.error.details[0].type).to.eql('any.required');
    });

    it('allows integers', () => {
      const data = { dex: 1 };
      const result = Joi.validate(data, CapturesListValidator);

      expect(result.error).to.not.exist;
    });

    it('disallows strings', () => {
      const data = { dex: 'test' };
      const result = Joi.validate(data, CapturesListValidator);

      expect(result.error.details[0].path).to.eql('dex');
      expect(result.error.details[0].type).to.eql('number.base');
    });

  });

});
