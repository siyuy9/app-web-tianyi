<template>
  <div class="grid">
    <div class="col-12">
      <div class="card">
        <div class="card" v-show="showSkeleton">
          <h4>Skeleton</h4>
          <div class="border-round border-1 surface-border p-4">
            <div class="flex mb-3">
              <Skeleton shape="circle" size="4rem" class="mr-2"></Skeleton>
              <div>
                <Skeleton width="10rem" class="mb-2"></Skeleton>
                <Skeleton width="5rem" class="mb-2"></Skeleton>
                <Skeleton height=".5rem"></Skeleton>
              </div>
            </div>
            <Skeleton width="100%" height="150px"></Skeleton>
            <div class="flex justify-content-between mt-3">
              <Skeleton width="4rem" height="2rem"></Skeleton>
              <Skeleton width="4rem" height="2rem"></Skeleton>
            </div>
          </div>
        </div>
        <iframe
          :src="swagger_src"
          width="100%"
          :height="outerHeight"
          @load="showSkeleton = false"
          style="border: 0"
        >
          Browser is not compatible
        </iframe>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      swagger_src: window.location.origin + "/api/v1/swagger/index.html",
      showSkeleton: true,
    };
  },
  computed: {
    outerHeight() {
      return window.outerHeight;
    },
  },
};
</script>
