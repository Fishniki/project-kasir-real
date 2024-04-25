document.addEventListener('DOMContentLoaded', function () {
    const navbarToggle = document.querySelector('[data-collapse-toggle="navbar-multi-level"]');
    const navbarMenu = document.getElementById('navbar-multi-level');

    navbarToggle.addEventListener('click', function () {
        navbarMenu.classList.toggle('hidden');
    });

    const dropdownNavbarLink = document.getElementById('dropdownNavbarLink');
    const dropdownNavbar = document.getElementById('dropdownNavbar');

    dropdownNavbarLink.addEventListener('click', function () {
        dropdownNavbar.classList.toggle('hidden');
    });

    // Sticky Navbar
    window.onscroll = function () {
        myFunction()
    };

    var navbar = document.getElementById("navbar-multi-level");

    var sticky = navbar.offsetTop;

    function myFunction() {
        if (window.pageYOffset >= sticky) {
            navbar.classList.add("sticky")
        } else {
            navbar.classList.remove("sticky");
        }
    }


        document.querySelectorAll('.decrement-button').forEach(function(button) {
        button.addEventListener('click', function() {
            var input = button.nextElementSibling;
            var value = parseInt(input.value);
            if (!isNaN(value) && value > 1) {
                input.value = value - 1;
            }
        });
    });

    document.querySelectorAll('.increment-button').forEach(function(button) {
    button.addEventListener('click', function() {
            var input = button.previousElementSibling;
            var value = parseInt(input.value);
            if (!isNaN(value)) {
                input.value = value + 1;
            } else {
                input.value = 1;
            }
        });
    });
});