export default {
    namespaced: true,
    state: {
        student: {
            loggedIn: false,
            full_name: "",
            email: "",
            personal_number: "",
            access_token: "",
            group: {
                id: -1,
                title: ""
            }
        }
    },
    mutations: {
        LOGIN(state, {student}) {
            state.student.loggedIn = true;
            state.student.full_name = student.full_name;
            state.student.email = student.email;
            state.student.personal_number = student.personal_number;
            state.student.access_token = student.access_token;
            state.student.group.id = student.group.id;
            state.student.group.title = student.group.title;
        },
        LOGOUT(state) {
            state.student.loggedIn = false;
            state.student.full_name = "";
            state.student.email = "";
            state.student.personal_number = "";
            state.student.access_token = "";
            state.student.group.id = -1;
            state.student.group.title = "";
        }
    },
    actions: {
        async login(context, { email, password }) {
            return fetch("http://localhost:4000/api/v1/students/login", {
                method: "POST",
                body: JSON.stringify({
                    email: email,
                    password: password })
            })
                .then(response => {
                    if (!response.ok) {
                        throw new Error("Cannot login!");
                    }
                    return response.json();
                }).then(data => {
                    let student = {
                        "access_token":    data.data.access_token,
                        "full_name": data.data.full_name,
                        "email": data.data.email,
                        "personal_number": data.data.personal_number,
                        "group": data.data.group,
                    }
                    context.commit("LOGIN",
                        { student: student });
                }).catch(error => {
                    context.commit("LOGOUT");
                    throw error;
                });
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
        },
        getTokenHeader(state) {
            return "Bearer " + state.student.access_token;
        }
    }
};