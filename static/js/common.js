
//__________________________________________________
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
    var selCat = $('#select-category');
    var selCmd = $('#select-cmd');
    selCat.val(selCat.attr('value')).change();
    selCmd.val(selCmd.attr('value')).change();
}

//__________________________________________________
function addOptions (array, valSelected){
    var i, len, v, html = ''
    array.sort();
    for ( i = 0, len = array.length; i < len; i++ ) {
        v = array[i];
        var sel = ( v == valSelected ) ? 'selected' : '';
        html += '<option ' + sel + ' value="' + v + '">' + v + '</option>';
    }
    return html
}



//__________________________________________________
function buildSelectHtml (numselect, Category, Cmd){
    var i, len, v, html = ''
    var options = [];

    for ( i = 0, len = JsonArray.length; i < len; i++ ) {
        v = JsonArray[i].CmdCategory;
        if ( numselect === 1 ) {
            options [i] = v
        } else if (v == Category ) {
            for ( var c = 0, l = JsonArray[i].CmdValues.length; c < l; c++ ) {
                v = JsonArray[i].CmdValues[c].CmdName;
                if ( numselect === 2 ) {
                    options [c] = v
                } else if (v == Cmd ) {
                    return  JsonArray[i].CmdValues[c].Args;
                }
            }
            return addOptions (options, Cmd);
        }
    }
    return addOptions (options, Category);
}


//__________________________________________________
function populateForm(array) {
    var html='';
    if (array) {
        for ( var i = 0, l = array.length; i < l; i++ ) {
            var v = array[i];
            if (v != "CompanyNum" ) {  // CompanyNum is common so strip here
                var val = $('#'+v).val();
                html += '<div class="block">';
                html += '<label for="' + v + '" class="form-label">' + v + ':</label>';
                html += '<input type="text" id="' + v + '" name="' + v + '" value="' + val +'">';
                html += '</div>';
            }
        }
    }
    $('#hidden-args').remove();
    $('#variable-args').html(html);
}


//__________________________________________________
function loadCmdList() {
    window.JsonArray = [];
    var txt = $('#json-cmds-list').text();
    if (txt.replace(/\s/g, '').length) {
        window.JsonArray = jQuery.parseJSON(txt);
    }
}

//__________________________________________________
function formatJson() {
    //var json = JSON.parse($('#output-area').text().replace(/(\r\n\t|\n|\r\t)/gm,''));
    var txt = $('#output-area').text().replace(/(\r\n\t|\n|\r\t)/gm,'');
    if (txt.replace(/\s/g, '').length) {
        var json = JSON.parse(txt);
        $('#output-area').text(JSON.stringify(json, null, 4));// Indented 4 spaces
        $('#output-area').addClass('prettyprint');
    }
}

//__________________________________________________
function initApp() {
    loadCmdList();
    populateSelects();
    formatJson();
}


