<template>
  <BaseSection>
    <div class="card-tr mb-2">
      <div class="card-body">
        <p style="color: #ffffff; font-weight: bolder; font-size: 2em;" class="nav-link mb-0">
          Clients </p>
      </div>
    </div>

    <div v-if="clients.length">
      <div v-for="client in clients" :key="client.id" class="card mb-2">
        <div class="card-body">
          <div class="d-flex mb-0">
            <div class="flex-shrink-0">
              <img v-if="client.gender == 'female'" src="../assets/f.png"
                   style="min-width: 50px; min-height: 50px; max-width: 50px;
                         max-height: 50px"
                   class="align-self-start mr-3 rounded-circle"
                   alt=""/>
              <img v-if="client.gender == 'male'" src="../assets/m.png"
                   style="min-width: 50px; min-height: 50px; max-width: 50px;
                         max-height: 50px"
                   class="align-self-start mr-3 rounded-circle"
                   alt=""/>
            </div>
            <div class="flex-grow-1 ms-3" style="padding-top: 0.125rem">
              <div> {{ client.name }}  {{ client.surname }} {{ client.patronymic }}</div>
              <div class="text-muted small"> {{ client.email }}
              </div>
            </div>

            <div class="dropdown">
              <button class="btn btn-sm btn-reaction dropdown-toggle" type="button"
                      id="dropdownMenuButton3" data-bs-toggle="dropdown"
                      aria-expanded="false">
                Actions
              </button>
              <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton3">
                <li><a class="dropdown-item small" href="#"
                       @click.prevent="deleteClient(client.id)">Delete</a></li>
              </ul>
            </div>

          </div>
        </div>
        <ul class="list-group list-group-flush">
          <li class="list-group-item">
            <div class="row">
              <div class="col-4 fw-bold">Birth date:</div>
              <div class="col-8">{{ new Date(client.birth_date).toLocaleString() }}</div>
            </div>
          </li>
          <li class="list-group-item">
            <div class="row">
              <div class="col-4 fw-bold">Registration date:</div>
              <div class="col-8">{{ new Date(client.registration_date).toLocaleString() }}</div>
            </div>
          </li>
          <li class="list-group-item" v-for="(value, key) in client.data" :key="key">
            <div class="row">
              <div class="col-4 fw-bold">{{ key }}:</div>
              <div class="col-8">{{ value }}</div>
            </div>
          </li>
        </ul>
      </div>
    </div>
    <div v-else>
      <div class="card">
        <div class="card-body">
          No clients
        </div>
      </div>
    </div>
  </BaseSection>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import BaseSection from '@/components/BaseSection.vue';
import { mapActions, mapGetters } from 'vuex';

export default defineComponent({
  name: 'ClientsView',
  components: { BaseSection },
  mounted() {
    this.userInfo().then(() => {
      if (!this.isAdmin) {
        window.location.replace('/'); // TODO
      }

      this.viewClients();
    });
  },
  methods: {
    ...mapActions(['viewClients', 'deleteClient', 'userInfo']),
  },
  computed: {
    ...mapGetters(['isAdmin', 'clients']),
    location() {
      return window.location.pathname;
    },
  },
});
</script>
