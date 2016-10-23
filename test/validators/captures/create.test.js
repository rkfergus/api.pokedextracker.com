'use strict';

const Joi = require('joi');

const CapturesCreateValidator = require('../../../src/validators/captures/create');

describe('captures create validator', () => {

  describe('pokemon', () => {

    it('is required', () => {
      const data = {};
      const result = Joi.validate(data, CapturesCreateValidator);

      expect(result.error.details[0].path).to.eql('pokemon');
      expect(result.error.details[0].type).to.eql('any.required');
    });

    it('allows an array of integers', () => {
      const data = { pokemon: [1] };
      const result = Joi.validate(data, CapturesCreateValidator);

      expect(result.error).to.not.exist;
      expect(result.value.pokemon).to.eql([1]);
    });

    it('allows a single integer', () => {
      const data = { pokemon: 1 };
      const result = Joi.validate(data, CapturesCreateValidator);

      expect(result.error).to.not.exist;
      expect(result.value.pokemon).to.eql([1]);
    });

    it('disallows an array of strings', () => {
      const data = { pokemon: ['test'] };
      const result = Joi.validate(data, CapturesCreateValidator);

      expect(result.error.details[0].path).to.eql('pokemon.0');
      expect(result.error.details[0].type).to.eql('number.base');
    });

  });

});
