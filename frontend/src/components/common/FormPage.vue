<template>
  <div
    class="surface-0 flex align-items-center justify-content-center min-h-screen min-w-screen overflow-hidden"
  >
    <div class="grid justify-content-center p-2 lg:p-0" style="min-width: 80%">
      <div
        class="surface-card p-4 border-round w-full lg:w-6"
        :class="shadowClass"
      >
        <div
          class="col-12"
          style="
            border-radius: 56px;
            padding: 0.3rem;
            background: linear-gradient(
              180deg,
              var(--primary-color),
              rgba(33, 150, 243, 0) 30%
            );
          "
        >
          <div
            class="h-full w-full m-0 py-7 px-4"
            style="
              border-radius: 53px;
              background: linear-gradient(
                180deg,
                var(--surface-50) 38.9%,
                var(--surface-0)
              );
            "
          >
            <div class="text-center mb-5">
              <img
                :src="logoSrc"
                alt="Tianyi logo"
                class="mb-5"
                style="width: 81px; height: 60px"
              />
            </div>
            <slot></slot>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    shadowClass() {
      return "shadow-" + this.shadow;
    },
    logoSrc() {
      var color;
      if (this.logoColor) {
        color = this.logoColor;
      } else if (this.$appState.darkTheme) {
        color = "white";
      } else {
        color = "dark";
      }
      return "/layout/images/logo-" + color + ".svg";
    },
  },
  props: {
    shadow: {
      type: Number,
      default: 2,
    },
    logoColor: {
      type: String,
    },
  },
};
</script>

<style lang="scss">
.pi-eye {
  transform: scale(1.6);
  margin-right: 1rem;
}

.pi-eye-slash {
  transform: scale(1.6);
  margin-right: 1rem;
}
</style>
