import axios from 'axios'
import api from '../api/index'
import {store} from "../store/store";
import { Notify } from 'vant';


const service = axios.create({
    baseURL: api.Host, // api base_url
    timeout: 30000 // 请求超时时间
});

const err = (error) => {
    if (error.response) {
        const data = error.response.data;
        if (error.response.status === 404) {
            Notify("API Not Found");
        }
        if (error.response.status === 403) {
            Notify("Forbidden");
        }
        if (error.response.status === 401 && !(data.code && data.code===0)) {
            store.commit("SHOW_MODAL", "login");
            store.dispatch("LOGOUT");
        }
    }
    return Promise.reject(error)
};

// request interceptor
service.interceptors.request.use(config => {
    const token = store.getters.GET_TOKEN;
    if (token) {
        config.headers['Authorization'] = token // 让每个请求携带自定义 token 请根据实际情况自行修改
    }
    return config
}, err);

// response interceptor
service.interceptors.response.use((response) => {
    if (response.data.code !== 0) {
        Notify(response.data.msg);
    }else{
        return response.data
    }
}, err);

export {service as axios}
