// Function to use X,Y coordinates instead of Y,X
  var yx = L.latLng;
  var xy = function(x, y) {
    if (L.Util.isArray(x)) {    // When doing xy([x, y]);
        return yx(-x[1], x[0]);
    }
    return yx(-y, x);  // -y because it's up to down and no down to up like by default
  };


// Zoom reinitializer
  $("#full-extent-btn").click(function() {
    map.fitBounds([xy(0,16410), xy(16410,0)]);
    $(".navbar-collapse.in").collapse("hide");
    return false;
  });


// Map initializer
  var map = L.map('map', {
    crs: L.CRS.Simple,
    minZoom: -4.6,
    wheelPxPerZoomLevel: 150,
    //layers: [cmarkers1, cmarkers2],
    zoomControl: false,
    attributionControl: false
  });
  var bounds = [xy(0,16410), xy(16410,0)];
  var image = L.imageOverlay('/images/map_tree_of_life.jpg', bounds).addTo(map);
  map.fitBounds(bounds);


// Mouse Position plugin
  var mouse = L.control.mousePosition({
    position: 'bottomleft',
    separator: ' , ',
    emptyString: 'Positionnez la souris...',
    lngFirst: true,
    lngFormatter: function (longitude) { return Math.floor(L.Util.formatNum(longitude, this.numDigits)) },
    latFormatter: function (latitude) { return Math.floor(L.Util.formatNum(-latitude, this.numDigits)) },
    }).addTo(map);


// Leaflet patch to make layer control scrollable on touch browsers
  var container = $(".leaflet-control-layers")[0];
  if (!L.Browser.touch) {
    L.DomEvent
    .disableClickPropagation(container)
    .disableScrollPropagation(container);
  } else {
    L.DomEvent.disableClickPropagation(container);
  }
