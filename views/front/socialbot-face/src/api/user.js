import {axios} from "../utils/request";
import api from "./index"

export function loginApi(account, password, type) {
    let parameter = new FormData;
    parameter.append("account", account);
    parameter.append("password", password);
    parameter.append("type", type);
    return axios({
        url: api.Login,
        method: 'post',
        data: parameter
    })
}

export function registerApi(account, password, type) {
    let parameter = new FormData;
    parameter.append("account", account);
    parameter.append("password", password);
    parameter.append("type", type);
    return axios({
        url: api.Register,
        method: 'post',
        data: parameter
    })
}


export function editProfile(nickname, intro, avatar) {
    let parameter = new FormData;
    parameter.append("nickname", nickname);
    parameter.append("intro", intro);
    if (avatar) {
        parameter.append("avatar", avatar);
    }
    return axios({
        url: api.EditProfile,
        method: 'post',
        data: parameter,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    })
}












