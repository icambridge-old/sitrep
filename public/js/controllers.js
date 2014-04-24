'use strict';

/* Controllers */

var jobControllers = angular.module('jobControllers', []);

jobControllers.controller('BuildListCtrl', ['$scope', 'build',
  function($scope, build) {
    $scope.builds = build.query();
  }]);

jobControllers.controller('JobPullRequestsCtrl', ['$scope', '$routeParams', 'job',
  function($scope, $routeParams, job) {
      $scope.jobId = $routeParams.jobId
      $scope.job = job.get({jobId: $routeParams.jobId}, function(job) {});

  }]);
jobControllers.controller('JobBranchesCtrl', ['$scope', '$routeParams', 'branches',
  function($scope, $routeParams, branches) {
      $scope.jobId = $routeParams.jobId
      $scope.job = branches.get({jobId: $routeParams.jobId}, function(job) {});

  }]);
jobControllers.controller('BuildBranchesCtrl', ['$scope', '$routeParams', 'buildBranches',
    function($scope, $routeParams, builds) {
        $scope.branchName = $routeParams.branchName
        $scope.response = builds.get({branchName: $routeParams.branchName}, function(job) {});

    }]);