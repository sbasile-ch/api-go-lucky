$(function() {
    var json = JSON.parse($('#output-area').text().replace(/(\r\n\t|\n|\r\t)/gm,''));
    $('#output-area').addClass('prettyprint');
    $('#output-area').text(JSON.stringify(json, null, 4));// Indented 4 spaces

    $('.submit-form').click(function() {
          .submit();
    });
})
