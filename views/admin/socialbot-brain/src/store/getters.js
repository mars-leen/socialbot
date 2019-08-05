import Vue from "vue"
import Constant from "./constant"
const getters = {
  User: state => {
    if (state.user) {
      return state.user
    }
    state.user = Vue.ls.get(Constant.USER_INFO_KEY)
    return  state.user
  },
  Token: state => {
    if (state.token) {
      return state.token
    }
    console.log(222222);
    return Vue.ls.get(Constant.USER_TOKEN_KEY)
  },
};
export default getters
