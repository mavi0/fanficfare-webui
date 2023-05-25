
(function ($) {
    "use strict";

     /*==================================================================
    [ Focus input ]*/
    $('.input100').each(function(){
        $(this).on('blur', function(){
            if($(this).val().trim() != "") {
                $(this).addClass('has-val');
            }
            else {
                $(this).removeClass('has-val');
            }
        })    
    })
  
  
    /*==================================================================
    [ Validate ]*/
    var input = $('.validate-input .input100');

    $('.validate-form').on('submit',function(){
        var check = true;

        for(var i=0; i<input.length; i++) {
            if(validate(input[i]) == false){
                showValidate(input[i]);
                check=false;
            }
        }

        return check;
    });


    $('.validate-form .input100').each(function(){
        $(this).focus(function(){
           hideValidate(this);
        });
    });

    function validate (input) {
        if($(input).attr('type') == 'email' || $(input).attr('name') == 'email') {
            if($(input).val().trim().match(/^([a-zA-Z0-9_\-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{1,5}|[0-9]{1,3})(\]?)$/) == null) {
                return false;
            }
        }
        else {
            if($(input).val().trim() == ''){
                return false;
            }
        }
    }

    function showValidate(input) {
        var thisAlert = $(input).parent();

        $(thisAlert).addClass('alert-validate');
    }

    function hideValidate(input) {
        var thisAlert = $(input).parent();

        $(thisAlert).removeClass('alert-validate');
    }
    
    function getURLParam() {
        const url_params = window.location.search.substring(1).split("&");
        for (let i=0;i<url_params.length;i++) {
            const key_value = url_params[i].split("=");
            if (key_value[0] == "url") {
                return key_value[1];
            }
        }    
    }

    const urlInParam = getURLParam();
    if (urlInParam) {
        $('.input100').each(function(){
            $(this).focus();
            $(this).val(urlInParam);
            $(this).blur();
        });
    }

})(jQuery);


function alert_click()
{
    alert("This may take some time - sometimes a few minutes... Please do not leave this page, your book will download automatically!");
    document.getElementById("progress-bar").style.display = "block"; 
    // $(".progress-bar").show();


}
