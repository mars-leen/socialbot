import {axios} from "../utils/request";
import api from "./index";

export function addCommissionProductApi (parameter) {
    return axios({
        url: api.AddCommissionProduct,
        method: 'post',
        data: parameter
    })
}

