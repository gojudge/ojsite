var tpApp = angular.module('tpApp',['ngRoute']);

tpApp.factory('Data', function () {
	return {};
})

tpApp.config(function($routeProvider) {
  $routeProvider
  	.when('/', {
  		// controller:'TeacherCtrl',
  		templateUrl:'/static/ng/problem_list.html'
  	})
    .when('/problem/add', {
		// controller:'TestCtrl',
		templateUrl:'/static/ng/problem_add.html'
    })
})

tpApp.directive("delete",function($document,$http){
  return{
    restrict:'A',
    require: 'ngModel',
    link:function(scope, element, attrs,ngModel){
      element.bind("click",function(){
        var id = ngModel.$modelValue.id;
        
        scope.$apply(function(){
          for(var i=0; i<scope.data.list.length; i++){
            if(scope.data.list[i].id==id){
            	scope.data.list.splice(i,1);

                $http.get("/api/problem/delete/"+id, {
					params: {"id":id}
				}).success(function(data){
					if (data.result) {
						console.log("delete success.");
						// TeacherCtrl.hello();
					} else{
						console.log("delete failed.", data.debug)
					};
				});
            }
          }
        })
      })
    }
  }
});

tpApp.controller("TeacherCtrl", function($scope,$http,Data) {
	var current_page = 1;
	$scope.data = Data;
	$scope.data.has_next = false;

	this.get_page = function (page) {
		$http.get("/api/problem/list/"+page, {
			params: {"page":page}
		}).success(function(data){
			$scope.data = data;
		});
	}

	this.get_page(current_page);

    $scope.prevPage = function() { 
		get_page(current_page--)
		console.log("ng clicked prevPage:",current_page)
	};

	$scope.nextPage = function() {
		get_page(current_page++)
		console.log("ng clicked nextPage:",current_page)
	};

	$scope.prevPageDisabled = function() {
		return current_page === 1 ? "disabled" : "";
	};

	$scope.nextPageDisabled = function() {
		return !$scope.data.has_next ? "disabled" : "";
	}

});

tpApp.controller("TestCtrl", function($scope,$http,Data) {
	console.log("hello world. TestCtrl.")
});