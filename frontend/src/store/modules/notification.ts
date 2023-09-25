export default {
  actions: {
    async addError(ctx: any, error: string) {
      ctx.commit('addError', error);
    },

    async deleteError(ctx: any, error: string) {
      ctx.commit('deleteError', error);
    },

    async addInfo(ctx: any, info: string) {
      ctx.commit('addInfo', info);
    },

    async deleteInfo(ctx: any, info: string) {
      ctx.commit('deleteInfo', info);
    },
  },

  mutations: {
    addError(state: any, error: string) {
      state.errors.unshift(error);
      console.error(state.errors);
    },

    deleteError(state: any, errorDel: string) {
      state.errors = state.errors.filter((error: string) => (error !== errorDel));
    },

    addInfo(state: any, info: string) {
      state.infos.unshift(info);
      console.info(state.infos);
    },

    deleteInfo(state: any, infoDel: string) {
      state.infos = state.infos.filter((info: string) => (info !== infoDel));
    },
  },

  state: {
    errors: Array<string>(),
    infos: Array<string>(),
  },

  getters: {
    errors(state: any) {
      return state.errors;
    },

    infos(state: any) {
      return state.infos;
    },
  },
};
