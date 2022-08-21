import axios from "axios";

const state = () => {
  return {
    id: null,
    parent_id: null,
    default_branch: null,
    path: null,
    description: null,
    source: null,
    name: null,
    branch: {},
  };
};

const mutations = {
  updateProject(state, data) {
    Object.assign(state, data);
  },
  updateBranch(state, data) {
    Object.assign(state.branch, data);
  },
};

const actions = {
  async loadProject({ commit }, path) {
    const projectResponse = await axios.get("/api/v1/projects", {
      params: {
        path: path,
      },
    });
    commit("updateProject", projectResponse.data[0]);
  },
  async loadBranch({ commit, state }, branch) {
    const branchResponse = await axios.get(
      `/api/v1/projects/${state.id}/branches/${branch}`
    );
    commit("updateBranch", branchResponse.data);
  },
};
const getters = {
  id: (state) => state.id,
  parent_id: (state) => state.parent_id,
  default_branch: (state) => state.default_branch,
  path: (state) => state.path,
  description: (state) => state.description,
  source: (state) => state.source,
  name: (state) => state.name,
  branch: (state) => state.branch,
  branchName: (state) => state.branch.name,
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
