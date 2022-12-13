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
        SET_ALL_SUBJECTS(state, subjects) {
            state.subjects = subjects;
        }
    },
    actions: {
        async getAllSubjects(context, { group_id }) {
            fetch("http://localhost:4000/api/v1/subjects/" + group_id)
                .then(response => {
                    if (response.ok) {
                        return response.json();
                    } else {
                        throw Error(response.body);
                    }
                })
                .then(data => {
                    console.log(data);
                    context.commit("SET_ALL_SUBJECTS", data.data);
                })
                .catch(error => {
                    console.log(error);
                });
        }
    },
    getters: {
        allSubjects(state) {
            return state.subjects;
        }
    }
};