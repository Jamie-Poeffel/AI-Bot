import { createRouter, createWebHistory } from "vue-router";

const routes = [
    {
        name: "Login",
        path: "/login",
        component: () => import("../views/LoginView.vue"),
    },
]  

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
});

export default router;