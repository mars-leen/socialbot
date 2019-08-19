import api from './index'
import {axios} from '../utils/request'

export function updateConfigApi (parameter) {
    return axios({
        url: api.UpdateConfig,
        method: 'post',
        data: parameter
    })
}

export function BaseConfigApi (parameter) {
    return axios({
        url: api.BaseConfig,
        method: 'get',
        params: parameter,
    })
}

