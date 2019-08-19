import api from './index'
import {axios} from '../utils/request'

export function addReverseConfigApi (parameter) {
    return axios({
        url: api.AddReverseConfig,
        method: 'post',
        data: parameter
    })
}

export function updateReverseConfigApi (parameter) {
    return axios({
        url: api.UpdateReverseConfig,
        method: 'post',
        data: parameter
    })
}

export function deleteReverseConfigApi (parameter) {
    return axios({
        url: api.DeleteReverseConfig,
        method: 'post',
        data: parameter
    })
}

export function listReverseConfigApi (parameter) {
    return axios({
        url: api.ListReverseConfig,
        method: 'get',
        params: parameter,
    })
}

