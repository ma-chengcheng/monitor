<template xmlns:V-on="http://www.w3.org/1999/xhtml">
  <div class="home">
    <div class="jumbotron">
      <div class="row">
        <div class="col-md-4 col-md-offset-4">
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
                  <a v-on:click="register" class="btn btn-success btn-lg btn-learn">注册</a>
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import {RegisterAPI} from '@/service/user'

export default {
  name: "Register",
  data() {
    return {
      username: "",
      password: "",
    }
  },
  methods: {
    register() {
      RegisterAPI(this.username, this.password).then(res => {
            if (200 === res.data.code) {
              this.$cookies.set('token', res.data.data['token']);
              this.$router.push('/');
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
