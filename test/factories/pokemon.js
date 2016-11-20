'use strict';

module.exports = Factory.define('pokemon')
  .sequence('national_id')
  .sequence('evolution_family_id')
  .attr('name', '')
  .attr('generation', 1);
