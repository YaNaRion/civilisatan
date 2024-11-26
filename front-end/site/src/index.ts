async function greeting() {
  //D
  const div = document.getElementById("main");
  fetch("localhost:3333", {
    method: "GET",
  })
    .then((res) => res.json())
    .then((res) => {
      if (div) {
        div.innerHTML = res;
      }
    });
}
