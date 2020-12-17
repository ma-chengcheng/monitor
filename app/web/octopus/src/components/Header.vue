<template>
  <nav class="fh5co-nav" role="navigation">
    <div class="top-menu">
      <div class="container">
        <div class="row">
          <div class="col-xs-2">
            <div id="fh5co-logo">
              <router-link to="/">八爪鱼</router-link>
            </div>
          </div>
          <div class="col-xs-10 text-right menu-1">
            <ul>
              <li v-bind:class="{'active': current_page === 'home'}">
                <router-link to="/">首页</router-link>
              </li>
              <li v-bind:class="{'active': current_page === 'login'}">
                <router-link to="/login">登陆</router-link>
              </li>
              <li v-bind:class="{'active': current_page === 'console'}">
                <router-link to="/console">控制台</router-link>
              </li>
              <span v-show="token === ''">1</span>
              <span v-show="token !== ''">2</span>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: "Header",
  data() {
    return {
      token: ''
    }
  },
  props: ['current_page'],
  created() {
    this.token = this.$cookies.get('token')
  }
};
</script>
