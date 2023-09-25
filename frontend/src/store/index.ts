import Vuex from 'vuex';

import ads from './modules/ads';
import eventTypes from './modules/event_types';
import users from './modules/users';
import user from './modules/user';
import notification from './modules/notification';
import clients from './modules/clients';
import filters from './modules/filters';
import stats from './modules/stats';

// Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    eventTypes,
    ads,
    clients,
    users,
    user,
    notification,
    filters,
    stats,
  },
});
