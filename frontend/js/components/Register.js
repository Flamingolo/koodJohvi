export function Register(){
    const registerSection = document.createElement("section");
    registerSection.innerHTML = `
        <h2>Register</h2>
        <form id="register-form">
            <div>
                <label for="email">Email:</label>
                <input type="email" id="email" name="email" required>
            </div>
            <div>
                <label for="username">Username:</label>
                <input type="text" id="username" name="username" required>
            </div>
            <div>
                <label for="password">Password:</label>
                <input type="password" id="password" name="password" required>
            </div>
            <div>
                <label for="nickname">Nickname:</label>
                <input type="text" id="nickname" name="nickname">
            </div>
            <div>
                <label for="age">Age:</label>
                <input type="number" id="age" name="age">
            </div>
            <div>
                <label for="gender">Gender:</label>
                <select id="gender" name="gender">
                    <option value="male>Male</option>
                    <option value="female">Female</option>
                    <option value="other">Other</option>
                </select>
            </div>
            <div>
                <label for="first-name">First Name:</label>
                <input type="text" id="first-name" name="first-name">
            </div>
            <div>
                <label for="last-name">Last Name:</label>
                <input type="text" id="last-name" name="last-name">
            </div>
            <button type="submit">Register</button>
        </form>
    `;


    registerSection.querySelector('#register-form').addEventListener('submit', async (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);

        const response = await fetch('http://localhost:8080/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(Object.fromEntries(formData.entries()))
        }).then(response => {
            if (response.ok){
                alert('Registered successfully');
            } else {
                alert('Registration failed');
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Registration failed');
        })
    });

    return registerSection;
}