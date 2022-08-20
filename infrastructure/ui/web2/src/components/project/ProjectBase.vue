<template>
  <Breadcrumb
    :home="{ icon: 'pi pi-home', to: { name: 'root' } }"
    :model="path"
    class="card"
  />
  <div
    class="card flex align-items-start flex-column lg:justify-content-between lg:flex-row"
  >
    <div>
      <div class="font-medium text-3xl text-900">{{ project.name }}</div>
      <div class="flex align-items-center text-700 flex-wrap">
        <span>
          Project ID: {{ project.id }}
          <Button
            icon="pi pi-copy"
            class="p-button-text p-button-plain h-1rem w-2rem"
            v-tooltip.bottom="'Copy project id'"
            @click="navigator.clipboard.writeText(project.id)"
          />
        </span>
      </div>
      <div class="flex align-items-center text-700 flex-wrap">
        <div class="mr-5 flex align-items-center mt-3">
          <i class="pi pi-globe mr-2"></i>
          <span>placeholder</span>
        </div>
        <div class="flex align-items-center mt-3">
          <i class="pi pi-clock mr-2"></i>
          <span>placeholder</span>
        </div>
      </div>
      <div
        v-if="project.description"
        class="flex align-items-center text-700 flex-wrap mt-3"
      >
        <span>{{ project.description }}</span>
      </div>
    </div>
    <div class="mt-3 lg:mt-0">
      <Button
        label="placeholder"
        class="p-button-outlined mr-2"
        icon="pi pi-check"
      ></Button>
    </div>
  </div>
  <router-view />
</template>

<script>
export default {
  data() {
    return {
      project: {},
    };
  },
  beforeMount() {
    this.$data.project = this.$store.getters["project/project"](
      this.$route.params.project_path
    );
  },
  computed: {
    path() {
      var oldPath = "";
      return this.$route.params.project_path.split("/").map((element) => {
        return {
          label: element,
          to: {
            name: "project",
            params: {
              project_path: oldPath ? [oldPath, element].join("/") : element,
            },
          },
        };
      });
    },
  },
  methods: {},
};
</script>
