<!-- jQuery -->
<script src="https://code.jquery.com/jquery-2.2.3.min.js"></script>

<!-- Bootstrap -->
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js" integrity="sha384-0mSbJDEHialfmuBBQP6A4Qrprq5OVfW37PRR3j5ELqxss1yVqOtnepnHVP9aJ7xS" crossorigin="anonymous"></script>

<!-- jQWidgets -->
<script type="text/javascript" src="http://www.jqwidgets.com/public/jqwidgets/jqxcore.js"></script>
<script type="text/javascript" src="http://www.jqwidgets.com/public/jqwidgets/jqx-all.js"></script>

<!-- Bootstrap Table -->
<script src="//cdnjs.cloudflare.com/ajax/libs/bootstrap-table/1.10.1/bootstrap-table.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/bootstrap-table/1.10.1/locale/bootstrap-table-fr-FR.min.js"></script>

<!-- Leaflet & Plugins -->
<script src="/js/leaflet.js"></script>
<script src="/js/L.Control.MousePosition.js"></script>
{{ if .onMap }}
  <script src="/js/app.js"></script>
{{ end }}

<script type="text/javascript">
  // Modal control
    $("#about-btn").click(function() {
      $("#aboutModal").modal("show");
      $(".navbar-collapse.in").collapse("hide");
      return false;
    });
    $("#nav-btn").click(function() {
      $(".navbar-collapse").collapse("toggle");
      return false;
    });
</script>