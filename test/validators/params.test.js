'use strict';

const Joi = require('joi');

const ParamsValidator = require('../../src/validators/params');

describe('params validator', () => {

  describe('id', () => {

    it('is required', () => {
      const data = {};
      const result = Joi.validate(data, ParamsValidator);

      expect(result.error.details[0].path).to.eql('id');
      expect(result.error.details[0].type).to.eql('any.required');
    });

    it('allows integers', () => {
      const data = { id: 1 };
      const result = Joi.validate(data, ParamsValidator);

      expect(result.error).to.not.exist;
    });

    it('disallows strings', () => {
      const data = { id: 'test' };
      const result = Joi.validate(data, ParamsValidator);

      expect(result.error.details[0].path).to.eql('id');
      expect(result.error.details[0].type).to.eql('number.base');
    });

  });

});
