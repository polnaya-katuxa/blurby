import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import HomeView from '../views/HomeView.vue';
import LoginView from '../views/LoginView.vue';
import RegisterView from '../views/RegisterView.vue';
import UsersView from '../views/UsersView.vue';
import AdView from '../views/AdView.vue';
import AdsView from '../views/AdsView.vue';
import ClientsView from '../views/ClientsView.vue';
import EventTypesView from '../views/EventTypesView.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView,
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterView,
  },
  {
    path: '/ad',
    name: 'ad',
    component: AdView,
  },
  {
    path: '/ads',
    name: 'ads',
    component: AdsView,
  },
  {
    path: '/et',
    name: 'event_types',
    component: EventTypesView,
  },
  {
    path: '/users',
    name: 'users',
    component: UsersView,
  },
  {
    path: '/clients',
    name: 'clients',
    component: ClientsView,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
  scrollBehavior(to, from, savedPosition) {
    return {
      top: 0,
      behavior: 'smooth',
    };
  },
});

export default router;
