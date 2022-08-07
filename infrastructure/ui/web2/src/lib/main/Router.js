import { createRouter, createWebHistory } from "vue-router";
import { AdminSidebarItems, BrowseSidebarItemsWrapper } from "./SidebarItems";
import App from "../../components/app/AppMain.vue";
import EventBus from "./EventBus";

function _import(importPath) {
  return () => import("../../components/" + importPath + ".vue");
}

const routes = [
  {
    path: "main",
    meta: {
      containerClassAddFlex: false,
      routerViewShowOutside: false,
      showTopBar: true,
      showFooter: true,
    },
    children: [
      {
        path: "browse",
        meta: {
          sidebarItems: BrowseSidebarItemsWrapper,
        },
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
        name: "admin",
        meta: {
          sidebarItems: AdminSidebarItems,
        },
        children: [
          {
            path: "swagger",
            name: "admin_swagger",
            component: _import("admin/AdminSwagger"),
          },
        ],
      },
    ],
  },
  {
    path: "detached",
    component: _import("common/FormPage"),
    meta: {
      containerClassAddFlex: true,
      routerViewShowOutside: true,
      showTopBar: false,
      showFooter: false,
    },
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

const routesWrapper = [
  {
    path: "",
    name: "root",
    redirect: { name: "browse_projects" },
  },
  {
    path: "/pages",
    component: App,
    children: routes,
    // if a root doesn't have a name, redirect to not_found
    // info - https://router.vuejs.org/guide/advanced/navigation-guards.html#global-before-guards
    beforeEnter: (to) => {
      return to.name ? true : { name: "invalid_not_found" };
    },
  },
  // catch all
  // https://stackoverflow.com/a/40194152
  {
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

export default Router;
