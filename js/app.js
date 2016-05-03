// Modal control
  $("#about-btn").click(function() {
    $("#aboutModal").modal("show");
    $(".navbar-collapse.in").collapse("hide");
    return false;
  });


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

// Custom markers (for testing)
  var marker1 = L.marker(xy(3361,8872), {icon: L.divIcon({className: 'icon-location blue'})}).bindPopup('<b>Titre un peu long parce qu\'il faut tester</b><br><em>Campement LCH</em><br><a href="#">supprimer</a> | <a href="#">modifier</a>');
  var marker2 = L.marker(xy(7958,14332), {icon: L.divIcon({className: 'icon-location blue'})}).bindPopup('<b>Titre</b><br><em>Campement LCH</em>');
  var marker3 = L.marker(xy(3568,7801), {icon: L.divIcon({className: 'icon-location blue'})}).bindPopup('<b>Titre</b><br><em>Campement LCH</em>');

  var marker4 = L.marker(xy(5120,6185), {icon: L.divIcon({className: 'icon-home green'})}).bindPopup('<b>Union</b><br><em>Campement principal</em>');
  var marker5 = L.marker(xy(3376,8633), {icon: L.divIcon({className: 'icon-home green'})}).bindPopup('<b>Deutschland</b><br><em>Campement</em>');
  // Define the Group of markers
    var markers1 = L.layerGroup([marker1, marker2, marker3]);
    var markers2 = L.layerGroup([marker4, marker5]);


// Map initializer
  var map = L.map('map', {
    crs: L.CRS.Simple,
    minZoom: -4.6,
    wheelPxPerZoomLevel: 150,
    layers: [markers1, markers2],
    zoomControl: false,
    attributionControl: false
  });
  var bounds = [xy(0,16410), xy(16410,0)];
  var image = L.imageOverlay('/images/map_tree_of_life.jpg', bounds).addTo(map);
  map.fitBounds(bounds);


// Layers initializer
  var overlayMaps = {
    '<i class="icon-location blue"></i> Campements LCH': markers1,
    '<i class="icon-home green"></i> Campements alliés': markers2
  };

// Add Layers
  L.control.layers(null, overlayMaps, {collapsed: false}).addTo(map);


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
