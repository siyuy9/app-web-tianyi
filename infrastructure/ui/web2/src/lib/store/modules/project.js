import axios from "axios";
import Error from "../../../lib/main/Error";

const state = () => {
  return {
    id: null,
    parent_id: null,
    path: null,
    description: null,
    source: null,
    name: null,
    Branches: [],
  };
};

const actions = {};
const getters = {
  project: (state) => (path) => {
    if (path === state.path) return state;
    axios
      .get("/api/v1/projects", {
        params: {
          path: path,
        },
      })
      .then((response) => {
        Object.assign(state, response.data[0]);
      })
      .catch(Error);
    return state;
  },
};

const mutations = {};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
