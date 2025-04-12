const validateEmail = (email) => {
    if (!email) {
        return "Имейлът е празен";
    }

    const splitEmail = email.split("@");
    if (splitEmail.length !== 2) {
        return "Имейлът е невалиден";
    }

    const [localPart, domainPart] = splitEmail;
    if (!localPart) {
        return "Липсва локалната част на имейла";
    }

    if (!domainPart) {
        return "Липсва домейн частта на имейла";
    }

    const splitDomain = domainPart.split(".");
    if (splitDomain.length !== 2) {
        return "Домейнът е невалиден";
    }

    if (splitDomain[0].length < 1 || splitDomain[1].length < 2) {
        return "Домейнът е невалиден";
    }

    return "";
}

const validatePassword = (password) => {
    if (!password) {
        return "Паролата е празна";
    }

    if (!/[0-9]/.test(password)) {
        return "Паролата трябва да съдържа поне една цифра";
    }

    if (!/[a-z]/.test(password)) {
        return "Паролата трябва да съдържа поне една малка буква";
    }

    if (!/[A-Z]/.test(password)) {
        return "Паролата трябва да съдържа поне една главна буква";
    }

    if (!/[!@#$^&*()_\-+]/.test(password)) {
        return "Паролата трябва да съдържа поне един специален символ";
    }

    return "";
}

document.getElementById("register-button").addEventListener("click", async (event) => {
    event.preventDefault()

    const email = document.getElementById("email-register").value.trim()
    const password = document.getElementById("password-register").value.trim()
    const confirmPassword = document.getElementById("password-register-confirm").value.trim()
    const errorText = document.getElementById("register-error-message")

    let errorMessage = validateEmail(email)
    if (errorMessage) {
        errorText.textContent = errorMessage
        errorText.style.display = "block";
        return
    }

    if (password !== confirmPassword) {
        errorText.textContent = "Паролите не съвпадат"
        errorText.style.display = "block";
        return
    }

    errorMessage = validatePassword(password)
    if (errorMessage) {
        errorText.textContent = errorMessage
        errorText.style.display = "block";
        return
    }

    try {
        const response = await fetch("https://server-project-production-b671.up.railway.app/users/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({email, password})
        })

        switch (response.status) {
            case 201:
                alert("Успешна регистрация!");
                window.location.href = "login.html";
                break;
            case 409:
                errorText.textContent = "Емайла вече се използва"
                errorText.style.display = "block";
                break
            default:
                errorText.textContent = "Възнинка грешка"
                errorText.style.display = "block";
                console.error(await response.json())
        }
    } catch (error) {
        console.error(error)
    }

})