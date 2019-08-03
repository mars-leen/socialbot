import {HOST, axios} from "./common";

const LOGIN = `${HOST}/v1/adminApi/login`;

export function loginApi(account, password, type) {
    let form = new FormData;
    form.append("account", account);
    form.append("password", password);
    form.append("type", type);
    return axios.post(LOGIN, form).then((res) => {
        return Promise.resolve(res);
    });
}