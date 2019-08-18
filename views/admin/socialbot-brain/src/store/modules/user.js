import Vue from 'vue'
import Constant from '../constant'
const user = {
  state: {
    token: '',
    info:{}
  },
  mutations: {
    SET_TOKEN: (state, token) => {
      state.token = token
    },
    SET_INFO: (state, info) => {
      state.info = info
    },
  },
  actions: {
    SaveUserInfo({ commit }, userInfo) {
      commit('SET_TOKEN', userInfo.token);
      commit('SET_INFO', userInfo);
      Vue.ls.set(Constant.USER_TOKEN_KEY, userInfo.token, 7 * 24 * 60 * 60 * 1000);
      Vue.ls.set(Constant.USER_INFO_KEY, userInfo, 7 * 24 * 60 * 60 * 1000);
    },
  }
};

export default user
