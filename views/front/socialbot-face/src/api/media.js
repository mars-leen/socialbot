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
