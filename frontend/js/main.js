import Header from "./components/Header.js";
import NavBar from "./components/NavBar.js";
import Posts from "./components/Posts.js";
// import Chat from "./components/Chat.js";
import Login from "./components/Login.js";
import Register from "./components/Register.js";
// import Footer from "./components/Footer.js";

export function isLoggedIn() {
    return localStorage.getItem("isLoggedOn") === "true";
}

export function setLoggedIn(status) {
    localStorage.setItem("isLoggedOn", status ? 'true' : 'false');
    init();
}

function init() {
    const root = document.getElementById("root");

    // Clear existing content
    root.innerHTML = "";

    // Append Header
    root.appendChild(Header());
    console.log("Header added");
    root.appendChild(NavBar(isLoggedIn, setLoggedIn));
    console.log("NavBar added");

    const hash = window.location.hash.substring(1);
    if (hash === 'login') {
        root.appendChild(Login(setLoggedIn));
        console.log("Login added");
    } else if (hash === 'register') {
        root.appendChild(Register());
        console.log("Register added");
    } else {
        root.appendChild(Posts());
        console.log("Posts added");
        // root.appendChild(Chat());
        // console.log("Chat added");
    }

    // Append Footer
    // root.appendChild(Footer());
    // console.log("Footer added");
}

// Initial loop
init();

// Handle navigation and component loading
window.addEventListener("hashchange", init);
