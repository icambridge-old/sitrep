'use strict';

/* App Module */

var jobApp = angular.module('jobApp', [
  'ngRoute',
  'jobServices',
  'jobControllers'
]);

jobApp.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider.
      when('/builds', {
        templateUrl: '/public/partials/build-list.html',
        controller: 'JobListCtrl'
      }).
      when('/job/:jobId', {
        templateUrl: '/public/partials/job-details.html',
        controller: 'JobDetailCtrl'
      }).
      otherwise({
        redirectTo: '/builds'
      });
  }]);
