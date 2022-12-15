export default {
    namespaced: true,
    state: {
        sections: [
        ]
    },
    mutations: {
        SET_ALL_SECTIONS(state, sections) {
            state.sections = sections;
        }
    },
    actions: {
        async getAllSections(context, { subject_id }) {
            fetch("http://localhost:4000/api/v1/tasks/categories/" + subject_id, {
                headers: {
                    Authorization: context.rootGetters["auth/getTokenHeader"]
                }
            })
                .then(response => {
                    if (response.ok) {
                        return response.json();
                    } else {
                        throw Error(response.body);
                    }
                })
                .then(data => {
                    console.log(data.data)
                    context.commit("SET_ALL_SECTIONS", data.data);
                })
                .catch(error => {
                    console.log(error);
                });
        }
    },
    getters: {
        allSections(state) {
            return state.sections;
        }
    }
};