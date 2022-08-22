<template>
  <div class="card">
    <div class="flex justify-content-between">
      <div class="font-medium text-3xl text-900 mb-3">{{ pipeline.name }}</div>
      <Button
        label="Run"
        class="p-button-outlined"
        :loading="pipelineLaunched"
        @click="createPipeline"
      ></Button>
    </div>
    <div class="text-500 mb-5">
      {{ pipeline.description ? pipeline.description : "no description" }}
    </div>
    <DataTable :value="jobs" responsiveLayout="scroll">
      <Column field="name" header="Name">
        <template #body="slotProps">
          <span>{{ slotProps.data.name }}</span>
        </template>
      </Column>
    </DataTable>
  </div>
</template>

<script>
import axios from "axios";
import { mapGetters } from "vuex";
import Error from "../../lib/main/Error";
import EventBus from "../../lib/main/EventBus";

export default {
  data() {
    return {
      pipelineData: {},
      pipelineLaunched: false,
    };
  },
  computed: {
    pipeline() {
      return this.pipelineData;
    },
    ...mapGetters("project", {
      branch: "branch",
      project_id: "id",
    }),
    jobs() {
      return this.branch.config.jobs;
    },
  },
  methods: {
    createPipeline() {
      if (this.pipelineLaunched) {
        return;
      }
      this.pipelineLaunched = true;
      axios
        .post(
          `/api/v1/projects/${this.project_id}/branches/${this.$route.params.pipeline_branch}/pipelines/${this.$route.params.pipeline_name}`
        )
        .then((response) =>
          EventBus.emit("app-toast-add", {
            severity: "success",
            summary: "pipeline launched",
            detail: response.data.web_url,
            life: 3000,
          })
        )
        .catch(Error)
        .finally(() => (this.pipelineLaunched = false));
    },
  },
  async beforeMount() {
    try {
      // load the current project
      await this.$store.dispatch(
        "project/loadProject",
        this.$route.params.project_path
      );
      await this.$store.dispatch(
        "project/loadBranch",
        this.$route.params.pipeline_branch
      );
      this.branch.config.pipelines.some((element) => {
        if (element.name !== this.$route.params.pipeline_name) {
          return false;
        }
        this.pipelineData = element;
        return true;
      });
    } catch (error) {
      Error(error);
    }
  },
};
</script>
