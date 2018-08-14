$(function() {
    var json = $('#output-area').text();
    json = JSON.stringify(json, null, 4); // Indented 4 spaces
    $('#output-area').text(json);
})
