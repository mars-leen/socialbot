import api from './index'
import {axios} from '../utils/request'

export function addServerConfigApi (parameter) {
    return axios({
        url: api.AddServerConfig,
        method: 'post',
        data: parameter
    })
}


export function updateServerConfigApi (parameter) {
    return axios({
        url: api.UpdateServerConfig,
        method: 'post',
        data: parameter
    })
}


export function deleteServerConfigApi (parameter) {
    return axios({
        url: api.DeleteServerConfig,
        method: 'post',
        data: parameter
    })
}

export function listServerConfigApi (parameter) {
    return axios({
        url: api.ListServerConfig,
        method: 'get',
        params: parameter,
    })
}

