import axios from "axios";

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
  clean: ({ commit }) => commit("user", {}),
  async login({ commit }, auth) {
    const login = await axios.post("/api/v1/users/login", {
      username: auth.username,
      password: auth.password,
    });
    commit("token", login.data.token);
    const user = await axios.get(`/api/v1/users/user/${login.data.id}`);
    commit("user", user.data);
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
    axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
  },
};

const getters = {
  isLoggedIn: (state) => !!state.token,
  isLoaded: (state) => !!state.token && state.username,
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
