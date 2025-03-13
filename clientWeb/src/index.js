document.addEventListener("DOMContentLoaded", () => {
  const loginForm = document.getElementById("login-form")

  loginForm.addEventListener("submit", (event) => {
    event.preventDefault();
    alert("test");

    const username = (document.getElementById("username")).value;
    const password = (document.getElementById("password")).value;

    console.log(`Username: ${username}`);
    console.log(`Password: ${password}`);
  });
});
