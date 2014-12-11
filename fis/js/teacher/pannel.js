function TeacherController($scope,$http){

	var current_page = 1;
	$scope.data = {}
	$scope.data.has_next = false;

	var get_page = function (page) {
		$http.get("/api/problem/list/"+page, {
			params: {"page":page}
		}).success(function(data){
			$scope.data = data;
		});
	}

	get_page(current_page);

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

};

