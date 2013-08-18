 "use strict";

var app = angular.module("Whale", ["ngCookies"]).config([
    "$routeProvider", "$locationProvider",
    function($routeProvider, $locationProvider) {

        $routeProvider.when("/landing", {
            templateUrl: "/public/app/landing/LandingIndex.html",
            controller: "LandingController"
        });

        $routeProvider.when("/signup", {
            templateUrl: "/public/app/signup/SignupIndex.html",
            controller: "SignupController"
        });

        $routeProvider.when("/login", {
            templateUrl: "/public/app/login/LoginIndex.html",
            controller: "LoginController"
        });

        $routeProvider.when("/logout", {
            templateUrl: "/public/app/logout/LogoutIndex.html",
            controller: "LogoutController"
        });


        $routeProvider.when("/profile", {
            templateUrl: "/public/app/profile/ProfileIndex.html",
            controller: "ProfileController"
        });


        $routeProvider.when("/photo", {
            templateUrl: "/public/app/photo/PhotoIndex.html",
            controller: "PhotoController"
        });

        $routeProvider.when("/photo", {
            templateUrl: "/public/app/photo/PhotoIndex.html",
            controller: "PhotoController"
        });

        $routeProvider.when("/whale", {
            templateUrl: "/public/app/whale/WhaleIndex.html",
            controller: "WhaleController"
        });

        // if($rootScope.UserService.isLoggedIn()) {
            $routeProvider.otherwise({redirectTo: "/landing"});
        // } else {
            // $routeProvider.otherwise({redirectTo: "/profile"});
        // }
}]);

app.run([
    "$rootScope", "$location", "UserService",
    function($rootScope, $location, UserService) {
    $rootScope.UserService = UserService;
    var routesThatRequireAuth = [
        "/logout",
    ];
    var routesThatRequireNoAuth = [
        "/login",
        "/signup",
    ];
    $rootScope.$on("$routeChangeStart", function(event, next, ccurrent) {
        if(_(routesThatRequireAuth).contains($location.path()) && !UserService.isLoggedIn()) {
            $location.path("/login")
        }
        if(_(routesThatRequireNoAuth).contains($location.path()) && UserService.isLoggedIn()) {
            $location.path("/profile")
        }
    });
}]);