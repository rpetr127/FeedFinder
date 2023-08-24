(function() {
    var ref_list = document.querySelector('p').querySelectorAll('a'),
    my_input = document.querySelector('input');

    for(i=0; i<=ref_list.length; i++) {
        ref_list[i].addEventListener('click', function(event) {
            event.preventDefault();
            my_input.value = event.target.text;
        })
    }
})()