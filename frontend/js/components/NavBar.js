import { isLoggedIn, setLoggedIn } from "../main.js";

export function NavBar(isLoggedIn, setLoggedIn){
    const nav = document.createElement("nav");
    nav.innerHTML = `
        <ul>
            <li><a href="#home">Home</a></li>
            <li><a href="#categories">Categories</a></li>
           ${isLoggedIn ? `
            <li><a href="#profile">Profile</a></li>
            <li><a href="#logout" id="logout">Logout</a></li>
            ` : `
            <li><a href="#login">Login</a></li>}
            <li><a href="#register">Register</a></li>
            `}
        </ul>
    `;

        if (isLoggedIn) {
            nav.querySelector("#logout").addEventListener("click", () => {
                e.preventDefault();
                setLoggedIn(false);
        });
    }

return nav;
}