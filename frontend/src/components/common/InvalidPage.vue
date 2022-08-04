<template>
  <FormPage shadow="0" :logoColor="logoColor">
    <slot></slot>
    <div class="col-12 mt-5 text-center">
      <i
        class="pi pi-fw pi-arrow-left text-blue-500 mr-2"
        style="vertical-align: center"
      ></i>
      <a @click="$router.go(-1)" href="#" class="text-blue-500"> Go Back </a>
    </div>
  </FormPage>
</template>

<script>
export default {
  props: {
    logoColor: String,
  },
};
</script>
