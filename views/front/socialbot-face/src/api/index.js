const api = {
    Host: 'http://localhost:8080',
    Login: '/v1/api/user/login',
    Register: '/v1/api/user/register',
    EditProfile: '/v1/api/user/editProfile',

    ListCategory: '/v1/api/category/list',

    HomeRecommend: '/v1/api/media/homeRecommend',
    ListMediaByCategory: '/v1/api/media/listByCategory',

    MediaDetail: '/v1/api/media/detail',
    LikeMedia: '/v1/api/media/like',
};

export default api
