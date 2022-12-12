import Vue from 'vue';
import App from './App.vue';
import ElementUI from './element/index';
import router from '@/router';
import store from '@/store';
import 'element-ui/lib/theme-chalk/index.css';
import './iconfont/iconfont.css';
import md5 from 'js-md5'
Vue.config.productionTip = false;

Vue.use(ElementUI);
Vue.prototype.$md5 = md5;
new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
