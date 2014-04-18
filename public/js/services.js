'use strict';

/* Services */

var jobServices = angular.module('jobServices', ['ngResource']);

jobServices.factory('job', ['$resource',
  function($resource){
    return $resource('job/:jobId', {}, {
      query: {method:'GET', params:{jobId:'jobs'}, isArray:true}
    });
  }]);

jobServices.factory('build', ['$resource',
    function($resource){
      return $resource('build/:buildId', {}, {
        query: {method:'GET', params:{buildId:'list'}, isArray:true}
      });
    }]);