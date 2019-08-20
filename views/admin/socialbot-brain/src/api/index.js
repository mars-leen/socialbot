
const api = {
    UploadSingle: '/v1/adminApi/upload/single',
    Login: '/v1/adminApi/login',
    Logout: '/v1/adminApi/logout',

    UpdateConfig: '/v1/adminApi/config/update',
    BaseConfig: '/v1/adminApi/config/base',

    AddServerConfig: '/v1/adminApi/config/addServer',
    UpdateServerConfig: '/v1/adminApi/config/updateServer',
    ListServerConfig: '/v1/adminApi/config/listServer',
    DeleteServerConfig: '/v1/adminApi/config/deleteServer',

    AddReverseConfig: '/v1/adminApi/config/addReverseHost',
    UpdateReverseConfig: '/v1/adminApi/config/updateReverseHost',
    ListReverseConfig: '/v1/adminApi/config/listReverseHost',
    DeleteReverseConfig: '/v1/adminApi/config/deleteReverseHost',


    AddCategory: '/v1/adminApi/category/add',
    UpdateCategory: '/v1/adminApi/category/update',
    ListCategory: '/v1/adminApi/category/list',
    DeleteCategory: '/v1/adminApi/category/delete',
    ListCategoryWithTags: '/v1/adminApi/category/listWithTags',

    AddCopywriter: '/v1/adminApi/copywriter/add',
    UpdateCopywriter: '/v1/adminApi/copywriter/update',
    ListCopywriter: '/v1/adminApi/copywriter/list',
    SearchCopywriter: '/v1/adminApi/copywriter/search',
    DeleteCopywriter: '/v1/adminApi/copywriter/delete',

    AddTag: '/v1/adminApi/tag/add',
    UpdateTag: '/v1/adminApi/tag/update',
    ListTag: '/v1/adminApi/tag/list',
    DeleteTag: '/v1/adminApi/tag/delete',


    AddRobotServer: '/v1/adminApi/robot/addServer',
    DeleteRobotServer: '/v1/adminApi/robot/deleteServer',
    ListRobotServer: '/v1/adminApi/robot/listServer',

    AddCrawler: '/v1/adminApi/crawler/add',
    UpdateCrawler: '/v1/adminApi/crawler/update',
    ListCrawler: '/v1/adminApi/crawler/list',
    ListCrawlerItem: '/v1/adminApi/crawler/listRandItem',
    DeleteCrawlerItem: '/v1/adminApi/crawler/deleteItem',


    AddCommissionProduct: '/v1/adminApi/media/addCommissionProduct',
    AddSocialMediaFromCrawler: '/v1/adminApi/media/addSocialMediaFromCrawler',

    AddGalleryTagApi: '/v1/adminApi/gallery/addGalleryTag',
    ListGalleryApi: '/v1/adminApi/gallery/listGallery',

};
export default api
