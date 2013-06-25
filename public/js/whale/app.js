"use strict";

var app = angular.module("Whale", ["ngCookies"]).config(["$routeProvider", "$locationProvider", function($routeProvider, $locationProvider) {
        $routeProvider.when("/home", {
            templateUrl: "/public/partials/HomeIndex.html",
            controller: "HomeController"
        });
        $routeProvider.when("/signup", {
            templateUrl: "/public/partials/UserSignup.html",
            controller: "UserController"
        });
        $routeProvider.when("/login", {
            templateUrl: "/public/partials/UserLogin.html",
            controller: "UserController"
        });
        $routeProvider.otherwise({redirectTo: "/home"});
}]);