import {HOST, axios} from "./common";


const GET_LEVEL_TAG = `${HOST}/v1/api/tag/getLevel`;
const GET_ALL_CATEGORY = `${HOST}/v1/api/category/getAll`;

export function getLevelTagApi() {
    return axios.get(GET_LEVEL_TAG, {
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