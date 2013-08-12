app.factory("FlashService", ["$rootScope", function ($rootScope) {
	return {
		show: function(message, level, delay) {
			$rootScope.flash = message;
		},
		clear: function() {
			$rootScope.flasg = "";
		}
	};
}])