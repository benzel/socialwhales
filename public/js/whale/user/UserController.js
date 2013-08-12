app.controller("UserController", ["$scope", "$location", "UserService", "FlashService", function($scope, $location, UserService, FlashService) {

    var showSuccess = function(response) {
        FlashService.show(response.flash, "error");
    }
    var showError = function(response) {
        FlashService.show(response.flash, "success");
    }
    $scope.credentials = {
        email: "test@test.org",
        password: "123456"
    };

    $scope.signup = function() {
        var signup = UserService.signup($scope.credentials);
        signup.success(function() {
            $UserService.login($scope.credentials());
        });
        signup.success(showSuccess)
        signup.error(showError);
    };
    $scope.login = function() {
        UserService.login($scope.credentials).success(function() {
            $location.path("/home");
        });
    };
    $scope.logout = function() {
        UserService.logout($scope.credentials);
        $location.path("/home");
    };
}]);