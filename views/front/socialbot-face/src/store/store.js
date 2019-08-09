import Vue from 'vue'
import Vuex from 'vuex'
import Constant from "./consant"
import {saveUserTag, loadUserTag, } from '../assets/js/cache';


Vue.use(Vuex);

const state = {
    user: false,
    token: "",
    clickStatus: true,
    modal: {
        show: false,
        loginShow: false,
        registerShow: false,
        editShow: false,
    },
    selectedTags: [],
};

const getters = {
    GET_USER: state => {
        if (state.user) {
            return state.user
        }
        return  state.user = Vue.ls.get(Constant.USER_INFO_KEY);
    },
    GET_TOKEN: state => {
        if (state.token) {
            return state.token
        }
        return state.token = Vue.ls.get(Constant.USER_TOKEN_KEY);
    },
    getClickStatus: state => {
        return state.clickStatus
    },
    GET_MODAL: state => {
        return state.modal
    },
    getSelectedTags: state => {
        if (state.selectedTags.length > 0) {
            return state.selectedTags
        }
        state.selectedTags = loadUserTag([]);
        return  state.selectedTags
    }
};

const mutations = {
    updateClickStatus: function(state, status){
        state.clickStatus = status;
    },
    updateSelectedTags: function (state, tagId) {
        const index = state.selectedTags.findIndex((t) => {
            return t === tagId;
        });
        if (index >= 0){
            state.selectedTags.splice(index, 1);
        }else{
            state.selectedTags.push(tagId);
        }
        state.selectedTags = saveUserTag(state.selectedTags);
    },
    resetSelectedTags: function (state, tagId) {
        state.selectedTags = saveUserTag([tagId]);
    },

    SHOW_MODAL : function (state, cmd) {
        if (cmd === "close") {
            state.modal.show = false;
            state.modal.loginShow = false;
            state.modal.registerShow = false;
            state.modal.editShow = false;
        }
        if (cmd === "login") {
            state.modal.show = true;
            state.modal.loginShow = true;
            state.modal.registerShow = false;
            state.modal.editShow = false;
        }
        if (cmd === "register") {
            state.modal.show = true;
            state.modal.loginShow = false;
            state.modal.registerShow = true;
            state.modal.editShow = false;
        }

        if (cmd === "profile") {
            state.modal.show = true;
            state.modal.loginShow = false;
            state.modal.registerShow = false;
            state.modal.editShow = true;
        }
    },
    SET_TOKEN: (state, token) => {
        state.token = token
    },
    SET_USER: (state, user) => {
        state.user = user
    },
};

const actions = {
    LOGIN({ commit }, userInfo) {
        Vue.ls.set(Constant.USER_TOKEN_KEY, userInfo.token, 7 * 24 * 60 * 60 * 1000);
        Vue.ls.set(Constant.USER_INFO_KEY, userInfo, 7 * 24 * 60 * 60 * 1000);
        commit('SET_TOKEN', userInfo.token);
        commit('SET_USER', userInfo);
    },
    LOGOUT({ commit }) {
        Vue.ls.remove(Constant.USER_TOKEN_KEY);
        Vue.ls.remove(Constant.USER_INFO_KEY);
        commit('SET_TOKEN', '');
        commit('SET_USER', null);
    },
    SAVE_PROFILE({ commit }, userInfo){
        Vue.ls.set(Constant.USER_INFO_KEY, userInfo, 7 * 24 * 60 * 60 * 1000);
        commit('SET_USER', userInfo);
    }
};

const store = new Vuex.Store({
    state,
    getters,
    mutations,
    actions
});

export {store}



