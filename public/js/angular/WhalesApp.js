var app = angular.module("WhalesApp", []).config(function($routeProvider) {
    $routeProvider.when("/home", {
        templateUrl: "/public/templates/HomeIndex.html",
        controller: "HomeController"
    });
    $routeProvider.when("/signup", {
        templateUrl: "/public/templates/UserSignup.html",
        controller: "SignupController"
    });
    $routeProvider.when("/login", {
        templateUrl: "/public/templates/UserLogin.html",
        controller: "LoginController"
    });
    $routeProvider.otherwise({redirectTo: "/home"});
});

app.controller("HomeController", function() {
    
});

app.controller("SignupController", function() {
    
});

app.controller("LoginController", function() {
    
});