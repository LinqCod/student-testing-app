import { createWebHistory, createRouter } from "vue-router";
import MySubjects from "@/views/StudentSubjects.vue";
import StudentLogin from "@/views/StudentLogin.vue";
import store from "@/store";

const routes = [
    {
        path: "/",
        name: "StudentSubjects",
        component: MySubjects,
        beforeEnter: (to, from, next) => {
            if (!store.getters["auth/isLoggedIn"]) {
                next({ name: "Login" });
            } else {
                next();
            }
        }
    },
    {
        path: "/login",
        name: "Login",
        component: StudentLogin,
        beforeEnter: (to, from, next) => {
            if (store.getters["auth/isLoggedIn"]) {
                next({ name: "StudentSubjects" });
            } else {
                next();
            }
        }
    }
];

const router = createRouter({
    history: createWebHistory(),
    routes});

export default router;