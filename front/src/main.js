// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import VueRouter from 'vue-router'
import Vuex from 'vuex'

// import App from './App'
import engines from './engines'

require('bootstrap/dist/css/bootstrap.css')

Vue.use(VueRouter)
Vue.use(Vuex)

const router = new VueRouter({
  routes: engines.routes()
})

/* eslint-disable no-new */
new Vue({
  router
}).$mount('#app')
