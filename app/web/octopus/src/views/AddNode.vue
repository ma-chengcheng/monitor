<template>
  <div>
    <Header current_page="add_server"/>
    <header id="fh5co-header" class="fh5co-cover js-fullheight" role="banner">
      <div class="overlay"></div>
      <div class="container">

        <div class="row">
          <div class="col-xs-4 text-left fh5co-heading fh5co-nav top-menu menu-1">
            <ul>
              <li>
                <router-link to="/console">资源概述</router-link>
              </li>
            </ul>
          </div>
        </div>

        <div class="row">
          <div class="col-md-8 col-md-offset-2 text-center fh5co-heading">
            <h2>添加主机</h2>
          </div>

          <div class="row text-center">
            <div class="col-md-4 col-md-offset-4">
              <div class=" form-group">
                <input v-model="host_name" type="text" class="form-control" placeholder="Host Name">
              </div>
              <div class=" form-group">
                <input v-model="ip" type="text" class="form-control" placeholder="IP">
              </div>
              <div class="form-group">
                <input v-model="ssh_username" type="text" class="form-control"
                       placeholder="SSH Username">
              </div>
              <div class="form-group">
                <input v-model="ssh_password" type="password" class="form-control"
                       placeholder="SSH Password">
              </div>
              <p><a class="btn btn-primary btn-lg btn-demo" v-on:click="addNode">添加</a></p>
            </div>
          </div>
        </div>
      </div>
    </header>
  </div>
</template>

<script>
import Header from "@/components/Header";
import {AddNodeAPI} from "@/service/node.js"

export default {
  name: "AddServer",
  components: {
    Header
  },
  data() {
    return {
      ip: "",
      host_name: "",
      ssh_username: "",
      ssh_password: ""
    }
  },
  methods: {
    addNode() {
      AddNodeAPI(this.ip, this.host_name, this.ssh_username, this.ssh_password).then(
          res => {
            if (200 === res.data.code) {
              this.$cookies.set('token', res.data.data['token'])
              this.$router.push('/console');
            } else {
              console.log(res.data.msg);
            }
          }
      ).catch(function (error) {
        console.log(error);
      });
    }
  }
}
</script>
