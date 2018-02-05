import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router);

import auth from './auth'
import Home from './components/Home.vue'
import Login from './components/Login.vue'
import Transfer from './components/Transfer.vue'
import History from './components/History.vue'

const routes = [
    { path: '/', component: Home, props: true, beforeEnter: auth.requireAuth },
    { path: '/login', component: Login },
    { path: '/history', component: History, beforeEnter: auth.requireAuth  },
    { path: '/transfer', component: Transfer, props: true, beforeEnter: auth.requireAuth },
];

// export router instance
export default new Router({
    mode: 'history',
    routes,
    linkActiveClass: 'is-active'
})
