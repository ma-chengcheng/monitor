<template xmlns:V-on="http://www.w3.org/1999/xhtml">
  <div class="home">

    <div class="jumbotron">
      <div class="row">
        <div class="col-md-4 col-md-offset-2">
          <img src="@/assets/logo.png"/>
        </div>

        <div class="col-md-offset-1 col-md-3">
          <blockquote class="blockquote-reverse">
            <p>一套开源的集群监管系统。</p>
            <footer>
              <cite title="八爪鱼">八爪鱼</cite>
            </footer>
          </blockquote>
          <div class="row text-center m-auto" style="padding-top: 100px;">
            <div class="col-md-12">
              <div class="form-group">
                <input
                    v-model="username"
                    type="text"
                    class="form-control"
                    placeholder="Username"
                />
              </div>
              <div class="form-group">
                <input
                    v-model="password"
                    type="password"
                    class="form-control"
                    placeholder="Password"
                />
              </div>
              <div class="form-group">
                <p>
                  <a v-on:click="login" class="btn btn-primary btn-lg btn-demo">登陆</a>
                  <router-link class="btn btn-success btn-lg btn-learn" to="/register">注册</router-link>
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="row copyright">
      <div class="col-md-12 text-center">
        <p>
          <small class="block">&copy; 2020 machengcheng 所有权</small>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import {LoginAPI} from '@/service/user'

export default {
  name: "Login",
  data() {
    return {
      username: "",
      password: "",
    }
  },
  created() {
    let token = this.$cookies.get('token');
    console.log(token);
  },
  methods: {
    login() {
      LoginAPI(this.username, this.password).then(res => {
          if (200 === res.data.code) {
            this.$cookies.set('token', res.data.data['token']);
            this.$router.push('/console');
          }
        }
      )
    }
  }
};
</script>

<style>
.jumbotron {
  background-color: white;
}
</style>
