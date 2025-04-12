document.addEventListener('DOMContentLoaded', async () => {
    const token = sessionStorage.getItem("token");

    if (token) {
        await getRegisteredEvents(token)
    } else {
        await getEvents()
    }
});

const getEvents = async () => {
    try {
        const response = await fetch("https://server-project-production-b671.up.railway.app/events")
        if (!response.ok) {
            console.error("Error getting events")
            alert("Грешка при зараждането на събития")
        }

        const data = await response.json()
        const eventsList = document.getElementById("events-list")

        data.forEach((event) => {
            const listItem = document.createElement("li")

            const eventTitle = document.createElement("h1")
            eventTitle.innerText = event.title;
            listItem.appendChild(eventTitle)

            const image = document.createElement("img")
            image.src = event.image_url
            image.alt = event.title
            listItem.appendChild(image)

            const eventDate = document.createElement("p")
            eventDate.innerText = `Дата: ${new Date(event.date).toLocaleDateString()}`
            listItem.appendChild(eventDate)


            const eventAddress = document.createElement("p")
            eventAddress.innerText = event.address
            listItem.appendChild(eventAddress)

            const eventDescription = document.createElement('p');
            eventDescription.innerText = event.description;
            listItem.appendChild(eventDescription)

            eventsList.appendChild(listItem)
        })

    } catch (error) {
        alert("Грешка при зараждането на събития")
        console.error(error)
    }
}

const getRegisteredEvents = async (token) => {
    try {
        const response = await fetch("https://server-project-production-b671.up.railway.app/events/registered", {
            headers: {
                "Authorization": `Bearer ${token}`
            }
        })

        if (!response.ok) {
            console.error("Error getting events")
            alert("Грешка при зараждането на събития")
        }

        const data = await response.json()
        const eventsList = document.getElementById("events-list")

        data.forEach((event) => {
                const listItem = document.createElement("li")

                const eventTitle = document.createElement("h1")
                eventTitle.innerText = event.title;
                listItem.appendChild(eventTitle)

                const image = document.createElement("img")
                image.src = event.image_url
                image.alt = event.title
                listItem.appendChild(image)

                const eventDate = document.createElement("p")
                eventDate.innerText = `Дата: ${new Date(event.date).toLocaleDateString()}`
                listItem.appendChild(eventDate)

                const eventAddress = document.createElement("p")
                eventAddress.innerText = event.address
                listItem.appendChild(eventAddress)

                const eventDescription = document.createElement('p');
                eventDescription.innerText = event.description;
                listItem.appendChild(eventDescription)


                const checkBoxContainer = document.createElement("div");
                checkBoxContainer.classList.add("checkbox-container");

                const checkBox = document.createElement("input");
                checkBox.type = "checkbox";
                checkBox.checked = event.is_registered;
                checkBox.id = `event-${event.id}`;

                const label = document.createElement("label");
                label.htmlFor = checkBox.id;
                label.textContent = checkBox.checked ? "Регистриран" : "Не регистриран";

                checkBox.addEventListener("change", async () => {
                    label.textContent = checkBox.checked ? "Регистриран" : "Не регистриран";

                    try {
                        const response = await fetch(
                            `https://server-project-production-b671.up.railway.app/events/register/${event.id}`,
                            {
                                method: checkBox.checked ? "POST" : "DELETE",
                                headers: {
                                    "Authorization": `Bearer ${token}`
                                }
                            }
                        );

                        if (!response.ok) {
                            console.error("Failed to register or unregister for event");
                            console.error(await response.json());
                            alert("Възникна грешка");
                        }
                    } catch (error) {
                        console.error("Failed to register or unregister for event");
                        alert("Възникна грешка");
                    }
                });

                checkBoxContainer.appendChild(checkBox);
                checkBoxContainer.appendChild(label);
                listItem.appendChild(checkBoxContainer);
                eventsList.appendChild(listItem)
            }
        )
    } catch
        (error) {
        alert("Грешка при зараждането на събития")
        console.error(error)
    }
}