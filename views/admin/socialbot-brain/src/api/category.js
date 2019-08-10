import api from './index'
import {axios} from '../utils/request'

export function addCategoryApi (parameter) {
    return axios({
        url: api.AddCategory,
        method: 'post',
        data: parameter,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    })
}


export function updateCategoryApi (parameter) {
    return axios({
        url: api.UpdateCategory,
        method: 'post',
        data: parameter,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    })
}


export function deleteCategoryApi (parameter) {
    return axios({
        url: api.DeleteCategory,
        method: 'post',
        data: parameter
    })
}

export function listCategoryApi (parameter) {
    return axios({
        url: api.ListCategory,
        method: 'get',
        params: parameter,
    })
}

export function listCategoryWithTagsApi (parameter) {
    return axios({
        url: api.ListCategoryWithTags,
        method: 'get',
        params: parameter,
    })
}
