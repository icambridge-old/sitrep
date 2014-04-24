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
        controller: 'BuildListCtrl'
      }).
      when('/job/:jobId', {
        templateUrl: '/public/partials/job-details.html',
        controller: 'JobPullRequestsCtrl'
      }).
      when('/job/:jobId/branches', {
        templateUrl: '/public/partials/job-branches.html',
        controller: 'JobBranchesCtrl'
      }).
      when('/build/branches/:branchName', {
        templateUrl: '/public/partials/build-branches.html',
        controller: 'BuildBranchesCtrl'
      }).
      otherwise({
        redirectTo: '/builds'
      });
  }]);
