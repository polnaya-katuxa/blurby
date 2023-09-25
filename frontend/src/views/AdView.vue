<template>
  <BaseSection>
    <AdCard/>
  </BaseSection>

</template>

<script lang="ts">
// eslint-disable-next-line max-classes-per-file
import { defineComponent } from 'vue';
import BaseSection from '@/components/BaseSection.vue';
import { mapActions, mapGetters } from 'vuex';
import AdCard from '@/components/AdCard.vue';
import PublishModal from '@/components/EventTypeCard.vue';

export default defineComponent({
  name: 'AdView',
  components: { AdCard, BaseSection },
  data() {
    return {
      perms: false,
      content: '',
    };
  },
  mounted() {
    this.userInfo().then(() => {
      if (this.isAdmin) {
        window.location.replace('/'); // TODO
      }
    });
  },
  created() {
    window.scrollTo(0, 0);
  },
  methods: {
    ...mapActions(['userInfo']),
  },
  computed: {
    ...mapGetters(['isAdmin']),
  },
});
</script>
