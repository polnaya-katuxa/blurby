<!--<template>-->
<!--  <div class="card mb-4">-->
<!--    <div class="card-body">-->
<!--      <div class="d-flex mb-3">-->
<!--        <div class="flex-shrink-0">-->
<!--          <img v-bind:src="post.author.picture"-->
<!--               style="min-width: 50px; min-height: 50px; max-width: 50px;-->
<!--                         max-height: 50px"-->
<!--               class="align-self-start mr-3 rounded-circle"-->
<!--               alt=""/>-->
<!--        </div>-->
<!--        <div class="flex-grow-1 ms-3" style="padding-top: 0.125rem">-->
<!--          <router-link style="color: black; text-decoration: none"-->
<!--                       v-bind:to="'/profile/' + post.author.login">{{ post.author.login }}-->
<!--          </router-link>-->
<!--          <div class="text-muted small">{{ new Date(post.pubTime).toLocaleString() }}-->
<!--          </div>-->
<!--        </div>-->

<!--        <div class="dropdown" v-if="isOwnProfile">-->
<!--          <button class="btn btn-sm btn-reaction dropdown-toggle" type="button"-->
<!--                  id="dropdownMenuButton1" data-bs-toggle="dropdown"-->
<!--                  aria-expanded="false">-->
<!--            Actions-->
<!--          </button>-->
<!--          <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton1">-->
<!--            <li>-->
<!--              <a class="dropdown-item small" href="#" v-if="post.perms"-->
<!--                 @click.prevent="changePermsPost(post.id)">Make free</a>-->
<!--              <a class="dropdown-item small" href="#" v-else-->
<!--                 @click.prevent="changePermsPost(post.id)">Make paid</a>-->
<!--            </li>-->
<!--            <li><a class="dropdown-item small" href="#"-->
<!--                   @click.prevent="deletePost(post.id)">Delete</a></li>-->
<!--          </ul>-->
<!--        </div>-->
<!--      </div>-->

<!--      <div class="styled-scrollbars" style="overflow: auto">-->
<!--        <div v-html="post.content"></div>-->
<!--      </div>-->
<!--    </div>-->

<!--    <ReactionsBlock :post="post"/>-->
<!--  </div>-->
<!--</template>-->

<!--<script lang="ts">-->
<!--import { defineComponent } from 'vue';-->
<!--import ReactionsBlock from '@/components/ReactionsBlock.vue';-->
<!--import { mapActions, mapGetters } from 'vuex';-->

<!--export default defineComponent({-->
<!--  name: 'StatCard',-->
<!--  props: ['stat'],-->
<!--  methods: {-->
<!--    ...mapActions(['changePermsPost', 'deletePost']),-->
<!--  },-->
<!--  computed: {-->
<!--    ...mapGetters(['curLogin']),-->
<!--    isOwnProfile() {-->
<!--      return (window.location.pathname === `/profile/${this.curLogin}`);-->
<!--    },-->
<!--    location() {-->
<!--      return window.location.pathname;-->
<!--    },-->
<!--  },-->
<!--});-->
<!--</script>-->
