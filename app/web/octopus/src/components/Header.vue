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
              <li>
                <router-link to="/console">{{ username }}</router-link>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<script>
import jwtDecode from 'jwt-decode'

export default {
  name: "Header",
  data() {
    return {
      username: ''
    }
  },
  props: ['current_page'],
  created() {
    let token = this.$cookies.get('token');
    if (token) {
      token = jwtDecode(token);
      this.username = token.username;
    } else {
      this.$router.push('/');
    }
  }
};
</script>
