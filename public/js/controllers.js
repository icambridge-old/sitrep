'use strict';

/* Controllers */

var jobControllers = angular.module('jobControllers', []);

jobControllers.controller('JobListCtrl', ['$scope', 'build',
  function($scope, build) {
    $scope.builds = build.query();
  }]);

jobControllers.controller('JobDetailCtrl', ['$scope', '$routeParams', 'job',
  function($scope, $routeParams, job) {
    $scope.pullRequests = job.get({jobId: $routeParams.jobId});

  }]);