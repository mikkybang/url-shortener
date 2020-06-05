const form = document.getElementById('shorten-form')

if (form){
    form.addEventListener('submit', async e => {
        e.preventDefault()

        const url = form.url.value

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

    })
}