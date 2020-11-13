import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

import { requireAuth } from "./auth";
import Callback from "./components/Callback.vue";
import Home from "./components/Home.vue";
import Login from "./components/Login.vue";
import Transfer from "./components/Transfer.vue";
import History from "./components/History.vue";

const routes = [
  { path: "/", component: Home, props: true, beforeEnter: requireAuth },
  { path: "/login", component: Login },
  { path: "/callback", component: Callback },
  { path: "/history", component: History },
  {
    path: "/transfer",
    component: Transfer,
    props: true,
  },
];

// export router instance
export default new Router({
  mode: "history",
  routes,
  linkActiveClass: "is-active",
});
