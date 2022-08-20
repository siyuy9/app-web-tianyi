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
  clean(context) {
    context.commit("user", {
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
    });
  },
};

const mutations = {
  user(state, user) {
    // update entire state
    Object.assign(state, user);
  },
  id(state, id) {
    state.id = id;
  },
  admin(state, admin) {
    state.admin = admin;
  },
  created_at(state, created_at) {
    state.created_at = created_at;
  },
  updated_at(state, updated_at) {
    state.updated_at = updated_at;
  },
  username(state, username) {
    state.username = username;
  },
  email(state, email) {
    state.email = email;
  },
  bio(state, bio) {
    state.bio = bio;
  },
  image(state, image) {
    state.image = image;
  },
  roles(state, roles) {
    state.roles = roles;
  },
  token(state, token) {
    state.token = token;
    axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
  },
};

const getters = {
  isLoggedIn(state) {
    return !!state.token;
  },
  admin(state) {
    return state.admin;
  },
  token(state) {
    return state.token;
  },
  user(state) {
    return state;
  },
  id(state) {
    return state.id;
  },
  created_at(state) {
    return state.created_at;
  },
  updated_at(state) {
    return state.updated_at;
  },
  username(state) {
    return state.username;
  },
  email(state) {
    return state.email;
  },
  bio(state) {
    return state.bio;
  },
  image(state) {
    return state.image;
  },
  roles(state) {
    return state.roles;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
