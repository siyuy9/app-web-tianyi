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
      <div class="font-medium text-3xl text-900">
        {{ project_name }}
      </div>
      <div class="flex align-items-center text-700 flex-wrap">
        <span>
          Project ID: {{ project_id }}
          <Button
            icon="pi pi-copy"
            class="p-button-text p-button-plain h-1rem w-2rem"
            v-tooltip.bottom="'Copy project id'"
            @click="writeToClipboard(project_id)"
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
        v-if="project_description"
        class="flex align-items-center text-700 flex-wrap mt-3"
      >
        <span>{{ project_description }}</span>
      </div>
    </div>
  </div>
  <router-view />
</template>

<script>
import { mapGetters } from "vuex";

import Error from "../../lib/main/Error";

export default {
  data() {
    return {};
  },
  async beforeMount() {
    try {
      // load the current project
      await this.$store.dispatch(
        "project/loadProject",
        this.$route.params.project_path
      );
    } catch (error) {
      Error(error);
    }
  },
  methods: {
    writeToClipboard: (text) => navigator.clipboard.writeText(text),
  },
  computed: {
    ...mapGetters("project", {
      project_id: "id",
      project_name: "name",
      project_description: "description",
    }),
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
};
</script>
