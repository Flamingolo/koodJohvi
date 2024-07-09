

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

    const hast = window.location.hash.substring(1);
    if (hash === 'login'){
        root.appendChild(Login(setLoggedIn));
    } else if (hast === 'register'){
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