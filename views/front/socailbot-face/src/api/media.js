import {axios} from "../utils/request";
import api from "./index";

export function listMediaByCategoryApi (parameter) {
    return axios({
        url: api.ListMediaByCategory,
        method: 'get',
        params: parameter,
    })
}

