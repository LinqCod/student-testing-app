import Vuex from "vuex";
import subjectsModule from "./subjectStore/index.js";
import sectionsModule from "./sectionStore/index.js";
import authModule from "./authStore/index.js";

export default new Vuex.Store({
    state: {},
    mutations: {},
    actions: {},
    modules: { auth: authModule, subjects: subjectsModule, sections: sectionsModule }
});