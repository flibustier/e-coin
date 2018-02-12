<template>
    <div>
        <div v-if="isLoggedIn">
            <navbar :assets="assets" />
            <section class="section columns">
                <div class="column is-one-quarter">
                    <sidebar />
                </div>
                <div class="column">
                    <router-view :assets="assets" v-on:transaction="update"></router-view>
                </div>
            </section>
        </div>
        <div v-else>
            <router-view></router-view>
        </div>
    </div>
</template>

<script>
import navbar from "./components/Navbar.vue";
import sidebar from "./components/Sidebar.vue";
import router from "./router";
import api from "./api";
import config from "./config.json";
import { isLoggedIn, login, logout } from './auth';

export default {
	router,

	data() {
		return {
			assets: config.assets,
		};
	},

	components: {
		navbar,
		sidebar,
	},

	methods: {
		fetch() {
			if (isLoggedIn()) {
				api
					.balance()
					.then(fetched => this.assets = fetched)
					.catch(_ => logout());
			}
		},
		update() {
			//console.log("Received event transaction, update balance…");
			this.fetch();
		},
	},

	computed: {
		isLoggedIn: function() {
			return isLoggedIn();
		},
	},

	mounted() {
		this.fetch();
	},
};
</script>
