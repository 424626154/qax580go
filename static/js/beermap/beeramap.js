var beermap;
// var placeSearch;
var citycode;

var searchResult;

var makers;

var selectpoi;

var server_makers;

function initMap() {
	beermap = new AMap.Map('beermap', {
		resizeEnable: true,
		zoom: 10
	});
	beermap.getCity(function(result) {
		citycode = result.citycode
	})
	initWidget();
	initAuto();
	regeocoder();
	// server_makers = new Array();
	// getServerMakers();
}

function initWidget() {
	AMap.plugin(['AMap.ToolBar', 'AMap.Scale', 'AMap.OverView', 'AMap.Geolocation'],
		function() {
			beermap.addControl(new AMap.ToolBar());

			beermap.addControl(new AMap.Scale());

			beermap.addControl(new AMap.OverView({
				isOpen: true
			}));
		});
}

function initSearch() {
	AMap.service('AMap.PlaceSearch', function() { //回调函数
		//实例化PlaceSearch
		placeSearch = new AMap.PlaceSearch({
			city: citycode, //城市
			map: beermap
		});
		//TODO: 使用placeSearch对象调用关键字搜索的功能
		//关键字查询
		placeSearch.search('北京', function(status, result) {
			if (status == 'complete' && result.info === 'OK') {
				alert(result.poiList.pois)
			} else {
				alert(result)
			}
		});
	})
}

function initAuto() {
	AMap.plugin(['AMap.Autocomplete', 'AMap.PlaceSearch'], function() { //回调函数
		//实例化Autocomplete
		var autoOptions = {
			city: "", //城市，默认全国
			input: "keyword" //使用联想输入的input的id
		};
		autocomplete = new AMap.Autocomplete(autoOptions);
		//TODO: 使用autocomplete对象调用相关功能
		var placeSearch = new AMap.PlaceSearch({
			city: '北京',
			// map: beermap
		});
		AMap.event.addListener(autocomplete, "select", function(e) {
			//TODO 针对选中的poi实现自己的功能
			// e.poi.name
			console.log(e.poi.name)
			placeSearch.search(e.poi.name, function(status, result) {
				if (status == 'complete' && result.info === 'OK') {
					// console.log(result)
					searchResult = result
						// console.log(beermap.getAllOverlays())
					this.addSearchMaker();
				} else {
					console.log(result)
				}
			});
		});
	})
}

//解析定位结果
function onComplete(data) {
	// beermap.setCenter([data.position.getLng(), data.position.getLat()]);
	console.log(data);
}
//解析定位错误信息
function onError(data) {
	console.log('定位失败');
}


function regeocoder() { //逆地理编码
	AMap.plugin('AMap.Geocoder', function() {
		lnglatXY = [beermap.getCenter().getLng(), beermap.getCenter().getLat()]; //已知点坐标
		var geocoder = new AMap.Geocoder({
			radius: 1000,
			extensions: "all"
		});
		geocoder.getAddress(lnglatXY, function(status, result) {
			if (status === 'complete' && result.info === 'OK') {
				// console.log(result)
			}
		});
	})
}

function getMarks() {
	beermap.getAllOverlays('marker')
	console.log(beermap.getAllOverlays())
}



function addSearchMaker() {
	if (searchResult != null) {
		if (makers != null) {
			for (var i = 0; i < makers.length; i++) {
				makers[i].setMap(null)
			}
		}
		makers = new Array();
		for (var h = 0; h < searchResult.poiList.pois.length; h++) { //返回搜索列表循环绑定marker
			var poi = searchResult.poiList.pois[h];
			var location = poi['location']; //经纬度
			// console.log(location)
			// 	var address = result.poiList.pois[h]['address']; //地址
			// var isAdd = false;
			// // console.log(poi.id)
			// if (server_makers != null) {
			// 	for (var j = 0; j < server_makers.length; j++) {
			// 		if (poi.id == server_makers[j].getExtData()['marker_data'].MId) {
			// 			console.log(server_makers[j].getExtData()['marker_data'].MId)
			// 			isAdd = true;
			// 			break;
			// 		}
			// 	}
			// }
			// poi.isAdd = isAdd;
			var image = "/static/img/beermap/beer_maker_type0.png";
			// if (poi.isAdd) {
			// 	image = "/static/img/beermap/beer_maker_type1.png";
			// }
			var bIcon = new AMap.Icon({
				size: new AMap.Size(30, 30), //图标大小
				image: image,
				imageSize: new AMap.Size(30, 30)
			})

			var marker = new AMap.Marker({ //加点
				map: beermap,
				position: location,
				icon: bIcon,
				extData: {
					poi: poi
				}
			});
			marker.on('click', function(e) {
				var poi = e.target.getExtData()['poi'];
				// console.log(poi.name);
				openInfo(poi);
			});
			makers.push(marker);

		}
	}
}

//构建信息窗体中显示的内容
function openInfo(poi) {
	// console.log(poi.isAdd)
	//构建信息窗体中显示的内容
	$('#save_bg').show();
	console.log($('#save_title'))
	$("#save_title").html(poi.name);
	$('#save_lng').html("纬度:" + poi.location.getLng())
	$('#save_lat').html("纬度:" + poi.location.getLat())
	$('#save_text').val('')
	selectpoi = poi;
	$('#save').click(function() {
		// console.log(poi)
		console.log("save_text" + $('#save_text').val())
		var describe = $('#save_text').val();
		addMaker(poi.id, poi.citycode, poi.name, poi.location.getLng(), poi.location.getLat(), describe)
	})
}

function save() {
	alert("save");
}

function closeInfo() {
	$('#save_bg').hide();
	// alert("closeInfo");
}


function addMaker(id, citycode, name, lng, lat, describe) {
	var url = "http://localhost:8080/beermap/ajax";
	var data = "op=addmaker&id=" + id + "&name=" + name + "&lng=" + lng + "&lat=" + lat + "&describe=" + describe;
	console.log(url)
	$.ajax({
		type: "POST",
		url: url,
		data: data,
		dataType: 'json',
		success: function(obj) {
			// alert(obj.errcode)
			if (obj.errcode == 0) {
				addMarkSuccess(obj.data)
				closeInfo()
			} else {
				alert(obj.errmsg)
			}
		}
	});
}

function addMarkSuccess(makes_data) {

}
// function getServerMakers() {
// 	var url = "http://localhost:8080/beermap/ajax";
// 	var data = "op=getmakers";
// 	$.ajax({
// 		type: "POST",
// 		url: url,
// 		data: data,
// 		dataType: 'json',
// 		success: function(obj) {
// 			// alert(obj.errcode)
// 			if (obj.errcode == 0) {
// 				// console.log(obj.data)
// 				addServerMakers(obj.data)
// 			} else {
// 				alert(obj.errmsg)
// 			}
// 		}
// 	});
// }

// function addServerMakers(makers_str) {
// 	var makes = JSON.parse(makers_str);
// 	if (makes.length > 0) {
// 		if (beermap != null) {
// 			for (var i = 0; i < makes.length; i++) {
// 				var maker_data = makes[i];
// 				// console.log(maker_data.Name)
// 				var bIcon = new AMap.Icon({
// 					size: new AMap.Size(30, 30), //图标大小
// 					image: "/static/img/beermap/beer_maker_type1.png",
// 					imageSize: new AMap.Size(30, 30)
// 				})
// 				var location = [maker_data.Lng, maker_data.Lat]
// 				var marker = new AMap.Marker({ //加点
// 					map: beermap,
// 					position: location,
// 					icon: bIcon,
// 					extData: {
// 						marker_data: maker_data
// 					}
// 				});
// 				marker.on('click', function(e) {
// 					var maker_data = e.target.getExtData()['marker_data'];
// 					// console.log(maker_data);
// 					openInfo(maker_data)
// 				});
// 				server_makers.push(marker);
// 			}
// 		}
// 	}
// }

// function addServerMaker(maker_str) {
// 	var maker = JSON.parse(maker_str);
// 	if (beermap != null) {
// 		var maker_data = maker;
// 		// console.log(maker_data.Name)
// 		var bIcon = new AMap.Icon({
// 			size: new AMap.Size(30, 30), //图标大小
// 			image: "/static/img/beermap/beer_maker_type1.png",
// 			imageSize: new AMap.Size(30, 30)
// 		})
// 		var location = [maker_data.Lng, maker_data.Lat]
// 		var marker = new AMap.Marker({ //加点
// 			map: beermap,
// 			position: location,
// 			icon: bIcon,
// 			extData: {
// 				marker_data: maker_data
// 			}
// 		});
// 		marker.on('click', function(e) {
// 			var maker_data = e.target.getExtData()['marker_data'];
// 			// console.log(maker_data);
// 			openInfo(maker_data)
// 		});
// 		server_makers.push(marker);
// 	}
// }