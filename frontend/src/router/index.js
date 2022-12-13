import { createWebHistory, createRouter } from "vue-router";
import MySubjects from "@/views/MySubjects.vue";

const routes = [
    {
        path: "/",
        name: "MySubjects",
        component: MySubjects
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes});

export default router;