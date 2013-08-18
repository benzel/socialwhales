"use strict";

app.factory("UserService", ["$http", "$location", "SessionService",
    function($http, $location, SessionService) {
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
                return $http.post("/auth/login", credentials).success(cacheSession);
            },

            logout: function() {
                var logout = $http.get("/auth/logout");
                logout.success(uncacheSession);
                logout.error(uncacheSession);
                return logout;
            },

            isLoggedIn: function() {
                return SessionService.get("authenticated") && SessionService.get("authenticated") != null;
            }
        };
    }
]);