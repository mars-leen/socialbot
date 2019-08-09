import api from './index'
import {axios} from '../utils/request'

export function loginApi (parameter) {
    return axios({
        url: api.Login,
        method: 'post',
        data: parameter
    })
}

export function registerApi (parameter) {
    return axios({
        url: api.Register,
        method: 'post',
        data: parameter
    })
}
