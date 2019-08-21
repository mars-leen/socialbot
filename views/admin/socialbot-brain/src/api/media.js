import {axios} from "../utils/request";
import api from "./index";

export function mediaDetailApi (parameter) {
    return axios({
        url: api.MediaDetail,
        method: 'get',
        params: parameter,
    })
}

export function editMediaApi (parameter) {
    return axios({
        url: api.EditMedia,
        method: 'post',
        data: parameter
    })
}


export function listMediasApi (parameter) {
    return axios({
        url: api.ListMedias,
        method: 'get',
        params: parameter,
    })
}

export function addCommissionProductApi (parameter) {
    return axios({
        url: api.AddCommissionProduct,
        method: 'post',
        data: parameter
    })
}


