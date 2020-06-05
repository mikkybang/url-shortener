const form = document.getElementById("shorten-form");

if (form) {
  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const url = form.url.value;
    const submitBtn = document.getElementById("submit-btn");
    const message = document.getElementById("message-text");

    submitBtn.innerHTML = `<span class="spinner-border spinner-border-sm" role="status" aria-hidden="false"></span>`;

    const body = {
      url,
    };

    const result = await fetch("/api/create", {
      method: "POST",
      headers: {
        "Content-Type": "application/json;charset=utf-8",
      },
      body: JSON.stringify(body),
    });

    const response = await result.json();

    const shortUrl = `${window.location.hostname}/${response.hash}`;
    submitBtn.innerHTML = "Shorten";
    message.innerHTML = `Your url has been shortened to: ${shortUrl}`;
  });
}
