app.controller("LoginController", ["$scope", "$location", "UserService", "FlashService", function($scope, $location, UserService, FlashService) {
    $scope.credentials = {
        email: "test2@test.org",
        password: "123456"
    };

    $scope.login = function() {
        var login = UserService.login($scope.credentials);
        login.success(function() {
            $location.path("/profile");
        });
        login.error(FlashService.showError);
    };

}]);