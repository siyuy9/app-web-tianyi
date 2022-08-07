import "primevue/resources/primevue.min.css";
import "primeflex/primeflex.css";
import "primeicons/primeicons.css";
import "prismjs/themes/prism-coy.css";

import "./assets/styles/layout.scss";
import "./assets/demo/flags/flags.css";

import { createApp, reactive } from "vue";
import PrimeVue from "primevue/config";
import BadgeDirective from "primevue/badgedirective";
import ConfirmationService from "primevue/confirmationservice";
import Ripple from "primevue/ripple";
import StyleClass from "primevue/styleclass";
import Tooltip from "primevue/tooltip";

import ToastService from "./lib/main/ToastService";
import Router from "./lib/main/Router";
import AppWrapper from "./components/app/AppWrapper.vue";
import CodeHighlight from "./lib/app/AppCodeHighlight";
import UniversalComponents from "./lib/main/UniversalComponents";

Router.beforeEach(function (to, from, next) {
  window.scrollTo(0, 0);
  next();
});

const app = createApp(AppWrapper);

app.config.globalProperties.$appState = reactive({
  theme: "lara-light-indigo",
  darkTheme: false,
});

app.use(PrimeVue, { ripple: true, inputStyle: "outlined" });
app.use(ConfirmationService);
app.use(ToastService);
app.use(Router);

const directives = {
  tooltip: Tooltip,
  ripple: Ripple,
  code: CodeHighlight,
  badge: BadgeDirective,
  styleclass: StyleClass,
};

for (const [key, value] of Object.entries(UniversalComponents)) {
  app.component(key, value);
}

for (const [key, value] of Object.entries(directives)) {
  app.directive(key, value);
}

app.mount("#app");
