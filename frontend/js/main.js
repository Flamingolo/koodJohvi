import Header from "./components/Header.js";
import NavBar from "./components/NavBar.js";
import Posts from "./components/Posts.js";
import Chat from "./components/Chat.js";
import Login from "./components/Login.js";
import Register from "./components/Register.js";
import Footer from "./components/Footer.js";

function isLoggedIn(){
    return localStorage.getItem("isLoggedOn") === "true"
}

function setLoggedIn(status){
    localStorage.setItem("isLoggedOn", status ? 'true' : 'false');
    init();
}


function init() {
    const root = document.getElementById("root");

    // Clear existing content
    root.innerHTML = "";

    // Append Header
    root.appendChild(Header());
    root.appendChild(NavBar(isLoggedIn(), setLoggedIn));

    const hash = window.location.hash.substring(1);
    if (hash === 'login'){
        root.appendChild(Login(setLoggedIn));
    } else if (hash === 'register'){
        root.appendChild(Register());
    } else {
        root.appendChild(Posts());
        root.appendChild(Chat());
    }

    // Append Footer
    root.appendChild(Footer());
}

// Initial loop
init();

// Handle navigation and component loading
window.addEventListener("hashchange", init);