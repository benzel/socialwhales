"use strict";

app.service("UserService", [function() {
    var sdo = {
        isLogged: false,
        username: ""
    };
    return sdo;
}]);