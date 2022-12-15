import { createWebHistory, createRouter } from "vue-router";
import StudentSubjects from "@/views/StudentSubjects.vue";
import TaskSections from "@/views/TaskSections.vue";
import StudentLogin from "@/views/StudentLogin.vue";
import store from "@/store";

const routes = [
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
    },
    {
        path: "/",
        name: "StudentSubjects",
        component: StudentSubjects,
        beforeEnter: (to, from, next) => {
            if (!store.getters["auth/isLoggedIn"]) {
                next({ name: "Login" });
            } else {
                next();
            }
        }
    },
    {
        path: "/sections/:subject_id",
        props: true,
        name: "TaskSections",
        component: TaskSections,
        beforeEnter: (to, from, next) => {
            if (!store.getters["auth/isLoggedIn"]) {
                next({ name: "Login" });
            } else {
                next();
            }
        }
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

export default router;