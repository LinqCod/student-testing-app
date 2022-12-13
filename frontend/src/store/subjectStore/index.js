export default {
    namespaced: true,
    state: {
        subjects: [
            {
                id: 1,
                title: "Математика"
            },
            {
                id: 2,
                title: "Программирование"
            }
        ]
    },
    mutations: {
        ADD_SUBJECT(state, subject) {
            state.subjects.push(subject);
        },
        DELETE_SUBJECT(state, id) {
            state.subjects = state.subjects.filter(subject => subject.id !== id);
        }
    },
    actions: {
        async addSubject(context, subject) {
            context.commit("ADD_SUBJECT", subject);
        },
        async deleteSubject(context, id) {
            context.commit("DELETE_SUBJECT", id);
        }
    },
    getters: {
        allSubjects(state) {
            return state.subjects;
        }
    }
};