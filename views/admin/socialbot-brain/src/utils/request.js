import axios from 'axios'
import api from '../api/index'
import store from "../store/";
import {router} from '../router'
import { notification } from 'ant-design-vue';
const service = axios.create({
    baseURL: api.Host, // api base_url
    timeout: 30000 // 请求超时时间
});

const err = (error) => {
    if (error.response) {
        const data = error.response.data
        if (error.response.status === 403) {
            notification.error({
                message: 'Forbidden',
                description: data.message
            })
        }
        if (error.response.status === 401 && !(data.code && data.code===0)) {
            /*notification.error({
                message: 'Unauthorized',
                description: 'Authorization verification failed'
            });*/
            router.push("/dashboard/login")
            /*if (token) {
                store.dispatch('Logout').then(() => {
                    setTimeout(() => {
                        window.location.reload()
                    }, 1500)
                })
            }*/
        }
    }
    return Promise.reject(error)
}

// request interceptor
service.interceptors.request.use(config => {
    const token = store.getters.Token();
    if (token) {
        config.headers['Authorization'] = token // 让每个请求携带自定义 token 请根据实际情况自行修改
    }
    return config
}, err);

// response interceptor
service.interceptors.response.use((response) => {
    if (response.data.code !== 0) {
        notification.error({
            message: "操作失败",
            description: response.data.msg
        });
    }else{
        return response.data
    }
}, err);


export {service as axios}
