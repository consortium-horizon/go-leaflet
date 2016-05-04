$(document).ready(function () {
  var icon_list = [
    { html: '<i class="icon-star"></i>', value: "icon-star" },
    { html: '<i class="icon-tree"></i>', value: "icon-tree" },
    { html: '<i class="icon-fire-1"></i>', value: "icon-fire-1" },
    { html: '<i class="icon-snow"></i>', value: "icon-snow" },
    { html: '<i class="icon-home"></i>', value: "icon-home" },
    { html: '<i class="icon-heart"></i>', value: "icon-heart" },
    { html: '<i class="icon-star-half-alt"></i>', value: "icon-star-half-alt" },
    { html: '<i class="icon-user"></i>', value: "icon-user" },
    { html: '<i class="icon-user-secret"></i>', value: "icon-user-secret" },
    { html: '<i class="icon-ok-circled"></i>', value: "icon-ok-circled" },
    { html: '<i class="icon-plus-circled"></i>', value: "icon-plus-circled" },
    { html: '<i class="icon-cancel-circled"></i>', value: "icon-cancel-circled" },
    { html: '<i class="icon-minus-circled"></i>', value: "icon-minus-circled" },
    { html: '<i class="icon-lock"></i>', value: "icon-lock" },
    { html: '<i class="icon-bookmark"></i>', value: "icon-bookmark" },
    { html: '<i class="icon-flag"></i>', value: "icon-flag" },
    { html: '<i class="icon-comment"></i>', value: "icon-comment" },
    { html: '<i class="icon-chat"></i>', value: "icon-chat" },
    { html: '<i class="icon-bell-alt"></i>', value: "icon-bell-alt" },
    { html: '<i class="icon-attention-circled"></i>', value: "icon-attention-circled" },
    { html: '<i class="icon-location"></i>', value: "icon-location" },
    { html: '<i class="icon-direction"></i>', value: "icon-direction" },
    { html: '<i class="icon-cog"></i>', value: "icon-cog" },
    { html: '<i class="icon-tint"></i>', value: "icon-tint" },
    { html: '<i class="icon-fire"></i>', value: "icon-fire" },
    { html: '<i class="icon-asterisk"></i>', value: "icon-asterisk" },
    { html: '<i class="icon-home"></i>', value: "icon-home" },
    { html: '<i class="icon-dot-circled"></i>', value: "icon-dot-circled" },
    { html: '<i class="icon-key"></i>', value: "icon-key" },
    { html: '<i class="icon-bug"></i>', value: "icon-bug" },
    { html: '<i class="icon-beaker"></i>', value: "icon-beaker" },
    { html: '<i class="icon-gift"></i>', value: "icon-gift" },
    { html: '<i class="icon-beer"></i>', value: "icon-beer" },
  ];
  // Define Default Index
  var ICON_SELECTED_INDEX = 0;
  // Create a jqxDropDownList
  $("#icon").jqxDropDownList({
    theme: 'bootstrap',
    selectedIndex: ICON_SELECTED_INDEX,
    source: icon_list,
    width: '100%',
    height: '35px'
  });
  // Select value of an item
  $('#icon').on('select', function (event){
    $('#icon-field').val(event.args.item.value);
  });
  $('#icon-field').val(icon_list[ICON_SELECTED_INDEX].value);


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

  // Define Default Index
  var COLOR_SELECTED_INDEX = 0;
  // Create a jqxDropDownList
  $("#color").jqxDropDownList({
    theme: 'bootstrap',
    selectedIndex: COLOR_SELECTED_INDEX,
    source: color_list,
    width: '100%',
    height: '35px'
  });
  // Select value of an item
  $('#color').on('select', function (event){
    $('#color-field').val(event.args.item.value);
  });
  $('#color-field').val(color_list[COLOR_SELECTED_INDEX].value);
});