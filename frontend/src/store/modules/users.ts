// eslint-disable-next-line max-classes-per-file
import * as openapi from '@/openapi';
import API from '@/api';
import { Marked } from '@ts-stack/markdown';

export default {
  actions: {
    async viewUsers(ctx: any) {
      try {
        const resp = await API.getUsers();

        ctx.commit('saveUsers', resp.data.users);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async deleteUser(ctx: any, login: string) {
      try {
        await API.deleteUser({
          login,
        });

        ctx.commit('deleteUser', login);
        ctx.dispatch('addInfo', 'Successfully deleted user.');
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async grantAdminUser(ctx: any, login: string) {
      try {
        await API.grantUserAdmin({
          login,
        });

        ctx.commit('updateUser', login);
        ctx.dispatch('addInfo', 'Successfully granted admin.');
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },
  },

  mutations: {
    saveUsers(state: any, users: Array<openapi.User>) {
      state.users = users;
    },

    deleteUser(state: any, login: string) {
      state.users = state.users.filter((user: openapi.User) => (user.login !== login));
    },

    updateUser(state: any, login: string) {
      state.users = state.users.map((user: openapi.User) => {
        if (user.login === login) {
          // eslint-disable-next-line no-param-reassign
          user.isAdmin = true;
        }
        return user;
      });
    },
  },

  state: {
    users: Array<openapi.User>(),
  },

  getters: {
    users(state: any) {
      return state.users;
    },
  },
};
