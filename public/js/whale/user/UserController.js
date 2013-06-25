app.controller("UserController", ["$scope", "$http", "UserService", function($scope, $http, UserService) {
    $scope.credentials = {
        email: "",
        password: ""
    };
    $scope.signup = function() {
        UserService.signup($scope.credentials);
    };
    $scope.login = function() {
        UserService.login($scope.credentials);
    };
    $scope.logout = function() {
        UserService.logout($scope.credentials);
    };
}]);