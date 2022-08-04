<template>
  <App>
    <AppMenu :model="menu" @menuitem-click="onMenuItemClick($event)" />
  </App>
</template>

<script>
import App from "./App.vue";
import AppMenu from "./AppMenu.vue";

export default {
  methods: {
    onMenuItemClick(event) {
      if (event.item && !event.item.items) {
        this.overlayMenuActive = false;
        this.mobileMenuActive = false;
      }
    },
  },
  components: {
    App,
    AppMenu,
  },
  data() {
    return {
      menu: [
        {
          label: "API",
          items: [
            {
              label: "Swagger",
              icon: "pi pi-fw pi-home",
              to: { name: "swagger" },
            },
          ],
        },
      ],
    };
  },
};
</script>
