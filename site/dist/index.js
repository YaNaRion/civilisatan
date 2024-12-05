"use strict";
// async function greeting() {
//   //D
//   const div = document.getElementById("main");
//   if (div) {
//     div.innerHTML = "Hot save x2";
//     console.log("lol x3");
//   }
// }
// function checkCredential(form: Event) {
//   alert(form);
//   // if (form.username.value === "yann" && form.password.value === "123")
//   // globalThis.location.href = "./admin/admin.html";
// }
document.addEventListener("DOMContentLoaded", function () {
    var form = document.getElementById("credential-form");
    form.addEventListener("submit", function (event) {
        event.preventDefault();
        var formData = new FormData(form);
        var credentials = {
            username: formData.get("username"),
            password: formData.get("password"),
        };
        console.log(credentials);
        globalThis.location.href = "./src/admin/admin.html";
        // Implement your logic here to submit the form data securely
    });
});
