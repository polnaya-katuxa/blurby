<template>
  <div class="card mb-2">

    <ul class="list-group list-group-flush">
      <li class="list-group-item">
        <div class="row">
          <div class="col-4 fw-bold">Name:</div>
          <div class="col-8">{{ et.name }}</div>
        </div>
      </li>

      <li class="list-group-item">
        <div class="row">
          <div class="col-4 fw-bold">Alias:</div>
          <div class="col-8">{{ et.alias }}</div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapActions } from 'vuex';

export default defineComponent({
  name: 'EventTypeCard',
  props: ['et'],
});
</script>
