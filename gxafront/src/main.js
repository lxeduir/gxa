import md5 from 'js-md5'
import Vue from 'vue';
import App from './App.vue';
import router from '@/router';
import store from '@/store';
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import './iconfont/iconfont.css';
import markdown from 'vue-markdown'
Vue.config.productionTip = false;
Vue.prototype.$md5 = md5;

Vue.use(ElementUI);
Vue.use(markdown)
new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')


  

