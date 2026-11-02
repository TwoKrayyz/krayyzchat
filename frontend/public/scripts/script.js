const signUpForm = document.getElementById("signUpForm");

async function signUpFormHandler(event) {
    event.preventDefault();

    const signUpData = new FormData(signUpForm);
    const username = signUpData.get("username");
    const password = signUpData.get("password");

    const signUpPayload = {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            username: username,
            password: password
        })
    };

    fetch("/api/v1/signup", signUpPayload)
        .then(response => response.json())
        .then(data => console.log(data))
        .catch(error => console.error(error));
}

if (signUpForm) signUpForm.addEventListener("submit", signUpFormHandler);