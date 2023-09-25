<template>
  <BaseSection>
    <div class="card-tr mb-2">
      <div class="card-body">
        <p style="color: #ffffff; font-weight: bolder; font-size: 2em;" class="nav-link mb-0">
          Active ads </p>
      </div>
    </div>

    <div v-if="ads.length">
      <AdsCard v-for="(ad, index) in ads" :key="index" v-bind:ad="ad"/>
    </div>
    <div v-else>
      <div class="card">
        <div class="card-body">
          No ads
        </div>
      </div>
    </div>
  </BaseSection>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import BaseSection from '@/components/BaseSection.vue';
import { mapActions, mapGetters } from 'vuex';
import AdsCard from '@/components/AdsCard.vue';

export default defineComponent({
  name: 'AdsView',
  components: { AdsCard, BaseSection },
  mounted() {
    this.userInfo().then(() => {
      if (this.isAdmin) {
        window.location.replace('/'); // TODO
      }

      this.getAllAds();
    });
  },
  methods: {
    ...mapActions(['getAllAds', 'userInfo']),
  },
  computed: {
    ...mapGetters(['isAdmin', 'ads']),
  },
});
</script>
