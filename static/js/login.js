/**
 * Created by sep on 16/3/7.
 */

var app = angular.module('loginApp', []);
app.config(function ($interpolateProvider) {
    $interpolateProvider.startSymbol('//');
    $interpolateProvider.endSymbol('//');

});

app.controller('loginFormController', function ($scope, $http) {
    $scope.user = {'userName': '', 'password': ''};
    $scope.loginFailed = false;
    $scope.login = function () {
        $scope.loginUser = angular.copy($scope.user);
        $scope.loginUser.password = faultylabs.MD5($scope.loginUser.password).toUpperCase()
        $http.post("/login", $scope.loginUser).success(function () {
            window.location.href = "/";
        }).error(function () {
            $scope.loginFailed = true;
            $scope.preUser = angular.copy($scope.user);
        });
    };
    $scope.isLoginFailed = function () {
        if ($scope.loginFailed) {
            return true;
        } else {
            return false;
        }

    };

    $scope.getUserNameExCSS = function () {
        if ($scope.user.userName.$dirty && $scope.user.userName.$invalid || $scope.isLoginFailed()) {
            return "has-error";
        } else {
            return "";
        }

    };
    $scope.getPasswordExCSS = function () {
        if ($scope.user.password.$dirty && $scope.user.password.$invalid || $scope.isLoginFailed()) {
            return "has-error";
        } else {
            return "";
        }
    };


});
