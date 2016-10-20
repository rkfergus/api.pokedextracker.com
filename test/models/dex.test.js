'use strict';

const Dex = require('../../src/models/dex');

const dex = Factory.build('dex');

describe('dex model', () => {

  describe('serialize', () => {

    it('returns the correct fields', () => {
      expect(Dex.forge(dex).serialize()).to.have.all.keys([
        'id',
        'user_id',
        'title',
        'slug',
        'shiny',
        'generation',
        'date_created',
        'date_modified'
      ]);
    });

  });

});
