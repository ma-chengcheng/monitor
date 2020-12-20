import Vue from "vue";
import VueRouter from "vue-router";
import Login from "../views/Login.vue";
import Register from "../views/Register.vue";
import Console from "../views/Console.vue";
import AddNode from "../views/AddNode.vue";
import Monitor from "../views/Monitor.vue";


Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Login",
    component: Login
  },
  {
    path: "/register",
    name: "Register",
    component: Register
  },
  {
    path: "/console",
    name: "Console",
    component: Console
  },
  {
    path: "/add_node",
    name: "AddNode",
    component: AddNode
  },
  {
    path: "/monitor",
    name: "Monitor",
    component: Monitor
  }
];

const router = new VueRouter({
  routes
});

export default router;
