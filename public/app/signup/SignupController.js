app.controller("SignupController", ["$scope", "$location", "UserService", "FlashService", function($scope, $location, UserService, FlashService) {

    $scope.credentials = {
        email: "test@test.org",
        password: "123456"
    };

    $scope.signup = function() {
        var signup = UserService.signup($scope.credentials);
        signup.success(function() {
            $UserService.login($scope.credentials());
        });
        signup.error(FlashService.showError);
    };

}]);