import api from './index'
import {axios} from '../utils/request'

export function listCategoryApi (parameter) {
    return axios({
        url: api.ListCategory,
        method: 'get',
        params: parameter,
    })
}
