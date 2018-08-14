function populateSelects () {
    $('#select-category').html(buildSelectHtml(1)).change(function() {
        var Category = this.value;
        $('#variable-args').html('');
        var html = buildSelectHtml(2, Category);
        if (html) {
            $('#select-cmd').html(html).change(function() {
                    populateForm(buildSelectHtml(3, Category, this.value));
            }).show();
        } else {
            $('#select-cmd').hide();
        }
    });
}

function buildSelectHtml (numselect, Category, Cmd){
    var i, len, v, html = ''

    for ( i = 0, len = JsonArray.length; i < len; i++ ) {
        v = JsonArray[i].CmdCategory;
        if ( numselect === 1 ) {
            html += '<option value="' + v + '">' + v + '</option>';
        } else if (v == Category ) {
            for ( var c = 0, l = JsonArray[i].CmdValues.length; c < l; c++ ) {
                v = JsonArray[i].CmdValues[c].CmdName;
                if ( numselect === 2 ) {
                    html += '<option value="' + v + '">' + v + '</option>';
                } else if (v == Cmd ) {
                    return  JsonArray[i].CmdValues[c].Args;
                }
            }
            return html;
        }
    }
    return html;
}


function populateForm(array) {
    var html='';
    if (array) {
        for ( var i = 0, l = array.length; i < l; i++ ) {
            var v = array[i];
            html += '<input type="text" id="' + v + '" name="' + v + '"><label for="' + v + '">' + v + '</label>';
        }
    }
    $('#variable-args').html(html);
}


function loadCmdList() {
    window.JsonArray = [];
    var txt = $('#json-cmds-list').text();
    if (txt.replace(/\s/g, '').length) {
        window.JsonArray = jQuery.parseJSON(txt);
    }
}

function formatJson() {
    //var json = JSON.parse($('#output-area').text().replace(/(\r\n\t|\n|\r\t)/gm,''));
    var txt = $('#output-area').text().replace(/(\r\n\t|\n|\r\t)/gm,'');
    if (txt.replace(/\s/g, '').length) {
        var json = JSON.parse(txt);
        $('#output-area').text(JSON.stringify(json, null, 4));// Indented 4 spaces
        $('#output-area').addClass('prettyprint');
    }
}

function initApp() {
    loadCmdList();
    populateSelects();
    formatJson();
}


