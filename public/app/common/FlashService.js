app.factory("FlashService", ["$rootScope", function ($rootScope) {
	return {
		show: function(message) {
			$rootScope.flashClass = "alert";
			$rootScope.message = message;
		},
		showSuccess: function(message) {
			$rootScope.flashClass = "alert alert-success";
			$rootScope.flash = message;
		},
		showError: function(message) {
			$rootScope.flashClass = "alert alert-error";
			$rootScope.flash = message;
		},
		showInfo: function(message) {
			$rootScope.flashClass = "alert alert-info";
			$rootScope.flash = message;
		},
		clear: function() {
			$rootScope.flash = "";
		}
	};
}])