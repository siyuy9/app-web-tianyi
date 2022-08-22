<template>
  <div class="card">
    <Dropdown
      v-model="currentBranchData"
      :options="branches"
      optionLabel="branch name"
      :filter="true"
      placeholder="Select a branch"
    >
      <template #value="slotProps">
        <div v-if="slotProps.value">
          {{ slotProps.value.name }}
        </div>
        <span v-else>
          {{ slotProps.placeholder }}
        </span>
      </template>
      <template #option="slotProps">
        <div>{{ slotProps.option.name }}</div>
      </template>
    </Dropdown>
    <Button
      label="Update"
      class="p-button-outlined mx-3"
      :loading="branchUpdating"
      @click="updateBranch"
    ></Button>
  </div>
  <DataTable :value="pipelines" responsiveLayout="scroll" class="card">
    <Column field="name" header="Name">
      <template #body="slotProps">
        <router-link
          :to="{
            name: 'project_pipeline',
            params: {
              pipeline_branch: currentBranch.name,
              pipeline_name: slotProps.data.name,
            },
          }"
        >
          <span>{{ slotProps.data.name }}</span>
        </router-link>
      </template>
    </Column>
  </DataTable>
</template>

<script>
import { mapGetters } from "vuex";

import axios from "axios";
import Error from "../../lib/main/Error";
import EventBus from "../../lib/main/EventBus";

export default {
  data() {
    return {
      currentBranchData: {},
      branches: [],
      branchUpdating: false,
    };
  },
  computed: {
    ...mapGetters("project", {
      branch: "branch",
      default_branch: "default_branch",
      project_id: "id",
    }),
    currentBranch() {
      return this.currentBranchData;
    },
    pipelines() {
      return this.branch.config ? this.branch.config.pipelines : [];
    },
  },
  watch: {
    async currentBranchData(newBranch, oldBranch) {
      try {
        // load the default branch
        await this.loadBranch(newBranch.name);
      } catch (error) {
        Error(error);
        this.currentBranchData = oldBranch;
      }
    },
  },
  methods: {
    loadBranch(branchName) {
      return this.$store.dispatch("project/loadBranch", branchName);
    },
    updateBranch() {
      if (this.branchUpdating) {
        return;
      }
      this.branchUpdating = true;
      axios
        .put(
          `/api/v1/projects/${this.project_id}/branches/${this.currentBranch.name}`
        )
        .then(() =>
          this.loadBranch(this.currentBranch.name).then(() =>
            EventBus.emit("app-toast-add", {
              severity: "success",
              summary: "branch updated",
              life: 5000,
            })
          )
        )
        .catch(Error)
        .finally(() => (this.branchUpdating = false));
    },
  },
  async beforeMount() {
    try {
      // load the current project
      await this.$store.dispatch(
        "project/loadProject",
        this.$route.params.project_path
      );
      this.currentBranchData = {
        name: this.default_branch,
        value: this.default_branch,
      };
      this.branches.push(this.currentBranchData);
    } catch (error) {
      Error(error);
    }
  },
};
</script>
