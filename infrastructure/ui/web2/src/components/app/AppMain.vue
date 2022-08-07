<template>
  <div :class="containerClass" @click="onWrapperClick">
    <AppTopBar v-if="$route.meta.showTopBar" />

    <div
      v-if="$route.meta.sidebarItems"
      class="layout-sidebar flex flex-column"
      :class="sidebarCustomZIndex"
      @click="onSidebarClick"
    >
      <div class="layout-menu-container flex-grow-1">
        <AppSubmenu
          :items="$route.meta.sidebarItems"
          class="layout-menu"
          :root="true"
          @menuitem-click="onMenuItemClick"
          @keydown="onMenuKeyDown"
        />
      </div>
      <Button
        class="p-button-rounded p-button-primary p-button-outlined align-self-center"
        :icon="sidebarToggleIconClass"
        @click="onMenuToggle($event)"
      />
    </div>

    <div
      v-if="$route.meta.sidebarItems && !isSidebarVisible"
      class="layout-sidebar-button-show-container"
    >
      <Button
        class="p-button-rounded p-button-primary"
        :icon="sidebarToggleIconClass"
        @click="onMenuToggle($event)"
      />
    </div>

    <router-view v-if="$route.meta.routerViewShowOutside" />

    <div class="layout-main-container" :class="contentClass">
      <div class="layout-main">
        <router-view v-if="!$route.meta.routerViewShowOutside" />
      </div>
      <AppFooter v-if="$route.meta.showFooter" />
    </div>

    <AppConfig :layoutMode="layoutMode" @layout-change="onLayoutChange" />
    <transition name="layout-mask">
      <div
        class="layout-mask p-component-overlay"
        v-if="mobileMenuActive"
      ></div>
    </transition>
  </div>
</template>

<script>
import AppTopBar from "./AppTopbar.vue";
import AppConfig from "./AppConfig.vue";
import AppFooter from "./AppFooter.vue";
import AppSubmenu from "./AppSubmenu.vue";
import EventBus from "../../lib/main/EventBus";

export default {
  emits: ["change-theme"],
  setup() {
    EventBus.on("app-toast-add", (toast_message) => {
      this.$toast.add(toast_message);
    });
  },
  data() {
    return {
      layoutMode: "static",
      staticMenuInactive: false,
      overlayMenuActive: false,
      mobileMenuActive: false,
    };
  },
  watch: {
    $route() {
      this.menuActive = false;
      this.$toast.removeAllGroups();
    },
  },
  methods: {
    onMenuItemClick(event) {
      this.$emit("menuitem-click", event);
    },
    onMenuKeyDown(event) {
      const nodeElement = event.target;
      if (event.code === "Enter" || event.code === "Space") {
        nodeElement.click();
        event.preventDefault();
      }
    },
    onWrapperClick() {
      if (!this.menuClick) {
        this.overlayMenuActive = false;
        this.mobileMenuActive = false;
      }

      this.menuClick = false;
    },
    menuToggle(event) {
      this.$emit("menu-toggle", event);
    },
    onMenuToggle(event) {
      this.menuClick = true;

      if (this.isDesktop()) {
        if (this.layoutMode === "overlay") {
          if (this.mobileMenuActive === true) {
            this.overlayMenuActive = true;
          }

          this.overlayMenuActive = !this.overlayMenuActive;
          this.mobileMenuActive = false;
        } else if (this.layoutMode === "static") {
          this.staticMenuInactive = !this.staticMenuInactive;
        }
      } else {
        this.mobileMenuActive = !this.mobileMenuActive;
      }

      event.preventDefault();
    },
    onSidebarClick() {
      this.menuClick = true;
    },
    onLayoutChange(layoutMode) {
      this.layoutMode = layoutMode;
    },
    addClass(element, className) {
      if (element.classList) element.classList.add(className);
      else element.className += " " + className;
    },
    removeClass(element, className) {
      if (element.classList) element.classList.remove(className);
      else
        element.className = element.className.replace(
          new RegExp(
            "(^|\\b)" + className.split(" ").join("|") + "(\\b|$)",
            "gi"
          ),
          " "
        );
    },
    isDesktop() {
      return window.innerWidth >= 992;
    },
  },
  computed: {
    sidebarToggleIconClass() {
      return (
        "pi " +
        (this.isSidebarVisible
          ? "pi-angle-double-left"
          : "pi-angle-double-right")
      );
    },
    // we need to dynamically add z-index, because if it's too small, then
    // it will no be able to cover menubar in the overlay mode,
    // and if it's too big, then in normal mode menubar popups will be covered
    sidebarCustomZIndex() {
      return this.overlayMenuActive ? "layout-sidebar-z-index-overlay" : "";
    },
    isSidebarVisible() {
      if (this.isDesktop()) {
        if (this.layoutMode === "static") return !this.staticMenuInactive;
        else if (this.layoutMode === "overlay") return this.overlayMenuActive;
      }

      return true;
    },
    containerClass() {
      return [
        "layout-wrapper",
        {
          "layout-overlay": this.layoutMode === "overlay",
          "layout-static": this.layoutMode === "static",
          "layout-static-sidebar-inactive":
            this.staticMenuInactive && this.layoutMode === "static",
          "layout-overlay-sidebar-active":
            this.overlayMenuActive && this.layoutMode === "overlay",
          "layout-mobile-sidebar-active": this.mobileMenuActive,
          "p-input-filled": this.$primevue.config.inputStyle === "filled",
          "p-ripple-disabled": this.$primevue.config.ripple === false,
          flex: this.$route.meta.containerClassAddFlex === true,
        },
      ];
    },
    logo() {
      return this.$appState.darkTheme
        ? "images/logo-white.svg"
        : "images/logo.svg";
    },
  },
  beforeUpdate() {
    if (this.mobileMenuActive)
      this.addClass(document.body, "body-overflow-hidden");
    else this.removeClass(document.body, "body-overflow-hidden");
  },
  components: {
    AppTopBar,
    AppConfig,
    AppFooter,
    AppSubmenu,
  },
};
</script>

<style lang="scss">
.p-toast.p-toast-top-right {
  z-index: 1000;
  top: 7rem;
}
</style>
