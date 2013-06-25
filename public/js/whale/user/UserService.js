"use strict";

app.factory("UserService", ["$http", "$location", function($http, $location) {
    return {
        signup: function(credentials) {
            return $http.post("/auth/signup", credentials);
        },
        login: function(credentials) {
            return $http.post("/auth/login", credentials);
        },
        logout: function() {
            return $http.post("/auth/logout", credentials);
        }
    };
}]);