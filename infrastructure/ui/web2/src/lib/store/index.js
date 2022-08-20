import { createStore } from "vuex";

import theme from "./modules/theme";
import user from "./modules/user";
import project from "./modules/project";

// https://vuex.vuejs.org/guide/
// https://github.com/vuejs/vuex/tree/main/examples/classic/shopping-cart/store
export default createStore({
  modules: {
    theme,
    user,
    project,
  },
});
