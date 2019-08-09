import Vue from 'vue'
import Constant from '../constant'
const user = {
  state: {
    token: '',
    info:{},
    modal: {
      modalShow: false,
      loginShow: false,
      registerShow: false,
      editShow: false,
    },
  },
  mutations: {
    SET_TOKEN: (state, token) => {
      state.token = token
    },
    SET_INFO: (state, info) => {
      state.info = info
    },
    SHOW_MODEL : function (state, cmd) {
      if (cmd === "close") {
        state.model.modalShow = false;
        state.model.loginShow = false;
        state.model.registerShow = false;
        state.model.editShow = false;
      }
      if (cmd === "login") {
        state.model.modalShow = true;
        state.model.loginShow = true;
        state.model.registerShow = false;
        state.model.editShow = false;
      }
      if (cmd === "register") {
        state.model.modalShow = true;
        state.model.loginShow = false;
        state.model.registerShow = true;
        state.model.editShow = false;
      }

      if (cmd === "profile") {
        state.model.modalShow = true;
        state.model.loginShow = false;
        state.model.registerShow = false;
        state.model.editShow = true;
      }
    },
  },
  actions: {
    LOGIN({ commit }, userInfo) {
      Vue.ls.set(Constant.USER_TOKEN_KEY, userInfo.token, 7 * 24 * 60 * 60 * 1000);
      Vue.ls.set(Constant.USER_INFO_KEY, userInfo, 7 * 24 * 60 * 60 * 1000);
      commit('SET_TOKEN', userInfo.token);
      commit('SET_INFO', userInfo);
    },
    LOGOUT({ commit }) {
      Vue.ls.remove(Constant.USER_TOKEN_KEY);
      Vue.ls.remove(Constant.USER_INFO_KEY);
      commit('SET_TOKEN', '');
      commit('SET_INFO', {});
    },
  }
};

export default user
