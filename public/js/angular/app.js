"use strict";

var app = angular.module("WhalesApp", ["ngCookies"]).config(["$routeProvider", "$locationProvider", function($routeProvider, $locationProvider) {
        $routeProvider.when("/home", {
            templateUrl: "/public/partials/HomeIndex.html",
            controller: "HomeController"
        });
        $routeProvider.when("/signup", {
            templateUrl: "/public/partials/UserSignup.html",
            controller: "UserSignupController"
        });
        $routeProvider.when("/login", {
            templateUrl: "/public/partials/UserLogin.html",
            controller: "UserLoginController"
        });
        $routeProvider.otherwise({redirectTo: "/home"});
}]);