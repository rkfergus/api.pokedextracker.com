'use strict';

const Joi = require('joi');

const PokemonRetrieveValidator = require('../../../src/validators/pokemon/retrieve');

describe('pokemon retrieve validator', () => {

  it('does not require all params', () => {
    const data = {};
    const result = Joi.validate(data, PokemonRetrieveValidator);

    expect(result.error).to.not.exist;
  });

  it('allows all params', () => {
    const data = { generation: 6, region: 'national' };
    const result = Joi.validate(data, PokemonRetrieveValidator);

    expect(result.error).to.not.exist;
  });

  describe('generation', () => {

    it('allows integers', () => {
      const data = { generation: 1 };
      const result = Joi.validate(data, PokemonRetrieveValidator);

      expect(result.error).to.not.exist;
    });

    it('disallows strings', () => {
      const data = { generation: 'test' };
      const result = Joi.validate(data, PokemonRetrieveValidator);

      expect(result.error.details[0].path).to.eql('generation');
      expect(result.error.details[0].type).to.eql('number.base');
    });

  });

  describe('region', () => {

    it('allows national', () => {
      const data = { region: 'national' };
      const result = Joi.validate(data, PokemonRetrieveValidator);

      expect(result.error).to.not.exist;
    });

    it('allows alola for gen 7', () => {
      const data = { region: 'alola' };
      const result = Joi.validate(data, PokemonRetrieveValidator);

      expect(result.error).to.not.exist;
    });

    it('disallows other regions', () => {
      const data = { region: 'kalos' };
      const result = Joi.validate(data, PokemonRetrieveValidator);

      expect(result.error.details[0].path).to.eql('region');
      expect(result.error.details[0].type).to.eql('any.allowOnly');
    });

  });

});
