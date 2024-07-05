document.addEventListener("DOMContentLoaded", () => {
    const navbar = document.getElementById("navbar")
    
    // Getting categories from database
    async function getCategories(){
        try {
            const response = await fetch("/categories");
            if (!response.ok){
                throw new Error("Network response was not ok")
            }
            const categories = await response.json();
            return categories.map(category => category.name);
        } catch (error) {
            console.error("Failed to catch categories:", error)
            return [];
        }
    }


    // Links for navbar
    const navLinks = [
        { name: "Home", href: "/"},
        { name: "Login", href: "/login", requiresAuth: false},
        { name: "Register", href: "/register", requiresAuth: false},
        { name: "Categories", href: "#", dropdown: true},
        { name: "Profile", href: "/profile"},
        { name: "Logout", href: "/logout", requiresAuth: true},
        { name: "", href: "/"},
        { name: "", href: "/"},
    ]

    navLinks.forEach(link => {
        const navItem = document.createElement("a");
        navItem.href = link.href;
        navItem.textContent = link.name;
        navItem.classList.add("nav-link");
        navbar.appendChild(navItem)
    })
})