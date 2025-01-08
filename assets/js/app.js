
const form = document.getElementById('loginForm');
const emailInput = document.getElementById('email');
const passwordInput = document.getElementById('password');
const emailError = document.getElementById('emailError');
const passError = document.getElementById('passError');


form.addEventListener('submit', function (e) {
  let isValid = true;

  
  if (emailInput.value.trim() === '') {
    emailError.style.display = 'block';
    isValid = false;
  } else {
    emailError.style.display = 'none';
  }

 
  if (passwordInput.value.trim() === '') {
    passError.style.display = 'block';
    passMessage.style.display = 'none';
    isValid = false;
  }  else {
    passError.style.display = 'none';
  }

  if (!isValid) {
    e.preventDefault();
  }
});
