import Vue from 'vue'
import app from './App.vue'
import router from './router'

import VueParticles from 'vue-particles'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'

Vue.use(VueParticles);
Vue.use(Buefy);

import { logout } from './auth'
import axios from 'axios'

axios.interceptors.response.use((response) => {
    return response
}, function (error) {
    // Do something with response error
    if (error.response.status === 401) {
        logout();
        router.replace('/login')
    }
    return Promise.reject(error)
});

new Vue({
    el: '#app',
    components: {
        app
    }
});