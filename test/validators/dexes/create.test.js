'use strict';

const Joi = require('joi');

const DexesCreateValidator = require('../../../src/validators/dexes/create');

describe('dexes create validator', () => {

  describe('title', () => {

    it('is required', () => {
      const data = { shiny: false, generation: 6 };
      const result = Joi.validate(data, DexesCreateValidator);

      expect(result.error.details[0].path).to.eql('title');
      expect(result.error.details[0].type).to.eql('any.required');
    });

    it('limits to 300 characters', () => {
      const data = { title: 'a'.repeat(301), shiny: false, generation: 6 };
      const result = Joi.validate(data, DexesCreateValidator);

      expect(result.error.details[0].path).to.eql('title');
      expect(result.error.details[0].type).to.eql('string.max');
    });

    it('trims excess whitespace', () => {
      const data = { title: '   a   ', shiny: false, generation: 6 };
      const result = Joi.validate(data, DexesCreateValidator);

      expect(result.value.title).to.eql('a');
    });

  });

  describe('shiny', () => {

    it('is required', () => {
      const data = { title: 'Test', generation: 6 };
      const result = Joi.validate(data, DexesCreateValidator);

      expect(result.error.details[0].path).to.eql('shiny');
      expect(result.error.details[0].type).to.eql('any.required');
    });

  });

  describe('generation', () => {

    it('is required', () => {
      const data = { title: 'Test', shiny: false };
      const result = Joi.validate(data, DexesCreateValidator);

      expect(result.error.details[0].path).to.eql('generation');
      expect(result.error.details[0].type).to.eql('any.required');
    });

    it('allows at least 6', () => {
      const data = { title: 'Test', shiny: false, generation: 6 };
      const result = Joi.validate(data, DexesCreateValidator);

      expect(result.error).to.not.exist;
    });

    it('allows at most 7', () => {
      const data = { title: 'Test', shiny: false, generation: 7 };
      const result = Joi.validate(data, DexesCreateValidator);

      expect(result.error).to.not.exist;
    });

    it('disallows less than 6', () => {
      const data = { title: 'Test', shiny: false, generation: 5 };
      const result = Joi.validate(data, DexesCreateValidator);

      expect(result.error.details[0].path).to.eql('generation');
      expect(result.error.details[0].type).to.eql('number.min');
    });

    it('disallows more than 7', () => {
      const data = { title: 'Test', shiny: false, generation: 8 };
      const result = Joi.validate(data, DexesCreateValidator);

      expect(result.error.details[0].path).to.eql('generation');
      expect(result.error.details[0].type).to.eql('number.max');
    });

  });

});
