<template>
  <FormPage>
    <div class="text-center mb-5">
      <div class="text-900 text-3xl font-medium mb-3">Welcome Back</div>
      <span class="text-600 font-medium line-height-3"
        >Don't have an account?</span
      >
      <router-link
        class="font-medium no-underline ml-2 text-blue-500 cursor-pointer"
        :to="{ name: 'registration' }"
      >
        Create today!
      </router-link>
    </div>

    <div class="w-full md:w-10 mx-auto">
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
        :feedback="false"
      ></Password>

      <div class="flex align-items-center justify-content-between mb-5">
        <div class="flex align-items-center">
          <Checkbox
            id="rememberme1"
            v-model="remember_password"
            :binary="true"
            class="mr-2"
          ></Checkbox>
          <label for="rememberme1">Remember me</label>
        </div>
        <a
          class="font-medium no-underline ml-2 text-right cursor-pointer"
          style="color: var(--primary-color)"
          >Forgot password?</a
        >
      </div>
      <Button label="Sign In" class="w-full p-3 text-xl"></Button>
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
