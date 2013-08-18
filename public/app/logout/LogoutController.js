/**
  *
  *
  */
app.controller("LogoutController", [
	"$scope", "$location", "FlashService", "UserService",
	function($scope, $location, FlashService, UserService) {

		$scope.timeRemained = 5;
		var timer;
		timer = setInterval(function() {
			if ($scope.timeRemained > 0) {
		        $scope.timeRemained -= 1;
		        $scope.$apply();
		    } else {
		    	clearInterval(timer);
		        var logout = UserService.logout();
		        logout.success(function() {
		            $location.path("/landing");
		        });
		        logout.error(FlashService.showError);
		        logout.error(function() {
		            $location.path("/landing");
		        });
		    }
		 }, 1000);

	}
]);