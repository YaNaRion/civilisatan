document.addEventListener("DOMContentLoaded", () => {
  const loginForm = document.getElementById("login-form") as HTMLFormElement;

  loginForm.addEventListener("submit", (event) => {
    event.preventDefault();
    alert("test");

    const username = (document.getElementById("username") as HTMLInputElement)
      .value;
    const password = (document.getElementById("password") as HTMLInputElement)
      .value;

    // Handle the form submission here.
    // For example, send a request to your server to authenticate the user.

    console.log(`Username: ${username}`);
    console.log(`Password: ${password}`);
  });
});
