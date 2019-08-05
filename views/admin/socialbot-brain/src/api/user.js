import api from './index'
import {axios} from '../utils/request'

export function loginApi (parameter) {
    return axios({
        url: api.Login,
        method: 'post',
        data: parameter
    })
}
