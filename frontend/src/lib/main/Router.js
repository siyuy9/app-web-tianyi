import { createRouter, createWebHistory } from "vue-router";
import { AdminSidebarItems, BrowseSidebarItemsWrapper } from "./SidebarItems";
import App from "../../components/app/AppMain.vue";

function _import(importPath) {
  return () => import("../../components/" + importPath + ".vue");
}

const routes = [
  {
    path: "/",
    name: "root",
    redirect: { name: "browse" },
  },
  {
    path: "/pages/browse",
    name: "browse",
    component: App,
    props: {
      sidebarItems: BrowseSidebarItemsWrapper,
    },
    children: [
      {
        path: "/pages/browse/groups",
        name: "browse_groups",
        component: _import("browse/BrowseGroups"),
      },
      {
        path: "/pages/browse/projects",
        name: "browse_projects",
        component: _import("browse/BrowseProjects"),
      },
    ],
  },
  {
    path: "/pages/admin",
    name: "admin",
    component: App,
    props: {
      sidebarItems: AdminSidebarItems,
    },
    children: [
      {
        path: "/pages/admin/swagger",
        name: "admin_swagger",
        component: _import("admin/AdminSwagger"),
      },
    ],
  },
  {
    path: "/pages/auth/login",
    name: "login",
    component: _import("auth/LoginPage"),
  },
  {
    path: "/pages/auth/registration",
    name: "registration",
    component: _import("auth/RegistrationPage"),
  },
  {
    path: "/pages/invalid/error",
    name: "error",
    component: _import("invalid/ErrorPage"),
  },
  {
    // catch all
    // https://stackoverflow.com/a/40194152
    path: "/:pathMatch(.*)*",
    name: "not_found",
    component: _import("invalid/NotFound"),
  },
];

const Router = createRouter({
  history: createWebHistory(),
  routes,
});

export default Router;
