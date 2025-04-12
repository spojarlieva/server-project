document.getElementById("login-button").addEventListener("click", async (event) => {
    event.preventDefault()

    const email = document.getElementById("email-login").value
    const password = document.getElementById("password-login").value
    const errorText = document.getElementById("login-error-message")

    try {
        const response = await fetch("https://server-project-production-d36b.up.railway.app/users/login", {
            method: "POST", headers: {
                "Content-Type": "application/json"
            }, body: JSON.stringify({email, password})
        })

        switch (response.status) {
            case 201:
                const body = await response.json()
                const token = body.token;
                sessionStorage.setItem("token", token)
                alert("Успешно влизане");
                window.location.href = "events.html"
                break;
            case 401:
                errorText.textContent = "Грешен емайл или парола"
                errorText.style.display = "block"
                console.error(await response.json())
                break
            default:
                errorText.textContent = "Възникна грешка"
                errorText.style.display = "block"
                console.error(await response.json())
        }
    } catch (error) {
        console.log(error)
    }
})