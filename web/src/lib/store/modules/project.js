import axios from "../../main/Axios";

const emptyProject = {
  id: null,
  parent_id: null,
  default_branch: null,
  path: null,
  description: null,
  source: null,
  name: null,
  branches: {},
};

const emptyBranch = {
  id: "",
  created_at: "",
  updated_at: "",
  deleted_at: {
    Time: "",
    Valid: false,
  },
  project_id: "",
  name: "",
  config: {
    jobs: [],
    pipelines: [],
  },
};

const state = () => {
  return {
    projects: {},
    currentID: null,
  };
};

const mutations = {
  updateProject: (state, { projectID, project_data }) => {
    // Object.assign doesn't work on undefined
    if (!state.projects[projectID]) {
      state.projects[projectID] = { ...emptyProject };
    }
    Object.assign(state.projects[projectID], project_data);
  },
  updateBranch(state, { projectID, branchName, branchData }) {
    // Object.assign doesn't work on undefined;
    if (!state.projects[projectID].branches) {
      state.projects[projectID].branches = {};
    }
    if (!state.projects[projectID].branches[branchName]) {
      state.projects[projectID].branches[branchName] = { ...emptyBranch };
    }
    Object.assign(state.projects[projectID].branches[branchName], branchData);
  },
  updateBranches(state, { projectID, branches }) {
    Object.assign(
      state.projects[projectID].branches,
      branches.map((branch) => [branch.name, branch])
    );
  },
  currentID(state, projectID) {
    state.currentID = projectID;
  },
};

const actions = {
  loadProject({ commit }, path) {
    return new Promise((resolve, reject) =>
      axios
        .get("/api/v1/projects", {
          params: {
            path: path,
          },
        })
        .then((response) => {
          var project = response.data.data[0];
          commit("updateProject", {
            projectID: project.id,
            project_data: project,
          });
          commit("currentID", project.id);
          resolve(response);
        })
        .catch(reject)
    );
  },
  loadBranch({ commit, getters }, branchName) {
    return new Promise((resolve, reject) =>
      axios
        .get(`/api/v1/projects/${getters.currentID}/branches/${branchName}`)
        .then((response) => {
          commit("updateBranch", {
            projectID: getters.currentID,
            branchName: branchName,
            branchData: response.data.data,
          });
          resolve(response);
        })
        .catch(reject)
    );
  },
  updateRemoteBranch({ commit, getters }, branchName) {
    return new Promise((resolve, reject) =>
      axios
        .put(`/api/v1/projects/${getters.currentID}/branches/${branchName}`)
        .then((response) => {
          commit("updateBranch", {
            projectID: getters.currentID,
            branchName: branchName,
            branchData: response.data.data,
          });
          resolve(response);
        })
        .catch(reject)
    );
  },
  loadBranches({ commit, getters }) {
    return new Promise((resolve, reject) =>
      axios
        .get(`/api/v1/projects/${getters.currentID}/branches`)
        .then((response) => {
          commit("updateBranches", {
            projectID: getters.currentID,
            branches: response.data.data,
          });
          resolve(response);
        })
        .catch(reject)
    );
  },
};
const getters = {
  id: (state) => (projectID) => state.projects[projectID].id,
  parentID: (state) => (projectID) => state.projects[projectID].parent_id,
  defaultBranch: (state) => (projectID) =>
    state.projects[projectID].default_branch,
  path: (state) => (projectID) => state.projects[projectID].path,
  description: (state) => (projectID) => state.projects[projectID].description,
  source: (state) => (projectID) => state.projects[projectID].source,
  name: (state) => (projectID) => state.projects[projectID].name,
  branch: (state) => (projectID, branchName) =>
    state.projects[projectID].branches[branchName],

  currentID: (state) => state.currentID,
  currentParentID: (state) => state.projects[state.currentID].parent_id,
  currentDefaultBranchName: (state) =>
    state.projects[state.currentID].default_branch,
  currentDefaultBranch: (state) =>
    state.projects[state.currentID].branches[
      state.projects[state.currentID].default_branch
    ],
  currentPath: (state) => state.projects[state.currentID].path,
  currentDescription: (state) => state.projects[state.currentID].description,
  currentSource: (state) => state.projects[state.currentID].source,
  currentName: (state) => state.projects[state.currentID].name,
  currentBranches: (state) => state.projects[state.currentID].branches,
  currentBranch: (state) => (branchName) =>
    state.projects[state.currentID].branches[branchName],
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
