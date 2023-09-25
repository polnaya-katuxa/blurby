<template>
  <section class="gradient-custom-2" style="min-height: 100vh">
    <div class="container-fluid py-4 h-100">
      <div class="row d-flex justify-content-center h-100">
        <div class="col col-lg-9 col-xl-7">

          <NavBar/>

          <div class="container">
            <br>

            <slot></slot>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import NavBar from '@/components/NavBar.vue';

export default defineComponent({
  name: 'BaseSection',
  components: {
    NavBar,
  },
});
</script>
