import {HOST, axios} from "./common";


const GET_ALL_TAG = `${HOST}/v1/api/tag/getAll`;
const GET_ALL_CATEGORY = `${HOST}/v1/api/category/getAll`;

export function getAllTagApi() {
    return axios.get(GET_ALL_TAG, {
        params: {}
    }).then((res) => {
        return Promise.resolve(res);
    });
}

export function getAllCategoryApi() {
    return axios.get(GET_ALL_CATEGORY, {
        params: {}
    }).then((res) => {
        return Promise.resolve(res);
    });
}