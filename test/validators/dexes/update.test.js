'use strict';

const Joi = require('joi');

const DexesUpdateValidator = require('../../../src/validators/dexes/update');

describe('dexes update validator', () => {

  describe('title', () => {

    it('is optional', () => {
      const data = { shiny: false, generation: 6 };
      const result = Joi.validate(data, DexesUpdateValidator);

      expect(result.error).to.not.exist;
    });

    it('limits to 300 characters', () => {
      const data = { title: 'a'.repeat(301), shiny: false, generation: 6 };
      const result = Joi.validate(data, DexesUpdateValidator);

      expect(result.error.details[0].path).to.eql('title');
      expect(result.error.details[0].type).to.eql('string.max');
    });

    it('trims excess whitespace', () => {
      const data = { title: '   a   ', shiny: false, generation: 6 };
      const result = Joi.validate(data, DexesUpdateValidator);

      expect(result.value.title).to.eql('a');
    });

  });

  describe('shiny', () => {

    it('is optional', () => {
      const data = { title: 'Test', generation: 6 };
      const result = Joi.validate(data, DexesUpdateValidator);

      expect(result.error).to.not.exist;
    });

  });

  describe('generation', () => {

    it('is optional', () => {
      const data = { title: 'Test', shiny: false };
      const result = Joi.validate(data, DexesUpdateValidator);

      expect(result.error).to.not.exist;
    });

    it('allows at least 1', () => {
      const data = { title: 'Test', shiny: false, generation: 1 };
      const result = Joi.validate(data, DexesUpdateValidator);

      expect(result.error).to.not.exist;
    });

    it('allows at most 6', () => {
      const data = { title: 'Test', shiny: false, generation: 6 };
      const result = Joi.validate(data, DexesUpdateValidator);

      expect(result.error).to.not.exist;
    });

    it('disallows less than 1', () => {
      const data = { title: 'Test', shiny: false, generation: 0 };
      const result = Joi.validate(data, DexesUpdateValidator);

      expect(result.error.details[0].path).to.eql('generation');
      expect(result.error.details[0].type).to.eql('number.min');
    });

    it('disallows more than 6', () => {
      const data = { title: 'Test', shiny: false, generation: 7 };
      const result = Joi.validate(data, DexesUpdateValidator);

      expect(result.error.details[0].path).to.eql('generation');
      expect(result.error.details[0].type).to.eql('number.max');
    });

  });

});
