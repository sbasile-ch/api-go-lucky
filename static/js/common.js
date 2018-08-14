$(function() {
    // Create the selects
    var JsonArray = jQuery.parseJSON($('#command-json-list').text())

    $('#select-category').html(buildSelectHtml(1)).change(function() {
        $('#select-cmd').html(buildSelectHtml(2, this.value));
    });

    // Format the Json reeived to display
    var json = JSON.parse($('#output-area').text().replace(/(\r\n\t|\n|\r\t)/gm,''));
    $('#output-area').addClass('prettyprint');
    $('#output-area').text(JSON.stringify(json, null, 4));// Indented 4 spaces

function buildSelectHtml (numselect, Category){
    var i, len, v, html = ''

    for ( i = 0, len = JsonArray.length; i < len; i++ ) {
        v = JsonArray[i].CmdCategory;
        if ( numselect === 1 ) {
            html += '<option value="' + v + '">' + v + '</option>';
        } else if (v == Category ) {
            //var a = JsonArray[i].CmdValues;
            //for ( var c = 0, l = a.length; i < l; c++ ) {
            for ( var c = 0, l = JsonArray[i].CmdValues.length; c < l; c++ ) {
                //v = a[c];
                v = JsonArray[i].CmdValues[c];
                html += '<option value="' + v + '">' + v + '</option>';
            }
            return html;
        }
    }
    return html;
}

})
