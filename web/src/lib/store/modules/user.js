import axios from "../../main/Axios";

const state = () => {
  return {
    id: null,
    token: null,
    admin: false,
    created_at: "",
    updated_at: "",
    username: "",
    email: "",
    bio: "",
    image: null,
    roles: null,
  };
};

const actions = {
  clean({ commit }) {
    commit("user", {});
  },
  login({ commit, dispatch }, auth) {
    return axios
      .post("/api/v1/users/login", {
        username: auth.username,
        password: auth.password,
      })
      .then((response) => {
        commit("user", response.data.data);
        if (auth.remember_me) {
          dispatch("persist");
        }
      });
  },
  initializeStore({ commit }) {
    var user = localStorage.getItem("user");
    if (user) {
      commit("user", JSON.parse(user));
    }
  },
  persist({ getters }) {
    localStorage.setItem("user", JSON.stringify(getters.user));
  },
};

const mutations = {
  // update entire state
  user: (state, user) => Object.assign(state, user),
  id: (state, id) => (state.id = id),
  admin: (state, admin) => (state.admin = admin),
  created_at: (state, created_at) => (state.created_at = created_at),
  updated_at: (state, updated_at) => (state.updated_at = updated_at),
  username: (state, username) => (state.username = username),
  email: (state, email) => (state.email = email),
  bio: (state, bio) => (state.bio = bio),
  image: (state, image) => (state.image = image),
  roles: (state, roles) => (state.roles = roles),
  token(state, token) {
    state.token = token;
  },
};

const getters = {
  isLoggedIn: (state) => !!state.id,
  isLoaded: (state) => !!state.id,
  admin: (state) => state.admin,
  token: (state) => state.token,
  user: (state) => state,
  id: (state) => state.id,
  created_at: (state) => state.created_at,
  updated_at: (state) => state.updated_at,
  username: (state) => state.username,
  email: (state) => state.email,
  bio: (state) => state.bio,
  image: (state) => state.image,
  roles: (state) => state.roles,
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
