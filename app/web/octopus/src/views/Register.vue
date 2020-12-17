<template xmlns:V-on="http://www.w3.org/1999/xhtml">
  <div class="home">
    <Header current_page="login"/>

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
import Header from "@/components/Header.vue";
import {RegisterAPI} from '@/service/user'

export default {
  name: "Register",
  components: {
    Header
  },
  data() {
    return {
      username: "",
      password: "",
    }
  },
  computed: {
    //判断手机号码
    // checkPhone: function () {
    //   return /^1\d{10}$/gi.test(this.phone)
    // },
    // ...mapState([
    //   'user_profile',
    //   'user_is_active'
    // ]),
  },
  methods: {
    // ...mapMutations([
    //   'RECORD_USERINFO',
    // ]),
    // goBack() {
    //   this.$router.go(-1);
    // },
    // phoneInputError() {
    //   if (!this.checkPhone) {
    //     this.inputPhoneErrorText = '手机格式错误'
    //   } else {
    //     this.inputPhoneErrorText = ''
    //   }
    // },
    register() {
      RegisterAPI(this.username, this.password).then(
          //     res => {
          //   if ('true' == res.data.state) {
          //     this.RECORD_USERINFO(res.data.user_profile)
          //     this.$router.push('/')
          //   }else{
          //     this.inputPasswordErrorText = "密码错误或用户不存在"
          //   }
          // }
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
