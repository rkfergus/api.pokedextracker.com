'use strict';

module.exports = Factory.define('pokemon')
  .sequence('national_id')
  .sequence('evolution_family_id')
  .attr('name', '')
  .attr('type1', 'fire')
  .attr('icon_url', '')
  .attr('avatar_url', '')
  .attr('generation', 1)
  .attr('x_location', '')
  .attr('y_location', '')
  .attr('or_location', '')
  .attr('as_location', '');
