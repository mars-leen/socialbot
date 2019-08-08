import api from './index'
import {axios} from '../utils/request'

export function deleteCrawlerItemApi (parameter) {
    return axios({
        url: api.DeleteCrawlerItem,
        method: 'post',
        data: parameter
    })
}

export function listRandCrawlerItemApi (parameter) {
    return axios({
        url: api.ListCrawlerItem,
        method: 'get',
        params: parameter,
    })
}


export function addSocialMediaFromCrawlerApi (parameter) {
    return axios({
        url: api.AddSocialMediaFromCrawler,
        method: 'post',
        data: parameter
    })
}
