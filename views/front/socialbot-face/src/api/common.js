import axios from 'axios';
import { Notify } from 'vant';
import {store} from "../store/store";

// config
//export const HOST = 'http://localhost:8080';
export const HOST = 'https://century.wedding';
//export const HOST = 'https://149.28.210.141';

axios.interceptors.request.use(
    config => {
        config.headers = {
            'Authorization': store.getters.GET_TOKEN,
            'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
        };
        return config
    },
    error => {
        return Promise.reject(error)
    }
);


axios.interceptors.response.use(
    response => {  //成功请求到数据
        if (checkCode(response)) {
            return Promise.resolve(response.data);
        }
    },
    error => {  //响应错误处理
        Notify("NET ERROR");
        // eslint-disable-next-line no-console
        console.log(error)
        return Promise.resolve(null);
    }
);

export {axios}



// constant
export const AUTH_CODE = 1002;

// public function
export function checkCode(res) {
    let result = res.data;
    if (!result || !result.msg) {
        Notify("NET ERROR");
        return false
    }

    if (result.code === AUTH_CODE) {
        store.commit("SHOW_MODAL", "login");
        store.dispatch("LOGOUT");
        return  false
    }

    if (result.code !== 0) {
        Notify(result.msg);
        return  false
    }
    return  true
}
