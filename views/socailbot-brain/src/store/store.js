import Vue from 'vue'
import Vuex from 'vuex'
import {saveToken, removeUser, saveUser, loadToken, loadUser} from '../assets/js/cache';


Vue.use(Vuex);

const state = {
    user: loadUser(false),
    token: loadToken(false),
};

const getters = {
    getUser: state => {
        if (state.user) {
            return state.user
        }
        state.user = loadUser(null);
        return  state.user
    },

    getToken: state => {
        if (state.token) {
            return state.token
        }
        console.log(22222)
        return loadToken("")
    },
};

const mutations = {
    setUser : function (state, user) {
        state.user = user
    },
    setToken : function (state, token) {
        state.token = token
    },
    loginSaveUser: function(state, user){
        state.user = user;
        state.token = user.token;
        saveToken(user.token);
        saveUser(user);
        console.log(1111111)
    }
};


const store = new Vuex.Store({
    state,
    getters,
    mutations,
});

export {store}



