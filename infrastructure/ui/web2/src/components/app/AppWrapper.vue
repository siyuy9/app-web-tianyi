<template>
  <router-view></router-view>
</template>

<script>
import { mapState } from "vuex";

export default {
  computed: mapState("theme", {
    theme: "theme",
  }),
  watch: {
    theme(newValue, oldValue) {
      const elementId = "theme-link";
      const linkElement = document.getElementById(elementId);
      const cloneLinkElement = linkElement.cloneNode(true);
      const newThemeUrl = linkElement
        .getAttribute("href")
        .replace(oldValue.name, newValue.name);

      cloneLinkElement.setAttribute("id", elementId + "-clone");
      cloneLinkElement.setAttribute("href", newThemeUrl);
      cloneLinkElement.addEventListener("load", () => {
        linkElement.remove();
        cloneLinkElement.setAttribute("id", elementId);
      });
      linkElement.parentNode.insertBefore(
        cloneLinkElement,
        linkElement.nextSibling
      );
    },
  },
};
</script>
