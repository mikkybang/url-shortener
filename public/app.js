const form = document.getElementById('shorten-form')

if (form){
    form.addEventListener('submit', async e => {
        e.preventDefault()

        const url = form.url.value
        const submitBtn = document.getElementById("submit-btn");

        submitBtn.innerHTML = `<span class="spinner-border spinner-border-sm" role="status" aria-hidden="false"></span>`;

        const body = {
            url
        }

       const result = await fetch("/api/create",{
            method: 'POST',
            headers: {
              'Content-Type': 'application/json;charset=utf-8'
            },
            body: JSON.stringify(body)
          }
        )

        const response = await result.json()

        console.log(response)

        submitBtn.innerHTML = "Shorten";

    })
}