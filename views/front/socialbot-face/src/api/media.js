import {axios} from "../utils/request";
import api from "./index"


export function homeRecommendApi (parameter) {
    return axios({
        url: api.HomeRecommend,
        method: 'get',
        params: parameter,
    })
}

export function listMediaByCategoryApi (parameter) {
    return axios({
        url: api.ListMediaByCategory,
        method: 'get',
        params: parameter,
    })
}

export function mediaDetailApi (parameter) {
    return axios({
        url: api.MediaDetail,
        method: 'get',
        params: parameter,
    })
}


export function likeMediaApi(uri, isLike) {
    if(isLike) {
        isLike = 1;
    }else{
        isLike = 0;
    }
    let parameter = new FormData;
    parameter.append("uri", uri);
    parameter.append("isLike", isLike);
    return axios({
        url: api.LikeMedia,
        method: 'post',
        data: parameter,
    })
}









import {HOST} from "./common";


const GET_TOP = `${HOST}/v1/api/media/getTop`;
const GET_DETAIL = `${HOST}/v1/api/media/getDetail`;
const LIKE = `${HOST}/v1/api/media/like`;

const GET_LIST_BY_CATEGORY = `${HOST}/v1/api/media/getListByCategory`;
const GET_LIST_BY_TAGS = `${HOST}/v1/api/media/getListByTags`;



export function getTopApi() {
    return axios.get(GET_TOP, {
        params: {}
    }).then((res) => {
        return Promise.resolve(res);
    });
}

export function getDetailApi(uri) {
    return axios.get(GET_DETAIL, {
        params: {uri: uri}
    }).then((res) => {
        return Promise.resolve(res);
    });
}

export function getListByCategoryApi(lastId, category, sort) {
    return axios.get(GET_LIST_BY_CATEGORY, {
        params: {lastId:lastId, category, sort}
    }).then((res) => {
       return Promise.resolve(res);
    });
}

export function getListByTagsApi(lastId, tags, sort) {
    return axios.get(GET_LIST_BY_TAGS, {
        params: {lastId, tags, sort}
    }).then((res) => {
        return Promise.resolve(res);
    });
}

export function likeApi(uri, isLike) {
    if(isLike) {
        isLike = 1;
    }else{
        isLike = 0;
    }
    let form = new FormData
    form.append("uri", uri);
    form.append("isLike", isLike);

    return axios.post(LIKE, form).then((res) => {
        return Promise.resolve(res);
    });
}
