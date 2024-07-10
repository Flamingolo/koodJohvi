import {} from "./frontend/js/"

document.addEventListener("DOMContentLoaded", () => {
    const rootElement = document.getElementById("root");

    function navigateTo(url) {
        window.history.pushState(null, null, url);
        router();
    }

    function router(){
        const routes = {
            "/": renderHomePage,
            "/login": renderLoginPage,
            "/register": renderRegisterPage,
            "/create-post": renderCreatePostPage,
            "/create-comment": renderCreateCommentPage,
        };

        const path = location.pathname;
        const route = routes[path]

        if (route) {
            route();
        } else {
            renderNotFoundPage();
        }
    }

    window.addEventListener('popstate', router);
    document.body.addEventListener('click', e => {
        if (e.target.matches("[data-link]")) {
            e.preventDefault();
            navigateTo(e.target.href);
        }
    });

    router()
});