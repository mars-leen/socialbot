import axios from 'axios';
import {store} from "../store/store";
import { message } from 'ant-design-vue';
import {router} from '../router'

// config
//export const HOST = 'http://localhost:8080';
export const HOST = 'https://century.wedding';
//export const HOST = 'https://149.28.210.141';

// constant
export const AUTH_CODE = 1002;
export const NET_ERROR = "NET ERROR";


// config
axios.interceptors.request.use(
    config => {
        config.headers = {
            'Authorization': store.getters.getToken,
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
    error => {
        message.error(NET_ERROR, 3);
        // eslint-disable-next-line no-console
        console.log(error);
        return Promise.resolve(null);
    }
);
export {axios}


// public function
export function checkCode(res) {
    let result = res.data;
    if (!result || !result.msg) {
        message.error(NET_ERROR, 3);
        return false
    }

    if (result.code === AUTH_CODE) {
        router.push("/login")
        return  false
    }

    if (result.code !== 0) {
        message.error(result.msg, 3);
        return  false
    }
    return  true
}