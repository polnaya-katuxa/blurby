// eslint-disable-next-line max-classes-per-file
import * as openapi from '@/openapi';
import API from '@/api';

export default {
  actions: {
    async testFilters(ctx: any, filters: Array<openapi.Filter>) {
      try {
        const resp = await API.filter({
          filters,
        });

        ctx.commit('saveTestRes', resp.data.count);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async addFilter(ctx: any, payload: { type: string, eventType: string, rate: number,
    min: number, hours: number, days: number, field: string, cmp: string, v1: string, v2: string,
      spanned: boolean}) {
      let span = '';
      if (payload.spanned) {
        span = `${payload.hours + payload.days * 24}h${payload.min}m`;
      }

      const filter = new class implements openapi.Filter {
        type = payload.type;

        filter = {
          field: payload.field,
          cmp: payload.cmp,
          value1: payload.v1,
          value2: payload.v2,
          alias: payload.eventType,
          rate: payload.rate,
          span,
        };
      }();

      ctx.commit('saveFilter', filter);
    },

    async deleteFilter(ctx: any, ind: number) {
      ctx.commit('deleteFilter', ind);
    },

    async deleteFilters(ctx: any) {
      ctx.commit('deleteFilters');
    },

    async zeroTest(ctx: any) {
      ctx.commit('saveTestRes', 0);
    },
  },

  mutations: {
    saveFilter(state: any, fil: openapi.Filter) {
      state.filters.unshift(fil);
    },

    deleteFilter(state: any, ind: number) {
      state.filters.splice(ind, 1);
    },

    deleteFilters(state: any) {
      state.filters = Array<openapi.Filter>();
    },

    saveTestRes(state: any, res: number) {
      state.testRes = res;
    },
  },

  state: {
    testRes: 0,
    filters: Array<openapi.Filter>(),
  },

  getters: {
    filters(state: any) {
      return state.filters;
    },

    getTestRes(state: any) {
      return state.testRes;
    },
  },
};
