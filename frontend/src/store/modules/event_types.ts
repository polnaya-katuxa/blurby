// eslint-disable-next-line max-classes-per-file
import * as openapi from '@/openapi';
import API from '@/api';
import Cookies from 'cookies-ts';

export default {
  actions: {
    async getEventTypes(ctx: any) {
      try {
        const resp = await API.getEventTypes();

        ctx.commit('saveEventTypes', resp.data.event_types);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async createEventType(ctx: any, payload: {alias: string, name: string}) {
      try {
        await API.createEventType({
          event_type: {
            id: '',
            name: payload.name,
            alias: payload.alias,
          },
        });

        const resp1 = await API.getEventTypes();

        ctx.commit('saveEventTypes', resp1.data.event_types);
        ctx.dispatch('addInfo', 'Successfully created event type.');
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },
  },

  mutations: {
    saveEventTypes(state: any, eventTypes: Array<openapi.EventType>) {
      state.eventTypes = eventTypes;
    },
  },

  state: {
    eventTypes: Array<openapi.EventType>(),
  },

  getters: {
    ets(state: any) {
      return state.eventTypes;
    },
  },
};
