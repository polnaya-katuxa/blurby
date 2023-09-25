// eslint-disable-next-line max-classes-per-file
import * as openapi from '@/openapi';
import API from '@/api';
import Cookies from 'cookies-ts';

export default {
  actions: {
    async userInfo(ctx: any) {
      try {
        const resp = await API.userInfo();

        ctx.commit('saveUser', resp.data.user);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async login(ctx: any, payload: {login: string, password: string}) {
      try {
        const resp = await API.login({
          login: payload.login,

          password: payload.password,
        });

        const cookies = new Cookies();
        cookies.set('user-token', resp.data.token);
        window.location.replace('/'); // TODO
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    // async getByID(ctx: any, id: string) {
    //   try {
    //     const resp = await API.getUserByID({ id });
    //
    //     ctx.commit('saveUserAuthor', resp.data.user);
    //   } catch (err: any) {
    //     ctx.dispatch('addError', err.response.data.message);
    //   }
    // },

    async register(ctx: any, payload: {login: string, password: string}) {
      try {
        const resp = await API.register({
          login: payload.login,
          password: payload.password,
        });

        const cookies = new Cookies();
        cookies.set('user-token', resp.data.token);
        window.location.replace('/'); // TODO
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },
  },

  mutations: {
    saveUser(state: any, user: openapi.User) {
      state.user = user;
    },

    // saveUserAuthor(state: any, user: openapi.User) {
    //   state.userAuthor = user;
    // },
  },

  state: {
    user: {} as openapi.User,
  },

  getters: {
    user(state: any) {
      return state.user;
    },

    isAdmin(state: any) {
      return state.user.isAdmin;
    },

    curLogin(state: any) {
      return state.user.login;
    },
  },
};
