import Vue from 'vue'
import Router from 'vue-router'
import Layout from './layouts/BasicLayout.vue'
import store from "./store/";

Vue.use(Router);

const routes = [
    {
        path: '/dashboard',
        component: Layout,
        children: [
            {
                path: '',
                name: 'default',
                component: () => import('./views/dashboard/Workplace.vue'),
                meta: {
                    keepAlive: true,
                    needLogin:true,
                }
            },
            {
                path: 'workplace',
                name: 'workplace',
                component: () => import('./views/dashboard/Workplace.vue'),
                meta: {
                    keepAlive: true,
                    needLogin:true,
                }
            },
            {
                path: 'config',
                name: 'config',
                component: () => import('./views/config/Config.vue'),
                meta: {
                    keepAlive: true,
                    needLogin:true,
                }
            },
            {
                path: 'config/category',
                name: 'categoryConfig',
                component: () => import('./views/config/Category.vue'),
                meta: {
                    keepAlive: true,
                    needLogin:true,
                }
            },
            {
                path: 'config/copywriter',
                name: 'copywriterConfig',
                component: () => import('./views/config/Copywriter.vue'),
                meta: {
                    keepAlive: true,
                    needLogin:true,
                }
            },
            {
                path: 'config/tag',
                name: 'tagConfig',
                component: () => import('./views/config/Tag.vue'),
                meta: {
                    keepAlive: true,
                    needLogin:true,
                }
            },
            {
                path: 'config/server',
                name: 'serverConfig',
                component: () => import('./views/config/Server.vue'),
                meta: {
                    keepAlive: true,
                    needLogin:true,
                }
            },

            {
                path: 'add-commission-product',
                name: 'addCommissionProduct',
                component: () => import('./views/media/AddCommissionProduct.vue'),
                meta: {
                    keepAlive: true,
                    needLogin:true,
                }
            },
            {
                path: 'media/gallery',
                name: 'gallery',
                component: () => import('./views/media/Gallery.vue'),
                meta: {
                    keepAlive: true,
                    needLogin:true,
                }
            },

            {
                path: 'robot/server',
                name: 'robotServer',
                component: () => import('./views/robot/Server.vue'),
                meta: {
                    keepAlive: true,
                    needLogin:true,
                }
            },
            {
                path: 'robot/crawler-item',
                name: 'crawlerItem',
                component: () => import('./views/robot/CrawlerItem.vue'),
                meta: {
                    keepAlive: true,
                    needLogin:true,
                }
            },
        ],
    },
    {
        path: '/dashboard/login',
        name: 'login',
        component:  () => import('./views/user/Login.vue'),
        meta: {
            keepAlive: false,
            needLogin:false,
        }
    },
];

const router = new Router({
    mode: 'history',
    routes
});

router.beforeEach((to, from, next) => {
    const token = store.getters.Token();
    if (to.meta.needLogin && !token) {
        next("/dashboard/login")
    }else {
        next();
    }
});

export {
    router
};
