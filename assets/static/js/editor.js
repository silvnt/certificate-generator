window.onload = function () {

  $('#bold_btn').on('click', function () {
    document.execCommand('bold');
    var text = document.getElementById('textarea').innerHTML;
    $('#textarea-show').html(text);
  });
  $('#italic_btn').on('click', function () {
    document.execCommand('italic');
    var text = document.getElementById('textarea').innerHTML;
    $('#textarea-show').html(text);
  });
  $('#left_btn').on('click', function () {
    document.execCommand('justifyLeft');
    var text = document.getElementById('textarea').innerHTML;
    $('#textarea-show').html(text);
  });
  $('#center_btn').on('click', function () {
    document.execCommand('justifyCenter');
    var text = document.getElementById('textarea').innerHTML;
    $('#textarea-show').html(text);
  });
  $('#right_btn').on('click', function () {
    document.execCommand('justifyRight');
    var text = document.getElementById('textarea').innerHTML;
    $('#textarea-show').html(text);
  });
  $('#justify_btn').on('click', function () {
    document.execCommand('justifyFull');
    var text = document.getElementById('textarea').innerHTML;
    $('#textarea-show').html(text);
  });
  /*$('#text-size-input').on(function () {
    document.execCommand('fontsize', true, $(this).val());
    var text = document.getElementById('textarea').innerHTML;
    $('#textarea-show').html(text);
  });*/
}

function readURL(input) {
  if (input.files && input.files[0]) {
      var reader = new FileReader();

      reader.onload = function (e) {
          $('#certif-bg').attr('src', e.target.result)
          $('#certif-bg').attr('style', 'visibility: visible')
          
      };

      reader.readAsDataURL(input.files[0]);
  }
}

function removeimage() {
      $('#certif-bg').attr('src', '#')
      $('#certif-bg').attr('style', 'visibility: hidden')
}