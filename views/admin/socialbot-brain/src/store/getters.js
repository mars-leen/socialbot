import Vue from "vue"
import Constant from "./constant"
const getters = {
  User: state => ()=> {
    if (state.info) {
      return state.info
    }
    state.info = Vue.ls.get(Constant.USER_INFO_KEY);
    return  state.info
  },
  Token: state =>()=> {
    if (state.token) {
      return state.token
    }
    state.token = Vue.ls.get(Constant.USER_TOKEN_KEY);
    return state.token
  },
};
export default getters
