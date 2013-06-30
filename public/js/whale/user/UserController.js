app.controller("UserController", ["$scope", "$location", "UserService", function($scope, $location, UserService) {
    $scope.credentials = {
        email: "test@test.org",
        password: "123456"
    };
    $scope.signup = function() {
        UserService.signup($scope.credentials).success(function() {
			UserService.login($scope.credentials).success(function() {
//				$location.path("home");
			})
		});
    };
    $scope.login = function() {
        UserService.login($scope.credentials);
    };
    $scope.logout = function() {
        UserService.logout($scope.credentials);
    };
}]);