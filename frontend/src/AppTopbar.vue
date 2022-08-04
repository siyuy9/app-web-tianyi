<template>
  <div class="layout-topbar">
    <router-link :to="{ name: 'root' }" class="layout-topbar-logo">
      <img alt="Logo" :src="topbarImage()" />
    </router-link>
    <button
      class="p-link layout-menu-button layout-topbar-button"
      @click="onMenuToggle($event)"
    >
      <i class="pi" :class="menuToggleIcon"></i>
    </button>

    <Button
      icon="pi pi-bars"
      label="Menu"
      class="p-button-lg p-button-text p-button-plain"
      @click="toggle('topbar-menu-main', $event)"
    />
    <TieredMenu
      :model="topbarMenuMainItems"
      :popup="true"
      ref="topbar-menu-main"
    />

    <ul class="layout-topbar-menu hidden lg:flex origin-top">
      <li>
        <button class="p-link layout-topbar-button">
          <i class="pi pi-calendar"></i>
          <span>Events</span>
        </button>
      </li>
      <li>
        <a
          href="#"
          class="p-link layout-topbar-button layout-config-button"
          id="layout-config-button"
          @click.stop="toggleConfigurator"
        >
          <i class="pi pi-cog"></i>
          <span>Settings</span>
        </a>
      </li>
      <li>
        <button class="p-link layout-topbar-button">
          <i class="pi pi-user"></i>
          <span>Profile</span>
        </button>
      </li>
    </ul>
  </div>
</template>

<script>
import EventBus from "./AppEventBus";

export default {
  props: {
    isSidebarVisible: Boolean,
  },
  methods: {
    toggleConfigurator(event) {
      EventBus.emit("configurator-toggle", event);
    },
    onMenuToggle(event) {
      this.$emit("menu-toggle", event);
    },
    onTopbarMenuToggle(event) {
      this.$emit("topbar-menu-toggle", event);
    },
    topbarImage() {
      return this.$appState.darkTheme
        ? "/images/logo-white.svg"
        : "/images/logo-dark.svg";
    },
    toggle(ref, event) {
      this.$refs[ref].toggle(event);
    },
  },
  computed: {
    darkTheme() {
      return this.$appState.darkTheme;
    },
    menuToggleIcon() {
      return this.$props.isSidebarVisible
        ? "pi-angle-double-left"
        : "pi-angle-double-right";
    },
  },
  data() {
    return {
      topbarMenuMainItems: [
        {
          label: "Projects",
          icon: "pi pi-folder",
          to: { name: "projects" },
        },
        {
          label: "Admin",
          icon: "pi pi-save",
          to: { name: "admin" },
        },
      ],
    };
  },
};
</script>
