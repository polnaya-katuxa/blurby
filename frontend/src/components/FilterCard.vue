<template>
  <div class="card mb-4">
    <div v-if="location != '/ads'" class="card-header">
      <div class="d-flex justify-content-between">
        <div></div>
        <div>
          <button type="button" id="closeFilter" class="btn-close me-0" @click.prevent="onDelete"
                  aria-label="Delete"></button>
        </div>
      </div>
    </div>

    <ul class="list-group list-group-flush">
      <li class="list-group-item">
        <div class="row">
          <div class="col-4 fw-bold">Type:</div>
          <div class="col-8">{{ filter.type }}</div>
        </div>
      </li>

      <div v-if="filter.type == 'by event'">
        <li class="list-group-item">
          <div class="row">
            <div class="col-4 fw-bold">Alias:</div>
            <div class="col-8">{{ filter.filter.alias }}</div>
          </div>
        </li>

        <li class="list-group-item">
          <div class="row">
            <div class="col-4 fw-bold">Span:</div>
            <div class="col-8">{{ filter.filter.span }}</div>
          </div>
        </li>

        <li class="list-group-item">
          <div class="row">
            <div class="col-4 fw-bold">Rate:</div>
            <div class="col-8">{{ filter.filter.rate }}</div>
          </div>
        </li>
      </div>

      <div v-else-if="filter.type == 'by field'">
        <li class="list-group-item">
          <div class="row">
            <div class="col-4 fw-bold">Field:</div>
            <div class="col-8">{{ filter.filter.field }}</div>
          </div>
        </li>

        <li class="list-group-item">
          <div class="row">
            <div class="col-4 fw-bold">Comparison:</div>
            <div class="col-8">{{ filter.filter.cmp }}</div>
          </div>
        </li>

        <li class="list-group-item">
          <div class="row">
            <div class="col-4 fw-bold">Value #1:</div>
            <div class="col-8">{{ filter.filter.value1 }}</div>
          </div>
        </li>

        <li v-if="filter.filter.cmp == 'between'" class="list-group-item">
          <div class="row">
            <div class="col-4 fw-bold">Value #2:</div>
            <div class="col-8">{{ filter.filter.value2 }}</div>
          </div>
        </li>
      </div>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapActions, mapGetters } from 'vuex';

export default defineComponent({
  name: 'FilterCard',
  props: ['filter', 'ind'],
  methods: {
    ...mapActions(['deleteFilter']),
    onDelete() {
      this.deleteFilter(this.ind);
    },
  },
  computed: {
    location() {
      return window.location.pathname;
    },
  },
});
</script>
