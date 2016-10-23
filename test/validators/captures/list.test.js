'use strict';

const Joi = require('joi');

const CapturesListValidator = require('../../../src/validators/captures/list');

describe('captures list validator', () => {

  it('requires either dex or user', () => {
    const data = {};
    const result = Joi.validate(data, CapturesListValidator);

    expect(result.error.details[0].path).to.eql('value');
    expect(result.error.details[0].type).to.eql('object.missing');
  });

  it('disallows both dex and user', () => {
    const data = { dex: 1, user: 1 };
    const result = Joi.validate(data, CapturesListValidator);

    expect(result.error.details[0].path).to.eql('value');
    expect(result.error.details[0].type).to.eql('object.xor');
  });

  describe('dex', () => {

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

  describe('user', () => {

    it('allows integers', () => {
      const data = { user: 1 };
      const result = Joi.validate(data, CapturesListValidator);

      expect(result.error).to.not.exist;
    });

    it('disallows strings', () => {
      const data = { user: 'test' };
      const result = Joi.validate(data, CapturesListValidator);

      expect(result.error.details[0].path).to.eql('user');
      expect(result.error.details[0].type).to.eql('number.base');
    });

  });

});
