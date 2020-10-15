const uploading = `<div class="ui icon message">
                        <i class="notched circle loading icon"></i>
                        <div class="content">
                            <div class="header">Just one second</div>
                            <p>We're uploading that content for you.</p>
                        </div>
                    </div>`;

const formCard = $('.ui.cards > .card > .content > .description');
const submitButton = $('#submit-button');
const filesUploadForm = $('#files-upload-form');
const filesInput = filesUploadForm.find('input[name="files[]"]');

filesInput.change(function(event) {
    if ($(this).val().length > 0) {
        submitButton.prop('disabled', false);
    }
});

submitButton.bind('click', function(event) {
    event.preventDefault();
    if (filesInput.val().length > 0) {
        formCard.append(uploading);
        filesUploadForm.submit();
    }
});