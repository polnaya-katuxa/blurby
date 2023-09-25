<template>
  <BaseSection>
    <div class="card-tr mb-2">
      <div class="card-body">
        <p style="color: #ffffff; font-weight: bolder; font-size: 2em;" class="nav-link mb-0">
          Create an event type </p>
      </div>
    </div>

    <div class="card">
      <div class="mx-4 my-4">

<!--        <p style="color: #000000; font-weight: bolder; font-size: 2em">-->
<!--          Create an event type </p>-->

        <form @submit.prevent="onSubmit">

          <input type="text" class="form-control col-3" placeholder="Alias" v-model="alias"
                 style="border-color: #0a97bf;" aria-label="Alias">

          <br>

          <div class="form-floating">
                  <textarea class="form-control text-pink" placeholder="Description"
                            id="floatingTextarea2" style="height: 100px" v-model="name"></textarea>
            <!-- eslint-disable-next-line -->
            <label for="floatingTextarea2">Description</label>
          </div>

          <br>

          <button type="submit" class="btn btn-reaction btn-block">Create</button>
        </form>

        <br>

        <p style="color: #000000; font-weight: bold"> Event types </p>

        <div v-if="ets.length">
          <EventTypeCard v-for="et in ets" :key="et.id" v-bind:et="et"/>
        </div>
        <div v-else>
          <div class="card">
            <div class="card-body">
              No event types
            </div>
          </div>
        </div>

      </div>
    </div>

  </BaseSection>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import * as openapi from '@/openapi';
import BaseSection from '@/components/BaseSection.vue';
import { mapActions, mapGetters } from 'vuex';
import EventTypeCard from '@/components/EventTypeCard.vue';

export default defineComponent({
  name: 'EventTypesView',
  components: { EventTypeCard, BaseSection },
  data() {
    return {
      alias: '',
      name: '',
    };
  },
  created() {
    this.userInfo().then(() => {
      if (this.isAdmin) {
        window.location.replace('/'); // TODO
      }
    });

    this.getEventTypes();
  },
  methods: {
    ...mapActions(['createEventType', 'getEventTypes', 'userInfo']),
    onSubmit() {
      this.createEventType({
        alias: this.alias,
        name: this.name,
      });

      this.alias = '';
      this.name = '';
    },
  },
  computed: {
    ...mapGetters(['isAdmin', 'ets']),
  },
});
</script>
