import Vue from "vue"
import Constant from "./constant"
const getters = {
  USER: state => {
    if (state.user) {
      return state.user
    }
    state.user = Vue.ls.get(Constant.USER_INFO_KEY)
    return  state.user
  },
  TOKEN: state => {
    if (state.token) {
      return state.token
    }
    return Vue.ls.get(Constant.USER_TOKEN_KEY)
  },
};
export default getters
