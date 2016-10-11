'use strict';

const Joi = require('joi');

const UserListValidator = require('../../../src/validators/user/list');

describe('user list validator', () => {

  describe('limit', () => {

    it('requires integers', () => {
      const data = { limit: 1.1 };
      const result = Joi.validate(data, UserListValidator);

      expect(result.error.details[0].path).to.eql('limit');
      expect(result.error.details[0].type).to.eql('number.integer');
    });

    it('has a minimum of 0', () => {
      const data = { limit: -1 };
      const result = Joi.validate(data, UserListValidator);

      expect(result.error.details[0].path).to.eql('limit');
      expect(result.error.details[0].type).to.eql('number.min');
    });

    it('has a maximum of 100', () => {
      const data = { limit: 101 };
      const result = Joi.validate(data, UserListValidator);

      expect(result.error.details[0].path).to.eql('limit');
      expect(result.error.details[0].type).to.eql('number.max');
    });

    it('has a default of 10', () => {
      const data = {};
      const result = Joi.validate(data, UserListValidator);

      expect(result.error).to.not.exist;
      expect(result.value.limit).to.eql(10);
    });

  });

  describe('offset', () => {

    it('requires integers', () => {
      const data = { offset: 1.1 };
      const result = Joi.validate(data, UserListValidator);

      expect(result.error.details[0].path).to.eql('offset');
      expect(result.error.details[0].type).to.eql('number.integer');
    });

    it('has a minimum of 0', () => {
      const data = { offset: -1 };
      const result = Joi.validate(data, UserListValidator);

      expect(result.error.details[0].path).to.eql('offset');
      expect(result.error.details[0].type).to.eql('number.min');
    });

    it('has a default of 0', () => {
      const data = {};
      const result = Joi.validate(data, UserListValidator);

      expect(result.error).to.not.exist;
      expect(result.value.offset).to.eql(0);
    });

  });

});
