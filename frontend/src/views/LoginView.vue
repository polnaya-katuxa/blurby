<template>
  <div class="position-relative vh-100">

    <div class="row" style="margin-top: 180px">
      <div class="col-md-4">

      </div>

      <div class="col-md-4 px-5">
        <form class="top-50 needs-validation"
              @submit.prevent="onSubmit">

          <div class="d-flex justify-content-center">
            <img src="../assets/logo.png"
                 style="min-width: 160px; min-height: 160px; max-width: 160px;
                         max-height: 160px"
                 alt=""/>
          </div>

          <h3>Login</h3>
          <div class="form-floating mb-3 has-validation">
            <input type="text" pattern=".{4,20}" required minlength="4" class="form-control"
                   id="inputLogin" placeholder="Your login" v-model="loginForm">
            <!-- eslint-disable-next-line -->
            <label for="inputLogin">
              Login
            </label>
            <div class="invalid-feedback">
              Please choose a login.
            </div>
          </div>

          <div class="form-floating mb-3">
            <input type="password" pattern=".{8,20}" required minlength="8"
                   title="8 characters minimum"
                   class="form-control" id="inputPassword" placeholder="Your password"
                   v-model="passwordForm">
            <!-- eslint-disable-next-line -->
            <label for="inputPassword">
              Password
            </label>
          </div>

          <div class="row w-100 me-0 ms-0">
            <button type="submit" class="btn btn-pink col">Submit</button>
            <router-link to="/register" class="col btn btn-pink-link">
              Not registered yet?
            </router-link>
          </div>
        </form>
      </div>

      <div class="col-md-4">

      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapActions } from 'vuex';

export default defineComponent({
  name: 'LoginView',
  data() {
    return {
      loginForm: '',
      passwordForm: '',
    };
  },
  methods: {
    ...mapActions(['login']),
    onSubmit() {
      console.error(this.loginForm, this.passwordForm);
      this.login({
        login: this.loginForm,
        password: this.passwordForm,
      });
    },
  },
});
</script>

<style scoped>

</style>
