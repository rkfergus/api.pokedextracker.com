'use strict';

const Joi = require('joi');

const UsersCreateValidator = require('../../../src/validators/users/create');

describe('users create validator', () => {

  describe('username', () => {

    it('allows alpha-numeric and underscore characters', () => {
      const data = { username: 'test_TEST', password: 'testtest' };
      const result = Joi.validate(data, UsersCreateValidator);

      expect(result.error).to.not.exist;
    });

    it('disallows anything besides alpha-numeric and underscore characters', () => {
      const data = { username: 'test-TEST', password: 'testtest' };
      const result = Joi.validate(data, UsersCreateValidator);

      expect(result.error.details[0].path).to.eql('username');
      expect(result.error.details[0].type).to.eql('string.token');
    });

    it('limits to 20 characters', () => {
      const data = { username: 'a'.repeat(21), password: 'testtest' };
      const result = Joi.validate(data, UsersCreateValidator);

      expect(result.error.details[0].path).to.eql('username');
      expect(result.error.details[0].type).to.eql('string.max');
    });

    it('trims whitespace', () => {
      const username = '  testing ';
      const data = { username, password: 'testtest' };
      const result = Joi.validate(data, UsersCreateValidator);

      expect(result.value.username).to.eql(username.trim());
    });

  });

  describe('password', () => {

    it('requires at least 8 characters', () => {
      const data = { username: 'testing', password: 'a'.repeat(7) };
      const result = Joi.validate(data, UsersCreateValidator);

      expect(result.error.details[0].path).to.eql('password');
      expect(result.error.details[0].type).to.eql('string.min');
    });

    it('limits to 72 characters', () => {
      const data = { username: 'testing', password: 'a'.repeat(73) };
      const result = Joi.validate(data, UsersCreateValidator);

      expect(result.error.details[0].path).to.eql('password');
      expect(result.error.details[0].type).to.eql('string.max');
    });

  });

  describe('friend_code', () => {

    it('converts null to undefined', () => {
      const data = { username: 'testing', password: 'testtest', friend_code: null };
      const result = Joi.validate(data, UsersCreateValidator);

      expect(result.value.friend_code).to.be.undefined;
    });

    it('converts the empty string to undefined', () => {
      const data = { username: 'testing', password: 'testtest', friend_code: '' };
      const result = Joi.validate(data, UsersCreateValidator);

      expect(result.value.friend_code).to.be.undefined;
    });

    it('allows codes in the format of 1234-1234-1234', () => {
      const data = { username: 'testing', password: 'testtest', friend_code: '1234-1234-1234' };
      const result = Joi.validate(data, UsersCreateValidator);

      expect(result.error).to.not.exist;
    });

    it('disallows codes not in the format of 1234-1234-1234', () => {
      const data = { username: 'testing', password: 'testtest', friend_code: '234-1234-1234' };
      const result = Joi.validate(data, UsersCreateValidator);

      expect(result.error.details[0].path).to.eql('friend_code');
      expect(result.error.details[0].type).to.eql('string.regex.base');
      expect(result.error).to.match(/"friend_code" must be a 12-digit number/);
    });

  });

  describe('referrer', () => {

    it('is optional', () => {
      const data = { username: 'testing', password: 'testtest', friend_code: '1234-1234-1234' };
      const result = Joi.validate(data, UsersCreateValidator);

      expect(result.error).to.not.exist;
    });

    it('can be a string', () => {
      const data = { username: 'testing', password: 'testtest', friend_code: '1234-1234-1234', referrer: 'http://test.com' };
      const result = Joi.validate(data, UsersCreateValidator);

      expect(result.error).to.not.exist;
      expect(result.value.referrer).to.eql(data.referrer);
    });

    it('converts null to undefined', () => {
      const data = { username: 'testing', password: 'testtest', friend_code: '1234-1234-1234', referrer: null };
      const result = Joi.validate(data, UsersCreateValidator);

      expect(result.value.referrer).to.be.undefined;
    });

    it('converts the empty string to undefined', () => {
      const data = { username: 'testing', password: 'testtest', friend_code: '1234-1234-1234', referrer: '' };
      const result = Joi.validate(data, UsersCreateValidator);

      expect(result.value.referrer).to.be.undefined;
    });

  });

});
