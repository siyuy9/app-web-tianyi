const state = () => {
  return {
    layoutMode: "static",
    theme: { name: "lara-light-indigo", dark: false },
    ripple: true,
    inputStyle: "outlined",
  };
};

const actions = {};
const getters = {
  layoutMode(state) {
    return state.layoutMode;
  },
  theme(state) {
    return state.theme;
  },
  isDark(state) {
    return state.theme.dark;
  },
  logo(state) {
    return state.theme.dark
      ? "/images/logo-white.svg"
      : "/images/logo-dark.svg";
  },
};

const mutations = {
  layoutMode(state, layoutMode) {
    state.layoutMode = layoutMode;
  },
  theme(state, value) {
    state.theme = { name: value.name, dark: value.dark };
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
