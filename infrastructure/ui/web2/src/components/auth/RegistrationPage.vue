<template>
  <form
    class="p-fluid w-full md:w-10 mx-auto"
    @submit.prevent="handleSubmit(validator.$invalid)"
  >
    <label for="email1" class="block text-900 text-xl font-medium mb-2"
      >Email</label
    >
    <InputText
      id="email1"
      v-model="validator.email.$model"
      type="text"
      class="w-full mb-3"
      placeholder="Email"
      style="padding: 1rem"
      :class="{ 'p-invalid': validator.email.$invalid && submitted }"
    />

    <label for="username1" class="block text-900 text-xl font-medium mb-2"
      >Username</label
    >
    <InputText
      id="username1"
      v-model="validator.username.$model"
      type="text"
      class="w-full mb-3"
      placeholder="Username"
      style="padding: 1rem"
      :class="{ 'p-invalid': validator.username.$invalid && submitted }"
    />

    <label for="password1" class="block text-900 font-medium text-xl mb-2"
      >Password</label
    >
    <Password
      id="password1"
      v-model="validator.password.$model"
      placeholder="Password"
      :toggleMask="true"
      class="w-full mb-3"
      inputClass="w-full"
      inputStyle="padding:1rem"
      :feedback="true"
      :class="{ 'p-invalid': validator.password.$invalid && submitted }"
    ></Password>

    <Button label="Register" type="submit" class="w-full p-3 text-xl"></Button>
  </form>
</template>

<script>
import { email, required } from "@vuelidate/validators";
import { useVuelidate } from "@vuelidate/core";

export default {
  setup: () => ({ validator: useVuelidate() }),
  data() {
    return {
      username: "",
      password: "",
      email: "",
    };
  },
  validations() {
    return {
      username: {
        required,
      },
      email: {
        required,
        email,
      },
      password: {
        required,
      },
    };
  },
  methods: {
    handleSubmit(isFormInvalid) {
      this.submitted = true;

      if (isFormInvalid) {
        return;
      }

      //this.toggleDialog();
    },
  },
};
</script>
