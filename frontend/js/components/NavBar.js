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


    async function createNavItems(){
        for (const link of navLinks){
            if ((link.requiresAuth && !isLoggedIn()) || (link.requiresAuth === false && isLoggedIn())){
                continue
            }

            if (link.dropdown) {
                const dropDown = document.createElement("div");
                dropDown.classList.add("dropdown");

                const dropDownButton = document.createElement("button");
                dropDownButton.classList.add("dropbtn");
                dropDownButton.textContent = link.name;
                dropDown.appendChild(dropDownButton);

                const dropDownContent = document.createElement("div");
                dropDownContent.classList.add("dropdown-content");

                const categories = await getCategories();
                categories.forEach(category => {
                    const categoryLink = document.createElement("a");
                    categoryLink.href = `/category/${category}`;
                    categoryLink.textContent = item;
                    dropDownContent.appendChild(categoryLink);
                });
                dropDown.appendChild(dropDownContent);
                navbar.appendChild(dropDown);
            } else {
                const navItem = document.createElement("a");
                navItem.href = link.href;
                navItem.textContent = link.name;
                navItem.classList.add("nav-link");
                navbar.appendChild(navItem);
            }
        }
    }
    createNavItems();
});