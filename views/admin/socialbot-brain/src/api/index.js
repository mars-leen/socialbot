
const api = {
    Host: 'http://localhost:8080',
    UploadSingle: '/v1/adminApi/upload/single',
    Login: '/v1/adminApi/login',
    Logout: '/v1/adminApi/logout',

    AddCategory: '/v1/adminApi/category/add',
    UpdateCategory: '/v1/adminApi/category/update',
    ListCategory: '/v1/adminApi/category/list',
    DeleteCategory: '/v1/adminApi/category/delete',
    ListCategoryWithTags: '/v1/adminApi/category/listWithTags',

    AddTag: '/v1/adminApi/tag/add',
    UpdateTag: '/v1/adminApi/tag/update',
    ListTag: '/v1/adminApi/tag/list',
    DeleteTag: '/v1/adminApi/tag/delete',

    AddCrawler: '/v1/adminApi/crawler/add',
    UpdateCrawler: '/v1/adminApi/crawler/update',
    ListCrawler: '/v1/adminApi/crawler/list',
    ListCrawlerItem: '/v1/adminApi/crawler/listItem',

    AddCommissionProduct: '/v1/adminApi/media/addCommissionProduct',
    AddSocialMediaFromCrawler: '/v1/adminApi/media/addSocialMediaFromCrawler',

};
export default api
