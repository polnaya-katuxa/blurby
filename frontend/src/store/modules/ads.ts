// eslint-disable-next-line max-classes-per-file
import * as openapi from '@/openapi';
import API from '@/api';

export default {
  actions: {
    async createAd(ctx: any, payload: { content: string, periodic: boolean,
      min: number, hours: number, days: number, userID: string, filters: Array<openapi.Filter>}) {
      const span = `${payload.hours + payload.days * 24}h${payload.min}m`;

      try {
        await API.createAd({
          ad: {
            content: payload.content,
            filters: payload.filters,
            userID: payload.userID,
            schedule: {
              periodic: payload.periodic,
              finished: false,
              nextTime: '',
              span,
            },
            create_time: '',
          },
        });

        const resp = await API.getAds();

        ctx.commit('saveAds', resp.data.ads);
        ctx.dispatch('addInfo', 'Successfully created ad.');
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    async getAllAds(ctx: any) {
      try {
        const resp = await API.getAds();

        ctx.commit('saveAds', resp.data.ads);
      } catch (err: any) {
        ctx.dispatch('addError', err.response.data.message);
      }
    },

    // async uncomment(ctx: any, payload: { postID: string, commID: string }) {
    //   try {
    //     console.error(payload);
    //     await API.uncomment(
    //       payload.postID,
    //       new class implements openapi.UncommentRequest {
    //         commID = payload.commID;
    //       }(),
    //     );
    //
    //     ctx.commit('deleteComment', payload.commID);
    //   } catch (err: any) {
    //     ctx.dispatch('uncommentError', err.response.data.message);
    //   }
    // },
  },

  mutations: {
    saveAds(state: any, ads: Array<openapi.Ad>) {
      state.ads = ads;
    },

    // savePost(state: any, post: openapi.Post) {
    //   state.post = post;
    // },
    //
    // deleteComment(state: any, commID: string) {
    //   state.comments = state.comments.filter((comm: openapi.Comment) => (comm.id !== commID));
    // },
  },

  state: {
    // post: {} as openapi.Post,
    ads: Array<openapi.Ad>(),
  },

  getters: {
    ads(state: any) {
      return state.ads;
    },
  },
};
