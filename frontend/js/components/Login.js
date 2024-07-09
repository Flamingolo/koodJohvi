import { setLoggedIn } from "../main";

export default function Login(setLoggedIn){
    const loginSection = document.createElement("section");
    loginSection.innerHTML = `
        <h2>Login</h2>
        <form id="login-form">
            <div>
                <label for="email">Email</label>
                <input type="email" id="email" name="email" required>
            </div>
            <div>
                <label for="password">Password</label>
                <input type="password" id="password" name="password" required>
            </div>
            <button type="submit">Login</button>
        </form>
    `;

    loginSection.querySelector('#login-form').addEventListener('submit', async (e) => {
        e.preventDefault();

        const email = e.target.email.value;
        const password = e.target.password.value;

        const response = await fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ email, password })
        });

        if (response.ok) {
            setLoggedIn(true);
        } else {
            alert('Invalid email or password');
        }
    });

    return loginSection;   

}