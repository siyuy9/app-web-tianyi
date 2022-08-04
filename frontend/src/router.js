import { createRouter, createWebHistory } from "vue-router";
import AppProjects from "./AppProjects.vue";
import AppAdmin from "./AppAdmin.vue";

const routes = [
  {
    path: "/",
    name: "root",
    redirect: { name: "projects" },
  },
  {
    path: "/pages/projects",
    name: "projects",
    component: AppProjects,
    children: [
      {
        path: "",
        name: "dashboard",
        component: () => import("./components/Dashboard.vue"),
      },
      {
        path: "/formlayout",
        name: "formlayout",
        component: () => import("./components/FormLayoutDemo.vue"),
      },
      {
        path: "/input",
        name: "input",
        component: () => import("./components/InputDemo.vue"),
      },
      {
        path: "/floatlabel",
        name: "floatlabel",
        component: () => import("./components/FloatLabelDemo.vue"),
      },
      {
        path: "/invalidstate",
        name: "invalidstate",
        component: () => import("./components/InvalidStateDemo.vue"),
      },
      {
        path: "/button",
        name: "button",
        component: () => import("./components/ButtonDemo.vue"),
      },
      {
        path: "/table",
        name: "table",
        component: () => import("./components/TableDemo.vue"),
      },
      {
        path: "/list",
        name: "list",
        component: () => import("./components/ListDemo.vue"),
      },
      {
        path: "/tree",
        name: "tree",
        component: () => import("./components/TreeDemo.vue"),
      },
      {
        path: "/panel",
        name: "panel",
        component: () => import("./components/PanelsDemo.vue"),
      },
      {
        path: "/overlay",
        name: "overlay",
        component: () => import("./components/OverlayDemo.vue"),
      },
      {
        path: "/media",
        name: "media",
        component: () => import("./components/MediaDemo.vue"),
      },
      {
        path: "/menu",
        component: () => import("./components/MenuDemo.vue"),
        children: [
          {
            path: "",
            component: () => import("./components/menu/PersonalDemo.vue"),
          },
          {
            path: "/menu/seat",
            component: () => import("./components/menu/SeatDemo.vue"),
          },
          {
            path: "/menu/payment",
            component: () => import("./components/menu/PaymentDemo.vue"),
          },
          {
            path: "/menu/confirmation",
            component: () => import("./components/menu/ConfirmationDemo.vue"),
          },
        ],
      },
      {
        path: "/messages",
        name: "messages",
        component: () => import("./components/MessagesDemo.vue"),
      },
      {
        path: "/file",
        name: "file",
        component: () => import("./components/FileDemo.vue"),
      },
      {
        path: "/chart",
        name: "chart",
        component: () => import("./components/ChartDemo.vue"),
      },
      {
        path: "/misc",
        name: "misc",
        component: () => import("./components/MiscDemo.vue"),
      },
      {
        path: "/crud",
        name: "crud",
        component: () => import("./pages/CrudDemo.vue"),
      },
      {
        path: "/timeline",
        name: "timeline",
        component: () => import("./pages/TimelineDemo.vue"),
      },
      {
        path: "/empty",
        name: "empty",
        component: () => import("./components/EmptyPage.vue"),
      },
      {
        path: "/documentation",
        name: "documentation",
        component: () => import("./components/Documentation.vue"),
      },
      {
        path: "/blocks",
        name: "blocks",
        component: () => import("./components/BlocksDemo.vue"),
      },
      {
        path: "/icons",
        name: "icons",
        component: () => import("./components/IconsDemo.vue"),
      },
      {
        path: "/pages/invalid/access_denied",
        name: "access",
        component: () => import("./pages/invalid/AccessDenied.vue"),
      },
    ],
  },
  {
    path: "/pages/admin",
    name: "admin",
    component: AppAdmin,
    children: [
      {
        path: "/pages/admin/swagger",
        name: "swagger",
        component: () => import("./pages/admin/Swagger.vue"),
      },
    ],
  },
  {
    path: "/pages/auth/login",
    name: "login",
    component: () => import("./pages/auth/LoginPage.vue"),
  },
  {
    path: "/pages/auth/registration",
    name: "registration",
    component: () => import("./pages/auth/RegistrationPage.vue"),
  },
  {
    path: "/landing",
    name: "landing",
    component: () => import("./pages/LandingDemo.vue"),
  },
  {
    path: "/pages/invalid/error",
    name: "error",
    component: () => import("./pages/invalid/ErrorPage.vue"),
  },
  {
    // catch all
    // https://stackoverflow.com/a/40194152
    path: "/:pathMatch(.*)*",
    name: "not_found",
    component: () => import("./pages/invalid/NotFound.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
