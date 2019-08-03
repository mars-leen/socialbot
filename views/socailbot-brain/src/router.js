import Vue from 'vue'
import Router from 'vue-router'
import Layout from './layout/Layout.vue'
import Blank from './pages/Blank.vue'
Vue.use(Router);

const Login = () => import('./pages/Login');
const Home = () => import('./pages/Home');
const Crawler = () => import('./pages/Crawler');
const Media = () => import('./pages/Media');

const routes = [
  {
    path: '/',
    name: 'blank',
    component: Blank,
  },
  {
    path: '/dashboard',
    component: Layout,
    children: [
      {
        path: '',
        name: 'home',
        component: Home
      },
      {
        path: 'crawler',
        name: 'crawler',
        component: Crawler
      },
      {
        path: 'media',
        name: 'media',
        component: Media
      }
    ],
  },
  {
    path: '/dashboard/login',
    name: 'login',
    component: Login,
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