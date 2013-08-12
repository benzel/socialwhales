"use strict";

app.factory("UserService", ["$http", "$location", "SessionService", function($http, $location, SessionService) {
    var cacheSession = function() {
        SessionService.set("authenticated", true);
    };
    var uncacheSession = function() {
        SessionService.unset("authenticated");
    };

    return {
        signup: function(credentials) {
            return $http.post("/auth/signup", credentials);
        },

        login: function(credentials) {
            return $http.post("/auth/login", credentials).success(cacheSession);;
        },

        logout: function() {
            return $http.get("/auth/logout", credentials).success(uncacheSession);;
        },

        isLogedIn: function() {
            return SessionService.get("authenticated")
        }
    };
}]);