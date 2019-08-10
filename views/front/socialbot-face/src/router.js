import Vue from 'vue'
import Router from 'vue-router'
import BLayout from './layouts/BasicLayout.vue'
Vue.use(Router);

const routes = [
  {
    path: '',
    component: BLayout,
    children: [
      {
        path: '',
        name: 'home',
        component:  () => import('./views/Home.vue'),
        meta: {
          title: 'Century.Wedding',
          keepAlive: true,
          showNavList:true,
        }
      },
      {
        path: '/media/detail/:uri/:time?',
        name: 'detailMedia',
        component: () => import('./views/media/Detail.vue'),
        props:true,
        meta: {
          title: 'detail',
          keepAlive: false,
          showNavList:false,
        }
      },
      {
        path: '/media/category/:id',
        name: 'categoryMedia',
        component: () => import('./views/media/Category.vue'),
        props:true,
        meta: {
          title: 'Category',
          keepAlive: true,
          showNavList:true,
        }
      },
    ],
  },

];

const router = new Router({
  mode: 'history',
  routes
});


router.beforeEach((to, from, next) => {
  const title = to.meta && to.meta.title;
  if (title) {
    document.title = title
  }
  next();
});
export {
  router
};
