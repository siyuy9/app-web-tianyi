<template>
  <InvalidPage logoColor="orange">
    <div class="grid flex flex-column align-items-center">
      <div
        class="flex justify-content-center align-items-center bg-orange-500 border-circle"
        style="width: 3.2rem; height: 3.2rem"
      >
        <i class="pi pi-fw pi-lock text-2xl text-50"></i>
      </div>
      <h1 class="text-900 font-bold text-4xl lg:text-5xl mb-2">
        Access Denied
      </h1>
      <span class="text-600 text-center"
        >You do not have the necesary permisions. Please contact admins.</span
      >
      <img
        src="/layout/images/asset-access.svg"
        alt="Access denied"
        class="mt-5"
        :width="width"
      />
    </div>
  </InvalidPage>
</template>

<script>
export default {
  computed: {
    width: "80%",
  },
};
</script>
