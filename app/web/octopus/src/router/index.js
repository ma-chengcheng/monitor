import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/login",
    name: "Login",
    component: () =>
      import(/* webpackChunkName: "login" */ "../views/Login.vue")
  },
  {
    path: "/register",
    name: "Register",
    component: () =>
        import(/* webpackChunkName: "register" */ "../views/Register.vue")
  },
  {
    path: "/console",
    name: "Console",
    component: () =>
        import(/* webpackChunkName: "console" */ "../views/Console.vue")
  },
  {
    path: "/add_node",
    name: "AddNode",
    component: () =>
        import(/* webpackChunkName: "add_node" */ "../views/AddNode.vue")
  },
  {
    path: "/monitor",
    name: "Monitor",
    component: () =>
        import(/* webpackChunkName: "monitor" */ "../views/Monitor.vue")
  }
];

const router = new VueRouter({
  routes
});

export default router;
