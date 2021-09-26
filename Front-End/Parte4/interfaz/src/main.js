import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify'
//import axios from 'axios'
//import VueAxios from 'vue-axios'

//Vue.use(VueAxios, axios)
//Vue.prototype.$http = axios;

Vue.config.productionTip = false
Vue.use(vuetify);

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
