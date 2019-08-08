import {axios} from "../utils/request";
import api from "./index";

export function addGalleryTagApi (parameter) {
    return axios({
        url: api.AddGalleryTagApi,
        method: 'post',
        data: parameter
    })
}

export function listGalleryApi (parameter) {
    return axios({
        url: api.ListGalleryApi,
        method: 'get',
        params: parameter,
    })
}
