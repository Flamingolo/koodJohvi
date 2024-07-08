
let isLoggedOn = false;

function setLoggedIn(){
    isLoggedOn = status;
    init();
}

function init() {
    const root = document.getElementById("root");

    // Clear existing content
    root.innerHTML = "";

    // Append Header
    root.appendChild(Header());
    root.appendChild(NavBar(isLoggedOn, setLoggedIn));
    root.appendChild(Posts());
    root.appendChild(Chat());
    root.appendChild(Footer());
}

init();


window.addEventListener("hashchange", function () {
    init();
});