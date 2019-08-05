import api from './index'
import {axios} from '../utils/request'

export function addTagApi (parameter) {
    return axios({
        url: api.AddTag,
        method: 'post',
        data: parameter
    })
}


export function updateTagApi (parameter) {
    return axios({
        url: api.UpdateTag,
        method: 'post',
        data: parameter
    })
}


export function deleteTagApi (parameter) {
    return axios({
        url: api.DeleteTag,
        method: 'post',
        data: parameter
    })
}

export function listTagApi (parameter) {
    return axios({
        url: api.ListTag,
        method: 'get',
        params: parameter,
        headers: {
            'Content-Type': 'application/json;charset=UTF-8'
        }
    })
}
