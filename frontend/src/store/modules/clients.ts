// eslint-disable-next-line max-classes-per-file
import * as openapi from '@/openapi';
import API from '@/api';
import { Marked } from '@ts-stack/markdown';
import { ClientStats } from '@/openapi';

export default {
  actions: {
    async viewClients(ctx: any) {
      try {
        const resp = await API.getClients();

        ctx.commit('saveClients', resp.data.clients);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async deleteClient(ctx: any, id: string) {
      try {
        await API.deleteClient({
          id,
        });

        ctx.commit('deleteClient', id);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async getStats(ctx: any) {
      try {
        const resp = await API.getClientStats();

        ctx.commit('saveStats', resp.data.clientStats);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },
  },

  mutations: {
    saveClients(state: any, clients: Array<openapi.Client>) {
      state.clients = clients;
    },

    deleteClient(state: any, id: string) {
      state.clients = state.clients.filter((client: openapi.Client) => (client.id !== id));
    },

    saveStats(state: any, stats: openapi.ClientStats) {
      state.stats = stats;
    },
  },

  state: {
    clients: Array<openapi.Client>(),
    stats: {} as openapi.ClientStats,
  },

  getters: {
    clients(state: any) {
      return state.clients;
    },

    stats(state: any) {
      return state.stats;
    },
  },
};
