"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
document.addEventListener("DOMContentLoaded", function() {
  var loginForm = document.getElementById("login-form");
  loginForm.addEventListener("submit", function(event) {
    event.preventDefault();
    alert("test");
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;

    // Handle the form submission here.
    // For example, send a request to your server to authenticate the user.
    console.log("Username: ".concat(username));
    console.log("Password: ".concat(password));
  });
});
