<template>
  <BaseSection>
    <div class="card-tr mb-2">
      <div class="card-body">
        <p style="color: #ffffff; font-weight: bolder; font-size: 2em;" class="nav-link mb-0">
          Users </p>
      </div>
    </div>

    <div v-for="user in users" :key="user.id" class="card mb-2">
      <div class="card-body">
        <div class="d-flex mb-3">
          <div class="flex-shrink-0">
            <img src="../assets/u.png"
                 style="min-width: 50px; min-height: 50px; max-width: 50px;
                         max-height: 50px"
                 class="align-self-start mr-3 rounded-circle"
                 alt=""/>
          </div>
          <div class="flex-grow-1 ms-3" style="padding-top: 0.125rem">
            {{ user.login }} <span v-if="user.isAdmin" class="badge bg-warning text-dark mx-1">
            ADMIN</span>
<!--            <div v-if="user.isAdmin" class="text-muted small"> ADMIN-->
<!--            </div>-->
          </div>

          <div class="dropdown">
            <button class="btn btn-sm btn-reaction dropdown-toggle" type="button"
                    id="dropdownMenuButton2" data-bs-toggle="dropdown"
                    aria-expanded="false">
              Actions
            </button>
            <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton2">
              <li v-if="!user.isAdmin">
                <a class="dropdown-item small" href="#"
                   @click.prevent="grantAdminUser(user.login)">Grant admin</a>
              </li>
              <li><a class="dropdown-item small" href="#"
                     @click.prevent="deleteUser(user.login)">Delete</a></li>
            </ul>
          </div>

        </div>

        <p>{{ user.description }}</p>
      </div>
    </div>
  </BaseSection>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import BaseSection from '@/components/BaseSection.vue';
import { mapActions, mapGetters } from 'vuex';

export default defineComponent({
  name: 'UsersView',
  components: { BaseSection },
  mounted() {
    this.userInfo().then(() => {
      if (!this.isAdmin) {
        window.location.replace('/'); // TODO
      }

      this.viewUsers();
    });
  },
  methods: {
    ...mapActions(['viewUsers', 'deleteUser', 'grantAdminUser', 'userInfo']),
  },
  computed: {
    ...mapGetters(['isAdmin', 'users']),
    location() {
      return window.location.pathname;
    },
  },
});
</script>

<style>
.card  {
  border-color: #d7c1f1 !important;
}
</style>
