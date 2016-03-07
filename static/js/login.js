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
        if ($scope.loginFailed && $scope.preUser.userName == $scope.user.userName
            && $scope.preUser.password == $scope.user.password) {
            return true;
        } else {
            return false;
        }

    }

});
