<template>
  <Sidebar v-model:visible="active">
    <a href="#" class="layout-config-close" @click="hideConfigurator">
      <i class="pi pi-times"></i>
    </a>

    <div class="layout-config-content">
      <h4>Component Scale</h4>
      <div class="config-scale">
        <Button
          icon="pi pi-minus"
          @click="decrementScale()"
          class="p-button-text"
          :disabled="scale === scales[0]"
        />
        <i
          class="pi pi-circle-fill"
          v-for="s of scales"
          :class="{ 'scale-active': s === scale }"
          :key="s"
        />
        <Button
          icon="pi pi-plus"
          @click="incrementScale()"
          class="p-button-text"
          :disabled="scale === scales[scales.length - 1]"
        />
      </div>

      <AppInputStyleSwitch />

      <h4>Ripple Effect</h4>
      <InputSwitch
        :modelValue="rippleActive"
        @update:modelValue="onRippleChange"
      />

      <h4>Menu Type</h4>
      <div class="p-formgroup-inline">
        <div class="field-radiobutton">
          <RadioButton
            id="static"
            name="layoutMode"
            value="static"
            v-model="d_layoutMode"
            @change="changeLayout($event, 'static')"
          />
          <label for="static">Static</label>
        </div>
        <div class="field-radiobutton">
          <RadioButton
            id="overlay"
            name="layoutMode"
            value="overlay"
            v-model="d_layoutMode"
            @change="changeLayout($event, 'overlay')"
          />
          <label for="overlay">Overlay</label>
        </div>
      </div>

      <h4>Themes</h4>
      <div v-for="item in computedThemes" :key="item.title">
        <h5>{{ item.title }}</h5>
        <div class="grid free-themes">
          <div
            class="col-3"
            v-for="item_inner in item.themes"
            :key="item_inner.name"
          >
            <p>
              <button
                class="p-link"
                type="button"
                @click="changeTheme($event, item_inner.name, item_inner.dark)"
              >
                <img :src="item_inner.source" :alt="item_inner.name" />
              </button>
              <span>{{ item_inner.title }}</span>
            </p>
          </div>
        </div>
      </div>
    </div>
  </Sidebar>
</template>

<script>
import EventBus from "../../lib/main/EventBus";

export default {
  props: {
    inputStyle: String,
    layoutMode: {
      type: String,
      default: null,
    },
  },
  outsideClickListener: null,
  themeChangeListener: null,
  watch: {
    $route() {
      if (this.active) {
        this.active = false;
        this.unbindOutsideClickListener();
      }
    },
  },
  beforeUnmount() {
    EventBus.off("theme-change", this.themeChangeListener);
  },
  mounted() {
    this.themeChangeListener = (event) => {
      if (event.theme === "nano") this.scale = 12;
      else this.scale = 14;

      this.applyScale();
    };
    EventBus.on("configurator-toggle", this.toggleConfigurator);
    EventBus.on("theme-change", this.themeChangeListener);
  },
  methods: {
    toggleConfigurator(event) {
      this.active = !this.active;
      // have to do this, otherwise listener will just hide theme sidebar
      // immediately, since the initial click is coming from the outside
      event.stopPropagation();
      event.preventDefault();

      if (this.active) {
        this.bindOutsideClickListener();
      } else {
        this.unbindOutsideClickListener();
      }
    },
    hideConfigurator(event) {
      this.active = false;
      this.unbindOutsideClickListener();
      event.preventDefault();
    },
    changeTheme(event, theme, dark) {
      EventBus.emit("theme-change", { theme: theme, dark: dark });
      event.preventDefault();
    },
    changeLayout(event, layoutMode) {
      this.$emit("layout-change", layoutMode);
      event.preventDefault();
    },
    bindOutsideClickListener() {
      if (!this.outsideClickListener) {
        this.outsideClickListener = (event) => {
          if (this.active && this.isOutsideClicked(event)) {
            this.active = false;
          }
        };
        document.addEventListener("click", this.outsideClickListener);
      }
    },
    unbindOutsideClickListener() {
      if (this.outsideClickListener) {
        document.removeEventListener("click", this.outsideClickListener);
        this.outsideClickListener = null;
      }
    },
    isOutsideClicked(event) {
      return !(
        this.$el.isSameNode(event.target) || this.$el.contains(event.target)
      );
    },
    decrementScale() {
      this.scale--;
      this.applyScale();
    },
    incrementScale() {
      this.scale++;
      this.applyScale();
    },
    applyScale() {
      document.documentElement.style.fontSize = this.scale + "px";
    },
    onRippleChange(value) {
      this.$primevue.config.ripple = value;
    },
  },
  computed: {
    containerClass() {
      return ["layout-config", { "layout-config-active": this.active }];
    },
    rippleActive() {
      return this.$primevue.config.ripple;
    },
    // have to normalize the list, because v-for does not allow to call methods
    // inside a template
    // https://forum.vuejs.org/t/call-method-inside-a-v-for-loop/37790
    computedThemes() {
      return this.$data.data_themes.map((theme_cluster) => {
        theme_cluster.themes = theme_cluster.themes.map((theme) => {
          theme.title = theme.name
            .split("-")
            .splice(1)
            .map((word) => word[0].toUpperCase() + word.substr(1))
            .join(" ");
          var source_name = theme.source_name ? theme.source_name : theme.name;
          theme.source =
            "/images/themes/" + source_name + "." + theme.extension;
          return theme;
        });
        return theme_cluster;
      });
    },
  },
  data() {
    return {
      active: false,
      scale: 14,
      scales: [12, 13, 14, 15, 16],
      d_layoutMode: this.layoutMode,
      data_themes: [
        {
          title: "Bootstrap",
          themes: [
            { name: "bootstrap4-light-blue", extension: "svg" },
            { name: "bootstrap4-light-purple", extension: "svg" },

            { name: "bootstrap4-dark-blue", extension: "svg", dark: true },
            { name: "bootstrap4-dark-purple", extension: "svg", dark: true },
          ],
        },
        {
          title: "PrimeOne Design - 2022",
          themes: [
            { name: "lara-light-indigo", extension: "png" },
            { name: "lara-light-blue", extension: "png" },
            { name: "lara-light-purple", extension: "png" },
            { name: "lara-light-teal", extension: "png" },

            { name: "lara-dark-indigo", extension: "png", dark: true },
            { name: "lara-dark-blue", extension: "png", dark: true },
            { name: "lara-dark-purple", extension: "png", dark: true },
            { name: "lara-dark-teal", extension: "png", dark: true },
          ],
        },
        {
          title: "Material Design",
          themes: [
            { name: "md-light-indigo", extension: "svg" },
            { name: "md-light-deeppurple", extension: "svg" },

            { name: "md-dark-indigo", extension: "svg", dark: true },
            { name: "md-dark-deeppurple", extension: "svg", dark: true },
          ],
        },
        {
          title: "Material Design Compact",
          themes: [
            {
              name: "mdc-light-indigo",
              extension: "svg",
              source_name: "md-light-indigo",
            },
            {
              name: "mdc-light-deeppurple",
              extension: "svg",
              source_name: "md-light-deeppurple",
            },

            {
              name: "mdc-dark-indigo",
              extension: "svg",
              source_name: "md-dark-indigo",
              dark: true,
            },
            {
              name: "mdc-dark-deeppurple",
              extension: "svg",
              source_name: "md-dark-deeppurple",
              dark: true,
            },
          ],
        },
        {
          title: "PrimeOne Design - 2021",
          themes: [
            { name: "saga-blue", extension: "png" },
            { name: "saga-green", extension: "png" },
            { name: "saga-orange", extension: "png" },
            { name: "saga-purple", extension: "png" },

            { name: "vela-blue", extension: "png", dark: true },
            { name: "vela-green", extension: "png", dark: true },
            { name: "vela-orange", extension: "png", dark: true },
            { name: "vela-purple", extension: "png", dark: true },

            { name: "arya-blue", extension: "png", dark: true },
            { name: "arya-green", extension: "png", dark: true },
            { name: "arya-orange", extension: "png", dark: true },
            { name: "arya-purple", extension: "png", dark: true },
          ],
        },
        {
          title: "Tailwind",
          themes: [{ name: "tailwind-light", extension: "png" }],
        },
        {
          title: "Fluent UI",
          themes: [{ name: "fluent-light", extension: "png" }],
        },
      ],
    };
  },
};
</script>
