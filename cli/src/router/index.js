import { createRouter, createWebHistory } from "vue-router";

const routes = [
    {
        name: "Login",
        path: "/",
        component: () => import("../views/LoginView.vue"),
    },
    {
        name: "Home",
        path: "/home",
        component: () => import("../views/HomeView.vue"),
    },
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
});

export default router;