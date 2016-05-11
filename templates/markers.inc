$(document).ready(function () {
  // Create the "icon select"
  var icon_list = [
    { html: '<i class="icon-home"></i>', value: "icon-home", group: "Générique" },
    { html: '<i class="icon-user"></i>', value: "icon-user", group: "Générique" },
    { html: '<i class="icon-flag"></i>', value: "icon-flag", group: "Générique" },
    //{ html: '<i class="icon-flag-3"></i>', value: "icon-flag-3" },
    { html: '<i class="icon-star"></i>', value: "icon-star", group: "Générique" },
    { html: '<i class="icon-star-half-alt"></i>', value: "icon-star-half-alt", group: "Générique" },
    { html: '<i class="icon-heart"></i>', value: "icon-heart", group: "Générique" },

    { html: '<i class="icon-tree"></i>', value: "icon-tree", group: "Nature" },
    { html: '<i class="icon-leaf"></i>', value: "icon-leaf", group: "Nature" },
    { html: '<i class="icon-garden"></i>', value: "icon-garden", group: "Nature" },
    { html: '<i class="icon-fire-1"></i>', value: "icon-fire-1", group: "Nature" },
    { html: '<i class="icon-asterisk"></i>', value: "icon-asterisk", group: "Nature" },
    { html: '<i class="icon-tint"></i>', value: "icon-tint", group: "Nature" },
    { html: '<i class="icon-water"></i>', value: "icon-water", group: "Nature" },
    { html: '<i class="icon-air"></i>', value: "icon-air", group: "Nature" },

    //{ html: '<i class="icon-tools"></i>', value: "icon-tools" },
    { html: '<i class="icon-wrench-3"></i>', value: "icon-wrench-3", group: "Outils" },
    { html: '<i class="icon-hammer"></i>', value: "icon-hammer", group: "Outils" },
    { html: '<i class="icon-bucket"></i>', value: "icon-bucket", group: "Outils" },
    { html: '<i class="icon-lock"></i>', value: "icon-lock", group: "Outils" },
    { html: '<i class="icon-key"></i>', value: "icon-key", group: "Outils" },

    { html: '<i class="icon-paw"></i>', value: "icon-paw", group: "Animaux" },
    { html: '<i class="icon-bug"></i>', value: "icon-bug", group: "Animaux" },
    { html: '<i class="icon-target-1"></i>', value: "icon-target-1", group: "Animaux" },

    { html: '<i class="icon-beaker"></i>', value: "icon-beaker", group: "Consommables" },
    { html: '<i class="icon-beer"></i>', value: "icon-beer", group: "Consommables" },

    { html: '<i class="icon-circle"></i>', value: "icon-circle", group: "Géométrie" },
    { html: '<i class="icon-dot-circled"></i>', value: "icon-dot-circled", group: "Géométrie" },
    { html: '<i class="icon-stop"></i>', value: "icon-stop", group: "Géométrie" },
    { html: '<i class="icon-up-dir-2"></i>', value: "icon-up-dir-2", group: "Géométrie" },
    { html: '<i class="icon-right-dir-2"></i>', value: "icon-right-dir-2", group: "Géométrie" },
    { html: '<i class="icon-down-dir-2"></i>', value: "icon-down-dir-2", group: "Géométrie" },
    { html: '<i class="icon-left-dir-2"></i>', value: "icon-left-dir-2", group: "Géométrie" },
    { html: '<i class="icon-cube"></i>', value: "icon-cube", group: "Géométrie" },
    { html: '<i class="icon-cubes"></i>', value: "icon-cubes", group: "Géométrie" },
  ];
  // Create a jqxDropDownList
  $("#icon").jqxDropDownList({
    theme: 'bootstrap',
    selectedIndex: 0,
    source: icon_list,
    width: '100%',
    height: '35px'
  });
  // Select value of an item
  $('#icon').on('select', function (event){
    $('#icon-field').val(event.args.item.value);
  });
  // Set the default state of icon field
  $('#icon-field').val(icon_list[$("#icon").jqxDropDownList('getSelectedIndex')].value);





  // Create the "color select"
  var color_list = [
    { html: '<div class="fleft"><div class="square bgblack"><div class="mlm lhm">Noir</div></div></div>', value: "black" },
    { html: '<div class="fleft"><div class="square bgwhite"><div class="mlm lhm">Blanc</div></div></div>', value: "white" },
    { html: '<div class="fleft"><div class="square bggrey"><div class="mlm lhm">Gris</div></div></div>', value: "grey" },
    { html: '<div class="fleft"><div class="square bgblue"><div class="mlm lhm">Bleu</div></div></div>', value: "blue" },
    { html: '<div class="fleft"><div class="square bgred"><div class="mlm lhm">Rouge</div></div></div>', value: "red" },
    { html: '<div class="fleft"><div class="square bggreen"><div class="mlm lhm">Vert</div></div></div>', value: "green" },
    { html: '<div class="fleft"><div class="square bgpurple"><div class="mlm lhm">Violet</div></div></div>', value: "purple" },
    { html: '<div class="fleft"><div class="square bgyellow"><div class="mlm lhm">Jaune</div></div></div>', value: "yellow" },
    { html: '<div class="fleft"><div class="square bgorange"><div class="mlm lhm">Orange</div></div></div>', value: "orange" },
    { html: '<div class="fleft"><div class="square bgbrown"><div class="mlm lhm">Marron</div></div></div>', value: "brown" },
  ];

  // Create a jqxDropDownList
  $("#color").jqxDropDownList({
    theme: 'bootstrap',
    selectedIndex: 0,
    source: color_list,
    width: '100%',
    height: '35px'
  });
  // Select value of an item
  $('#color').on('select', function (event){
    $('#color-field').val(event.args.item.value);
  });
  // Set the default state of color field
  $('#color-field').val(color_list[$("#color").jqxDropDownList('getSelectedIndex')].value);







  // Create the "type select"
  var group_list = [
    {{ range $key, $g := call .groups }}
      { html: '<i class="{{$g.Icon}} {{$g.Color}}"></i> {{$g.Name}}', value: "group{{$key}}", GroupName: "{{$g.Name}}" },
    {{ end }}
  ];
  // Create a jqxDropDownList
  $("#group").jqxDropDownList({
    theme: 'bootstrap',
    selectedIndex: 0,
    source: group_list,
    width: '100%',
    height: '35px'
  });
  // Select value of an item
  $('#group').on('select', function (event){
    $('#group-field').val(event.args.item.value);
    $('#inputTitle').val(group_list[$("#group").jqxDropDownList('getSelectedIndex')].GroupName);

    //$('#group-field').attr('Gname', group_list[GROUP_SELECTED_INDEX].Gname);
    //$('#inputTitle').val($('#group-field').attr('Gname'));
    
  });
  // Set the default state of group & inputTitle field
  $('#group-field').val(group_list[$("#group").jqxDropDownList('getSelectedIndex')].value);
  $('#inputTitle').val(group_list[$("#group").jqxDropDownList('getSelectedIndex')].GroupName);

});
