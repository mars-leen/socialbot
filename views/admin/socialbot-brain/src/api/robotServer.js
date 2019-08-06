import api from './index'
import {axios} from '../utils/request'

export function addRobotServerApi (parameter) {
    return axios({
        url: api.AddRobotServer,
        method: 'post',
        data: parameter
    })
}

export function deleteRobotServerApi (parameter) {
    return axios({
        url: api.DeleteRobotServer,
        method: 'post',
        data: parameter
    })
}

export function listRobotServerApi (parameter) {
    return axios({
        url: api.ListRobotServer,
        method: 'get',
        params: parameter,
    })
}
