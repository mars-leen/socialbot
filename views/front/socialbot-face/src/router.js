import Vue from 'vue'
import Router from 'vue-router'
import Home from './pages/Home.vue'
Vue.use(Router);


// media
const MediaCategory = () => import('./pages/media/Category.vue');
const MediaDetail = () => import('./pages/media/Detail.vue');

//user

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home,
    meta: {
      title: 'Century.Wedding',
      keepAlive: true,
    }
  },

  // media
  {
    path: '/media/category/:id',
    name: 'categoryMedia',
    component: MediaCategory,
    props:true,
    meta: {
      title: 'Category',
      keepAlive: true,
    }
  },
  {
    path: '/media/detail/:uri/:time?',
    name: 'detailMedia',
    component: MediaDetail,
    props:true,
    meta: {
      title: 'detail',
      keepAlive: false,
    }
  },
  // user
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
