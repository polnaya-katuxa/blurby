<template>
  <div>
    <router-view/>
    <NotificationsBlock/>
  </div>
</template>

<style>

</style>

<script lang="ts">
import { defineComponent } from 'vue';
import Cookies from 'cookies-ts';
import NotificationsBlock from '@/components/NotificationBlock.vue';

export default defineComponent({
  name: 'App',
  components: {
    NotificationsBlock,
  },
  beforeCreate() {
    const cookies = new Cookies();
    const curURL = window.location.pathname;
    const f = cookies.get('user-token');

    if ((curURL !== '/register') && (curURL !== '/login') && (f === null)) {
      window.location.replace('/login');
    }
  },
});
</script>
