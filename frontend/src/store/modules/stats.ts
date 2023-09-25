// eslint-disable-next-line max-classes-per-file
import * as openapi from '@/openapi';
import API from '@/api';
import { Marked } from '@ts-stack/markdown';
import { ClientStats } from '@/openapi';

export default {
  actions: {
    async getStats(ctx: any) {
      try {
        const resp = await API.getClientStats();

        ctx.commit('saveClientStats', resp.data.clientStats);
        ctx.commit('saveAdStats', resp.data.adStats);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },
  },

  mutations: {
    saveClientStats(state: any, stats: openapi.ClientStats) {
      state.clientStats = stats;
    },

    saveAdStats(state: any, stats: Array<openapi.AdStat>) {
      state.adStats = stats;
    },
  },

  state: {
    clientStats: {} as openapi.ClientStats,
    adStats: Array<openapi.AdStat>(),
  },

  getters: {
    clientStats(state: any) {
      return state.clientStats;
    },

    adStats(state: any) {
      return state.adStats;
    },

    chartAdNums(state: any) {
      return state.adStats.map((row: openapi.AdStat) => row.num);
    },

    chartAdDates(state: any) {
      return state.adStats.map((row: openapi.AdStat) => row.date);
    },
  },
};
