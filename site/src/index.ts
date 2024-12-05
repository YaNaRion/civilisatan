document.addEventListener("DOMContentLoaded", () => {
  const form = document.getElementById("credential-form") as HTMLFormElement;
  form.addEventListener("submit", (event: Event) => {
    event.preventDefault();
    const formData = new FormData(form);
    const credentials = {
      username: formData.get("username") as string,
      password: formData.get("password") as string,
    };
    console.log(credentials);
    globalThis.location.href = "./src/admin/admin.html";
  });
});
