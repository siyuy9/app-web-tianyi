<template>
  <FormPage>
    <div class="w-full md:w-10 mx-auto">
      <label for="email1" class="block text-900 text-xl font-medium mb-2"
        >Email</label
      >
      <InputText
        id="email1"
        v-model="email"
        type="text"
        class="w-full mb-3"
        placeholder="Email"
        style="padding: 1rem"
      />

      <label for="username1" class="block text-900 text-xl font-medium mb-2"
        >Username</label
      >
      <InputText
        id="username1"
        v-model="username"
        type="text"
        class="w-full mb-3"
        placeholder="Username"
        style="padding: 1rem"
      />

      <label for="password1" class="block text-900 font-medium text-xl mb-2"
        >Password</label
      >
      <Password
        id="password1"
        v-model="password"
        placeholder="Password"
        :toggleMask="true"
        class="w-full mb-3"
        inputClass="w-full"
        inputStyle="padding:1rem"
        :feedback="true"
      ></Password>

      <Button label="Register" class="w-full p-3 text-xl"></Button>
    </div>
  </FormPage>
</template>

<script>
export default {
  data() {
    return {
      username: "",
      password: "",
      remember_password: false,
    };
  },
};
</script>
