export default {
    namespaced: true,
    state: {
        student: {
            id: 1,
            full_name: "Балин Максим Максимович",
            email: "danny12123@yandex.ru",
            personal_number: "23ИЕ23",
            group: {
                id: 1,
                title: "ИКБО-24-20"
            },
            loggedIn: false
        }
    },
    mutations: {
        LOGIN(state) {
            state.student.loggedIn = true;
        },
        LOGOUT(state) {
            state.student.loggedIn = false;
        }
    },
    actions: {
        async login(context) {
            context.commit("LOGIN");
        },
        async logout(context) {
            context.commit("LOGOUT");
        }
    },
    getters: {
        currentStudent(state) {
            return state.student;
        },
        isLoggedIn(state) {
            if (!state.student) return false;
            return state.student.loggedIn;
        }
    }
};