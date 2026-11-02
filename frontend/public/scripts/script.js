const signUpForm = document.getElementById("signUpForm");
const signInForm = document.getElementById("signInForm");

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

async function signInFormHandler(event) {
    event.preventDefault();

    const signInData = new FormData(signInForm);
    const username = signInData.get("username");
    const password = signInData.get("password");

    const signInPayload = {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            username: username,
            password: password
        })
    };

    fetch("/api/v1/signin", signInPayload)
        .then(response => response.json())
        .then(data => console.log(data))
        .catch(error => console.error(error));
}

if (signInForm) signInForm.addEventListener("submit", signInFormHandler);
if (signUpForm) signUpForm.addEventListener("submit", signUpFormHandler);