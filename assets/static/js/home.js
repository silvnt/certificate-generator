window.onload = function () {
  $(document).ready(function () {
    $('#editor').jqxEditor({
      tools: 'bold italic underline | font size color background | left center right | ul ol | clean'
    });
  });

  $('#input-bg').on('change', function () {
    if (this.files && this.files[0]) {
      var reader = new FileReader();

      reader.onload = function (e) {
        $('#editor-bg').attr('src', e.target.result)
        $('#editor-bg').css('visibility', 'visible')

      };

      reader.readAsDataURL(this.files[0]);
    }
  });

  $('#remove-bg').on('click', function () {
    $('#editor-bg').attr('src', '#')
    $('#editor-bg').css('visibility', 'hidden')
  });

  $('#editor').on('change', function () {
    $('#input-editor').html($('#editor').html())
  })
}
