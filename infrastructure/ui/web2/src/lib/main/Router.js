import { createRouter, createWebHistory } from "vue-router";
import App from "../../components/app/AppMain.vue";
import EventBus from "./EventBus";
import VuexStore from "../store";
import SidebarItems from "../../lib/main/SidebarItems";

function _import(importPath) {
  return () => import(`../../components/${importPath}.vue`);
}

const routes = [
  {
    path: "main",
    component: App,
    children: [
      {
        path: "browse",
        children: [
          {
            path: "groups",
            name: "browse_groups",
            component: _import("browse/BrowseGroups"),
          },
          {
            path: "projects",
            name: "browse_projects",
            component: _import("browse/BrowseProjects"),
          },
        ],
      },
      {
        path: "admin",
        meta: {
          sidebarItems: SidebarItems.admin.items,
          sidebarAsRoot: SidebarItems.admin.asRoot,
        },
        beforeEnter: () => {
          return VuexStore.getters["user/admin"];
        },
        children: [],
      },
      {
        path: "user",
        meta: {
          sidebarItems: SidebarItems.user.items,
          sidebarAsRoot: SidebarItems.user.asRoot,
        },
        children: [
          {
            path: "profile",
            name: "user_profile",
            component: _import("user/UserProfile"),
          },
          {
            path: "appearance",
            name: "user_appearance",
            component: _import("user/UserAppearance"),
          },
        ],
      },
      {
        path: "help",
        meta: {
          sidebarItems: SidebarItems.help.items,
          sidebarAsRoot: SidebarItems.help.asRoot,
        },
        children: [
          {
            path: "swagger",
            name: "help_swagger",
            component: _import("help/HelpSwagger"),
          },
        ],
      },
      {
        path: "projects",
        children: [
          {
            path: "create",
            name: "project_create",
            component: _import("project/ProjectCreate"),
          },
        ],
      },
    ],
  },
  {
    path: "detached",
    component: _import("common/FormPage"),
    children: [
      {
        path: "auth",
        children: [
          {
            path: "login",
            name: "auth_login",
            component: _import("auth/LoginPage"),
          },
          {
            path: "registration",
            name: "auth_registration",
            component: _import("auth/RegistrationPage"),
          },
          {
            path: "forgot_password",
            name: "auth_forgot_password",
            component: _import("auth/ForgotPassword"),
          },
        ],
      },
      {
        path: "invalid",
        meta: {
          showGoBackLink: true,
          shadow: 0,
        },
        children: [
          {
            path: "error",
            name: "invalid_error",
            component: _import("invalid/ErrorPage"),
            meta: { logoColor: "error" },
          },
          {
            path: "not_found",
            name: "invalid_not_found",
            component: _import("invalid/NotFound"),
            meta: { logoColor: "blue" },
          },
          {
            path: "access_denied",
            name: "invalid_access_denied",
            component: _import("invalid/AccessDenied"),
            meta: { logoColor: "orange" },
          },
        ],
      },
    ],
  },
];

const projectRoutes = [
  {
    path: "pipelines",
    name: "project_pipelines",
    component: _import("project/ProjectPipelines"),
  },
  {
    path: "pipelines/:pipeline_branch/:pipeline_name",
    name: "project_pipeline",
    component: _import("project/ProjectPipeline"),
  },
];

const routesWrapper = [
  {
    path: "",
    name: "root",
    redirect: { name: "browse_projects" },
  },
  {
    path: "/-",
    children: routes,
  },
  {
    component: App,
    children: [
      {
        path: "/:project_path(.*)/-",
        meta: {
          sidebarItems: SidebarItems.project.items,
          sidebarAsRoot: SidebarItems.project.asRoot,
        },
        name: "project",
        component: _import("project/ProjectBase"),
        redirect: { name: "project_pipelines" },
        children: projectRoutes,
      },
    ],
  },
  {
    // catch all
    // https://router.vuejs.org/guide/essentials/dynamic-matching.html#catch-all-404-not-found-route
    path: "/:pathMatch(.*)*",
    redirect: { name: "invalid_not_found" },
  },
];

const Router = createRouter({
  history: createWebHistory(),
  routes: routesWrapper,
});

Router.onError((error) => {
  EventBus.emit("app-toast-add", {
    severity: "error",
    summary: "A routing error has occured",
    detail: error,
    life: 3000,
  });
});

Router.beforeEach((to) => {
  if (
    // make sure the user is authenticated
    !VuexStore.getters["user/isLoggedIn"] &&
    // allowed pages
    !["auth_login", "auth_registration", "auth_forgot_password"].includes(
      to.name
    )
  ) {
    // redirect the user to the login page
    return { name: "auth_login" };
  }
  // if the route doesn't have a name, redirect to not_found
  // https://router.vuejs.org/guide/advanced/navigation-guards.html#global-before-guards
  return to.name ? true : { name: "invalid_not_found" };
});

export default Router;
