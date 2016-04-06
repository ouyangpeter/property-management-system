/**
 * Created by sep on 16/3/7.
 */

var app = angular.module('loginApp', []);
app.config(function ($interpolateProvider) {
    $interpolateProvider.startSymbol('//');
    $interpolateProvider.endSymbol('//');

});

app.controller('loginFormController', function ($scope, $http) {
    $scope.user = {identifier: '', credential: '', identityType: "username"};
    $scope.loginFailed = false;
    $scope.login = function () {
        $scope.loginUser = angular.copy($scope.user);
        if ($scope.loginUser.identityType == "username"){
            $scope.loginUser.credential= md5("#86" + $scope.loginUser.credential+ "#86").toUpperCase()
        }
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
        if ($scope.user.identifier.$dirty && $scope.user.identifier.$invalid || $scope.isLoginFailed()) {
            return "has-error";
        } else {
            return "";
        }

    };
    $scope.getPasswordExCSS = function () {
        if ($scope.user.credential.$dirty && $scope.user.credential.$invalid || $scope.isLoginFailed()) {
            return "has-error";
        } else {
            return "";
        }
    };


});
