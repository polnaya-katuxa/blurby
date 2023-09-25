<template>
  <div class="notification-block">
    <div v-for="error in errors" :key="error"
         class="alert alert-danger alert-dismissible fade show" role="alert">
      <strong>Error occurred. </strong> {{ error }}
      <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"
              @click="deleteError(error)"></button>
    </div>
    <div v-for="info in infos" :key="info"
         class="alert alert-info alert-dismissible fade show" role="alert">
      <strong>Info message. </strong> {{ info }}
      <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"
              @click="deleteInfo(info)"></button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapActions, mapGetters } from 'vuex';

export default defineComponent({
  name: 'NotificationsBlock',
  methods: {
    ...mapActions(['addError', 'deleteError', 'addInfo', 'deleteInfo']),
  },
  computed: {
    ...mapGetters(['errors', 'infos']),
  },
});
</script>
