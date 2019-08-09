import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
Vue.use(Router);

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
  {
    path: '/media/category',
    name: 'mediaCategory',
    component: () => import('./views/media/Category.vue'),
    meta: {
      title: 'Century.Wedding',
      keepAlive: true,
    }
  },
];


const router = new Router({
  mode: 'history',
  routes
});


router.beforeEach((to, from, next) => {
  /*const title = to.meta && to.meta.title;
  if (title) {
    document.title = title
  }*/
  next();
});
export {
  router
};
