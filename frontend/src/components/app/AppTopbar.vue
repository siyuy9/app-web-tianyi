<template>
  <!--
      Menubar documentation - https://www.primefaces.org/primevue/menubar
  -->
  <Menubar :model="topbarLeftMenuItems" class="layout-topbar">
    <template #start>
      <div style="display: flex">
        <router-link :to="{ name: 'root' }" class="layout-topbar-logo p-1">
          <img alt="Logo" :src="topbarImage" />
        </router-link>
      </div>
    </template>
    <template #end>
      <span class="p-input-icon-left">
        <i class="pi pi-search" />
        <InputText type="text" v-model="topbarSearch" placeholder="Search" />
      </span>
    </template>
  </Menubar>
</template>

<script>
import EventBus from "../../lib/app/AppEventBus";
import {
  AdminSidebarItems,
  BrowseSidebarItems,
} from "../../lib/main/SidebarItems";

export default {
  computed: {
    topbarImage() {
      return this.$appState.darkTheme
        ? "/images/logo-white.svg"
        : "/images/logo-dark.svg";
    },
  },
  data() {
    return {
      topbarSearch: "",
      topbarLeftMenuItems: [
        {
          label: "Browse",
          icon: "pi pi-bars",
          items: BrowseSidebarItems,
        },
        {
          label: "Admin",
          icon: "pi pi-shield",
          items: AdminSidebarItems,
        },
        {
          label: "Profile",
          icon: "pi pi-user",
          items: [
            {
              label: "Theme",
              icon: "pi pi-image",
              command: (event) => {
                // toggle theme configurator on the right
                // AppConfig.vue listens to this event
                EventBus.emit("configurator-toggle", event.originalEvent);
              },
            },
            {
              separator: true,
            },
            {
              label: "Logout",
              icon: "pi pi-sign-out",
            },
          ],
        },
        {
          label: "Help",
          icon: "pi pi-question-circle",
          items: [
            {
              label: "Source code",
              icon: "pi pi-external-link",
              url: "https://gitlab.com/kongrentian-groups/golang/tianyi",
            },
          ],
        },
      ],
    };
  },
};
</script>
