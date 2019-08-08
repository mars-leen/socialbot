import api from './index'
import {axios} from '../utils/request'

export function addCopywriterApi (parameter) {
    return axios({
        url: api.AddCopywriter,
        method: 'post',
        data: parameter
    })
}


export function updateCopywriterApi (parameter) {
    return axios({
        url: api.UpdateCopywriter,
        method: 'post',
        data: parameter
    })
}


export function deleteCopywriterApi (parameter) {
    return axios({
        url: api.DeleteCopywriter,
        method: 'post',
        data: parameter
    })
}

export function listCopywriterApi (parameter) {
    return axios({
        url: api.ListCopywriter,
        method: 'get',
        params: parameter,
    })
}

export function searchCopywriterApi (parameter) {
    return axios({
        url: api.ListCopywriter,
        method: 'get',
        params: parameter,
    })
}
