import Vuex from "vuex";
import subjectsModule from "./subjectStore/index.js";
import authModule from "./authStore/index.js";

export default new Vuex.Store({
    state: {},
    mutations: {},
    actions: {},
    modules: { subjects: subjectsModule, auth: authModule }
});