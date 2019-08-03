import {axios, HOST} from "./common";

const GET_RAND_LIST = `${HOST}/v1/adminApi/content/getRandList`;
const DELETE = `${HOST}/v1/adminApi/content/delete`;
const PUBLISH = `${HOST}/v1/adminApi/content/publish`;


export function getRandListApi() {
    return axios.get(GET_RAND_LIST, {
        params: {}
    }).then((res) => {
        return Promise.resolve(res);
    });
}



export function deleteApi(id) {
    let form = new FormData;
    form.append("id", id);
    return axios.post(DELETE, form).then((res) => {
        return Promise.resolve(res);
    });
}

export function publishApi(contentId, title,desc, cover, medias, tags, categoryId, type, recommend) {
    let re = (recommend ? 1 : 0);

    let form = new FormData;
    form.append("contentId", contentId);
    form.append("title", title);
    form.append("desc", desc);
    form.append("cover", cover);
    form.append("medias", medias);
    form.append("tags",tags);
    form.append("categoryId",categoryId);
    form.append("type",type);
    form.append("recommend", re);
    console.log("--->",form)
    return axios.post(PUBLISH, form).then((res) => {
        return Promise.resolve(res);
    });
}
