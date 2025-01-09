const phone = document.getElementById("mobileno");
const registrationForm = document.getElementById('registrationForm');
const email = document.getElementById('email');
const Fname = document.getElementById('uname');
const fullName = document.getElementById('fullName');
const invalidName = document.getElementById('invalidName');
const usernameError = document.getElementById('usernameError');
const invalidUsername = document.getElementById('invalidUsername');
const emailError = document.getElementById('emailError');
const username = document.getElementById('username');
const password = document.getElementById('upassword');
const passnameError = document.getElementById('passnameError');
const passUsername = document.getElementById('passUsername');

function PhoneValidation() {
    const pattern = /^\d{10}$/;
    const isValid = pattern.test(phone.value);
    document.getElementById("phoneValidationMsg").textContent = isValid ? '' : 'Please enter a valid 10-digit phone number.';
    return isValid;
}

registrationForm.addEventListener('submit', function (e) {
    let isValid = true;

    if (Fname.value.trim() === '') {
        fullName.style.display = 'block';
        isValid = false;
    } else if (!/^[a-zA-Z\s]+$/.test(Fname.value)) {
        fullName.style.display = 'none';
        invalidName.style.display = 'block';
        isValid = false;
    } else {
        fullName.style.display = 'none';
        invalidName.style.display = 'none';
    }

    if (email.value.trim() === '') {
        emailError.style.display = 'block';
        isValid = false;
    } else {
        emailError.style.display = 'none';
    }

    if (username.value.trim() === '') {
        usernameError.style.display = 'block';
        invalidUsername.style.display = 'none';
        isValid = false;
    } else if (!/^[a-zA-Z0-9]+$/.test(username.value)) {
        usernameError.style.display = 'none';
        invalidUsername.style.display = 'block';
        isValid = false;
    } else {
        usernameError.style.display = 'none';
        invalidUsername.style.display = 'none';
    }
    if (password.value.trim() === '') {
        passnameError.style.display = 'block';
        passUsername.style.display = 'none';
        isValid = false;
    } else if (!PasswordValidation()) {
        passnameError.style.display = 'none';
        passUsername.style.display = 'block';
        isValid = false;
    } else {
        passnameError.style.display = 'none';
        passUsername.style.display = 'none';
    }

   
    if (!PhoneValidation()) {
        isValid = false;
    }

    
    if (!isValid) {
        e.preventDefault();
    }
});
