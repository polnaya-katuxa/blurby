<template>

  <div v-if="!ad.schedule.finished" class="card mb-2">
    <div class="card-body">

      <p style="color: #000000; font-weight: bold"> Content </p>

      <div class="card mb-4 px-2 py-2">
        <div v-html="ad.content"></div>
      </div>

      <div class="row">
        <div class="col-4 fw-bold">Create time</div>
        <div class="col-8">{{ new Date(ad.create_time).toLocaleString() }}</div>
      </div>

      <br>

      <p style="color: #000000; font-weight: bold"> Schedule </p>

      <div class="card mb-4">
        <ul class="list-group list-group-flush">
          <li class="list-group-item">
            <div class="row">
              <div class="col-4 fw-bold">Finished:</div>
              <div v-if="ad.schedule.finished" class="col-8">{{ ad.schedule.finished }}</div>
              <div v-else class="col-8">false</div>
            </div>
          </li>

          <li class="list-group-item">
            <div class="row">
              <div class="col-4 fw-bold">Periodic:</div>
              <div v-if="ad.schedule.periodic" class="col-8">{{ ad.schedule.periodic }}</div>
              <div v-else class="col-8">false</div>
            </div>
          </li>

          <li class="list-group-item">
            <div class="row">
              <div class="col-4 fw-bold">Span:</div>
              <div class="col-8">{{ ad.schedule.span }}</div>
            </div>
          </li>

          <li class="list-group-item">
            <div class="row">
              <div class="col-4 fw-bold">Next time:</div>
              <div class="col-8">{{ new Date(ad.schedule.nextTime).toLocaleString() }}</div>
            </div>
          </li>
        </ul>
      </div>

      <p style="color: #000000; font-weight: bold"> Filters </p>

      <div v-if="ad.filters.length">
        <FilterCard v-for="(fil, index) in ad.filters" :key="index" v-bind:filter="fil"
                    v-bind:ind="index"/>
      </div>
      <div v-else>
        <div class="card">
          <div class="card-body">
            No filters
          </div>
        </div>
      </div>

    </div>

  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapActions, mapGetters } from 'vuex';
import FilterCard from '@/components/FilterCard.vue';

export default defineComponent({
  name: 'AdsCard',
  components: { FilterCard },
  props: ['ad'],
  data() {
    return {
      content: '',
    };
  },
  mounted() {
    // this.getByID(this.ad.userID);
  },
  methods: {
    // ...mapActions(['getByID']),
  },
  computed: {
    // ...mapGetters(['userAuthor']),
  },
});
</script>
