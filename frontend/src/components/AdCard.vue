<template>
  <div class="card-tr mb-2">
    <div class="card-body">
      <p style="color: #ffffff; font-weight: bolder; font-size: 2em;" class="nav-link mb-0">
        Create an advertisement </p>
    </div>
  </div>

  <div class="card mb-2">
    <div class="mx-4 my-4">

<!--      <p style="color: #000000; font-weight: bolder; font-size: 2em">-->
<!--        Create an advertisement </p>-->

      <form @submit.prevent="onSubmit">

        <div v-if="preview" class="card px-2 py-2" style="height: 300px">
          <div v-html="content" style="overflow: scroll;"></div>
        </div>

        <div v-else class="form-floating">
              <textarea class="form-control text-pink" placeholder="Ad content in HTML"
                        id="floatingTextarea2" required style="height: 100px"
                        v-model="content"></textarea>
          <!-- eslint-disable-next-line -->
          <label for="floatingTextarea2">Ad content in HTML</label>
        </div>

        <br>

        <div class="form-check form-switch">
          <input class="form-check-input check-pink" type="checkbox" id="flexSwitchCheckDefault"
                 v-model="preview">
          <!-- eslint-disable-next-line -->
          <label class="form-check-label" for="flexSwitchCheckDefault"
          >Preview</label>
        </div>

        <br>

        <p style="color: #000000; font-weight: bold">Schedule preferences </p>

        <div class="form-check form-switch">
          <input class="form-check-input check-pink" type="checkbox" id="flexSwitchCheckDefault"
                 v-model="periodic">
          <!-- eslint-disable-next-line -->
          <label class="form-check-label" for="flexSwitchCheckDefault"
          >Send periodically</label>
        </div>

        <br>

        <div class="input-group">
          <span class="input-group-text text-dark">Span (m:h:d):</span>
          <input type="number" class="form-control" min="0" max="59" step="1" placeholder="Minutes"
                 aria-label="Minutes" required v-model="ssMin">
          <input type="number" class="form-control" min="0" max="23" step="1" placeholder="Hours"
                 aria-label="Hours" required v-model="ssHours">
          <input type="number" class="form-control" min="0" max="365" step="1" placeholder="Days"
                 aria-label="Days" required v-model="ssDays">
        </div>

        <br>

        <p style="color: #000000; font-weight: bold">Filtering preferences </p>

        <div>
          <select required class="form-select" v-model="type"
                  aria-label=".form-select-sm example">
            <option value="by field">Field filter</option>
            <option value="by event">Event filter</option>
          </select>

          <div v-if="type === 'by event'">
            <br>

            <select required class="form-select" v-model="eventType"
                    aria-label=".form-select-sm example">
              <option v-for="et in ets" :key="et.id" :value="et.alias">{{ et.name }}</option>
            </select>

            <br>

            <div class="form-check form-switch">
              <input class="form-check-input check-pink" type="checkbox" id="flexSwitchCheckDefault"
                     v-model="spannedFilter">
              <!-- eslint-disable-next-line -->
              <label class="form-check-label" for="flexSwitchCheckDefault"
              >Spanned filter</label>
            </div>

            <br>

            <div class="input-group" v-if="spannedFilter">
              <span class="input-group-text text-dark">Span (m:h:d):</span>
              <input required type="number" v-model="fsMin" class="form-control" min="0" max="59"
                     step="1" placeholder="Minutes" aria-label="Minutes">
              <input required type="number" v-model="fsHours" class="form-control" min="0" max="23"
                     step="1" placeholder="Hours" aria-label="Hours">
              <input required type="number" v-model="fsDays" class="form-control" min="0" max="365"
                     step="1" placeholder="Days" aria-label="Days">
            </div>

            <br v-if="spannedFilter">

            <div class="input-group">
              <span class="input-group-text text-dark">Rate:</span>
              <input type="number" class="form-control" min="1" max="100" step="1" v-model="rate"
                     placeholder="Rate" required aria-label="Rate">
            </div>

          </div>

          <div v-if="type === 'by field'">
            <br>

            <div class="input-group">
              <select required class="col-sm-3 form-select form-select-sm" v-model="field"
                      aria-label=".form-select-sm example">
                <option value="age">Age filter</option>
                <option value="name">Name filter</option>
                <option value="birth date">Birth date filter</option>
                <option value="gender">Gender filter</option>
              </select>

              <select class="form-select form-select-sm col-sm-3" v-model="cmp"
                      aria-label=".form-select-sm example" required>
                <option v-if="field == 'age' || field == 'birth date'" value="<">&lt;</option>
                <option v-if="field == 'age' || field == 'birth date'" value=">">&gt;</option>
                <option value="=">=</option>
                <option v-if="field == 'age' || field == 'birth date'"
                        value="between">between</option>
              </select>

              <input type="text" class="form-control col-3" placeholder="Value #1" v-model="v1"
                     aria-label="Value #1" required>

              <input type="text" class="form-control col-3" v-if="cmp == 'between'" v-model="v2"
                     placeholder="Value #2" aria-label="Value #2" required>
            </div>

          </div>

          <br>

          <div class="row mx-0">
            <button type="button" class="btn btn-reaction col-4" @click.prevent="onAdd"
                    data-mdb-ripple-color="dark" style="z-index: 1;">
              Add filter
            </button>
          </div>

          <br>

          <div class="row mx-0">
            <p class="px-0" style="color: #000000; font-weight: bold"> Actions </p>

            <button v-if="filters.length" type="button" class="btn btn-reaction col-4"
                    @click.prevent="onTest"
                    data-mdb-ripple-color="dark" style="z-index: 1;">
              Test filters
            </button>
            <div v-if="filters.length" class="col-4 text-center"
                 style="border-color: #0a97bf; border-width: 1px;
                 border-radius: 5px; padding-top: 0.375em; padding-bottom: 0.375em;">
              {{ getTestRes }} clients
            </div>
            <button type="submit" class="btn btn-reaction col-4" data-mdb-ripple-color="dark"
                    style="z-index: 1;">
              Create ad
            </button>
          </div>

          <br>

          <p style="color: #000000; font-weight: bold"> Ready filters </p>

          <div v-if="filters.length">
            <FilterCard v-for="(fil, index) in filters" :key="index" v-bind:filter="fil"
                        v-bind:ind="index"/>
          </div>
          <div v-else>
            <div class="card">
              <div class="card-body">
                No ready filters
              </div>
            </div>
          </div>

        </div>

      </form>
    </div>

  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import FilterCard from '@/components/FilterCard.vue';
import { mapActions, mapGetters } from 'vuex';

export default defineComponent({
  name: 'AdCard',
  components: { FilterCard },
  data() {
    return {
      content: '',
      periodic: false,
      type: 'by field',
      spannedFilter: false,
      eventType: '',
      field: '',
      cmp: '',
      v1: '',
      v2: '',
      fsMin: 0,
      fsHours: 0,
      fsDays: 0,
      rate: 0,
      ssMin: 0,
      ssHours: 0,
      ssDays: 0,
      preview: false,
    };
  },
  mounted() {
    this.userInfo().then(() => {
      if (this.isAdmin) {
        window.location.replace('/'); // TODO
      }
    });

    this.getEventTypes();
  },
  methods: {
    ...mapActions(['userInfo', 'getEventTypes', 'testFilters', 'addFilter', 'createAd',
      'deleteFilters', 'zeroTest']),
    onSubmit() {
      this.createAd({
        content: this.content,
        periodic: this.periodic,
        min: this.ssMin,
        hours: this.ssHours,
        days: this.ssDays,
        userID: this.user.id,
        filters: this.filters,
      });

      this.content = '';
      this.periodic = false;
      this.ssMin = 0;
      this.ssHours = 0;
      this.ssDays = 0;
      this.deleteFilters();
      this.type = '';
      this.eventType = '';
      this.rate = 0;
      this.fsMin = 0;
      this.fsHours = 0;
      this.fsDays = 0;
      this.field = '';
      this.cmp = '';
      this.v1 = '';
      this.v2 = '';
      this.spannedFilter = false;
      this.zeroTest();
    },
    onTest() {
      this.testFilters(this.filters);
    },
    onAdd() {
      this.addFilter({
        type: this.type,
        eventType: this.eventType,
        rate: this.rate,
        min: this.fsMin,
        hours: this.fsHours,
        days: this.fsDays,
        field: this.field,
        cmp: this.cmp,
        v1: this.v1,
        v2: this.v2,
        spanned: this.spannedFilter,
      });
    },
  },
  computed: {
    ...mapGetters(['isAdmin', 'user', 'ets', 'getTestRes', 'filters']),
  },
});
</script>
