document.addEventListener("DOMContentLoaded", async () => {
    try {
        let response = await fetch("https://server-project-production-d36b.up.railway.app/gallery/images")
        if (!response.ok) {
            console.error("Error getting images")
            alert("Грешка при зареждане на снимките")
        }

        const data = await response.json()
        let div = document.getElementById("gallery-div")

        data.forEach((image) => {
            let el = document.createElement("img")
            el.src = image.url
            el.alt = "gallery image"
            el.loading = "lazy"
            div.appendChild(el)
        })

    } catch (error) {
        alert("Грешка при зареждане на снимките")
        console.error(error)
    }
})