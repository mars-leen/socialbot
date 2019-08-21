import Vue from 'vue'
import App from './App.vue'
import {router} from './router.js'
import {store} from './store/store.js'
import VueStorage from 'vue-ls'

import './assets/style/style.css'
import VueMasonry from 'vue-masonry-css'

Vue.config.productionTip = false;

Vue.use(VueMasonry);
Vue.use(VueStorage);

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app');
