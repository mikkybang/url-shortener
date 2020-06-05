const form = document.getElementById('shorten-form')

if (form){
    form.addEventListener('submit', e => {
        e.preventDefault()
        
        const url = form.url.value

        console.log(url)

    })
}