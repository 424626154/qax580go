var beermap;

var makers;

function initMap() {
	beermap = new AMap.Map('beermap', {
		resizeEnable: true,
		zoom: 10
	});
	makers = new Array();
	getMakers();
}

function getMakers() {
	var url = "http://localhost:8080/beermap/ajax";
	var data = "op=getmakers";
	$.ajax({
		type: "POST",
		url: url,
		data: data,
		dataType: 'json',
		success: function(obj) {
			// alert(obj.errcode)
			if (obj.errcode == 0) {
				// console.log(obj.data)
				addMakers(obj.data)
			} else {
				alert(obj.errmsg)
			}
		}
	});
}

function addMakers(makes_str) {
	var makes = JSON.parse(makes_str);
	if (makes.length > 0) {
		if (beermap != null) {
			for (var i = 0; i < makes.length; i++) {
				var maker_data = makes[i];
				// console.log(maker_data.Name)
				var bIcon = new AMap.Icon({
					size: new AMap.Size(30, 30), //图标大小
					image: "/static/img/beermap/beer_maker_type1.png",
					imageSize: new AMap.Size(30, 30)
				})
				var location = [maker_data.Lng, maker_data.Lat]
				var marker = new AMap.Marker({ //加点
					map: beermap,
					position: location,
					icon: bIcon,
					extData: {
						marker_data: maker_data
					}
				});
				marker.on('click', function(e) {
					var maker_data = e.target.getExtData()['marker_data'];
					// console.log(maker_data);
					openInfo(maker_data)
				});
				makers.push(marker);
			}
		}
	}
}


function openInfo(marker_data) {
	//构建信息窗体中显示的内容
	var info = [];
	info.push("<div>" + marker_data.Name + "<div>");
	info.push("<div>" + marker_data.Describe + "<div>");
	infoWindow = new AMap.InfoWindow({
		content: info.join("<br>"), //使用默认信息窗体框样式，显示信息内容
		offset: new AMap.Pixel(0, -30)
	});
	infoWindow.open(beermap, [marker_data.Lng, marker_data.Lat]);
}