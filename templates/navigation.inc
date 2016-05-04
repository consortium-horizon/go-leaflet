<div role="navigation" class="navbar navbar-inverse navbar-fixed-top">
  <div class="container-fluid">
    <div class="navbar-header">
      <div class="navbar-icon-container">
        <a id="nav-btn" class="navbar-icon pull-right visible-xs" href="#"><i class="icon-bars icon-lg white"></i></a>
        <a id="sidebar-toggle-btn" class="navbar-icon pull-right visible-xs" href="#"><i class="icon-search icon-lg white"></i></a>
      </div>
      <a href="/" class="navbar-brand">GameMap</a>
    </div>
    <div class="navbar-collapse collapse">
      <ul class="nav navbar-nav">
        <li><a id="about-btn" data-target=".navbar-collapse.in" data-toggle="collapse" href="#"><i class="icon-help-circled white"></i>&nbsp;&nbsp;A propos</a></li>
        <li><a id="about-btn" data-target=".navbar-collapse.in" data-toggle="collapse" href="/markers/new"><i class="icon-plus white"></i>&nbsp;&nbsp;Ajouter un marqueur</a></li>
        <li><a id="about-btn" data-target=".navbar-collapse.in" data-toggle="collapse" href="/markers/manage"><i class="icon-pencil white"></i>&nbsp;&nbsp;Modifier les marqueurs</a></li>
        <li><a id="full-extent-btn" data-target=".navbar-collapse.in" data-toggle="collapse" href="#"><i class="icon-resize-small-2 white"></i>&nbsp;&nbsp;Réinitialiser le zoom</a></li>
      </ul>
    </div><!--/.navbar-collapse -->
  </div>
</div><!--/.navigation -->

<div class="modal fade" id="aboutModal" tabindex="-1" role="dialog">
  <div class="modal-dialog modal-lg">
    <div class="modal-content">
      <div class="modal-header">
        <button class="close" type="button" data-dismiss="modal" aria-hidden="true">&times;</button>
        <h4 class="modal-title">Welcome to the BootLeaf template!</h4>
      </div>
      <div class="modal-body">
        <div class="tab-content" id="aboutTabsContent">
          <div class="tab-pane fade active in" id="about">
            <p>A simple, responsive template for building web mapping applications with <a href="http://getbootstrap.com/">Bootstrap 3</a>, <a href="http://leafletjs.com/" target="_blank">Leaflet</a>, and <a href="http://twitter.github.io/typeahead.js/" target="_blank">typeahead.js</a>. Open source, MIT licensed, and available on <a href="https://github.com/bmcbride/bootleaf" target="_blank">GitHub</a>.</p>
            <div class="panel panel-primary">
              <div class="panel-heading">Features</div>
              <ul class="list-group">
                <li class="list-group-item">Fullscreen mobile-friendly map template with responsive navbar and modal placeholders</li>
                <li class="list-group-item">jQuery loading of external GeoJSON files</li>
                <li class="list-group-item">Logical multiple layer marker clustering via the <a href="https://github.com/Leaflet/Leaflet.markercluster" target="_blank">leaflet marker cluster plugin</a></li>
                <li class="list-group-item">Elegant client-side multi-layer feature search with autocomplete using <a href="http://twitter.github.io/typeahead.js/" target="_blank">typeahead.js</a></li>
                <li class="list-group-item">Responsive sidebar feature list synced with map bounds, which includes sorting and filtering via <a href="http://listjs.com/" target="_blank">list.js</a></li>
                <li class="list-group-item">Marker icons included in grouped layer control via the <a href="https://github.com/ismyrnow/Leaflet.groupedlayercontrol" target="_blank">grouped layer control plugin</a></li>
              </ul>
            </div>
          </div>
        </div>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
      </div>
    </div><!-- /.modal-content -->
  </div><!-- /.modal-dialog -->
</div><!-- /.modal -->
