<template>
  <div class="home">
    <Header current_page="console"/>
    <div class="overlay"></div>
    <div class="container">

      <div class="row">
        <div class="col-xs-4 text-left fh5co-heading fh5co-nav top-menu menu-1">
          <ul>
            <li class="active">
              <router-link to="/console">资源概述</router-link>
            </li>
          </ul>
        </div>
        <div class="col-xs-offset-4 col-xs-4 text-right">
          <RouterLink class="btn btn-success btn-sm btn-learn" to="/add_node">添加主机</RouterLink>
        </div>
      </div>

      <div class="row">
        <div v-bind:key="nodeInfo.ip" v-for="nodeInfo in nodeInfoList" class="col-md-6">
          <div class="fh5co-blog">
            <div class="blog-text">
              <h2>{{ nodeInfo.host_name }}</h2>
              <div class="row">
                <div class="col-xs-6">
                  <span class="posted_on">IP <b>{{ nodeInfo.ip }}</b></span>
                </div>
                <div class="col-xs-6 text-right">
                  <span v-if="nodeInfo.status" class="posted_on text-right" style="color: green;">运行中</span>
                  <span v-else class="posted_on text-right" style="color: gray;">已停止</span>
                </div>
              </div>

              <div class="row">
                <div class="col-xs-4">
                  <h5>cpu使用率</h5>
                </div>
                <div class="col-xs-4">
                  <h5>内存使用率</h5>
                </div>
                <div class="col-xs-4">
                  <h5>磁盘使用率</h5>
                </div>
              </div>


              <div class="row text-center">
                <div class="col-xs-4">
                  <h2>{{ nodeInfo.cpu_percent }}%</h2>
                </div>
                <div class="col-xs-4">
                  <h2>{{ nodeInfo.memory_percent }}%</h2>
                </div>
                <div class="col-xs-4">
                  <h2>{{ nodeInfo.disk_percent }}%</h2>
                </div>
              </div>

              <ul class="stuff text-right">
                <router-link :to="{ name: 'Monitor', params: { ip: nodeInfo.ip }}" class="btn btn-sm btn-info">
                  详情
                </router-link>
                <a v-on:click="deleteNode(nodeInfo.ip)" class="btn btn-sm btn-danger">删除</a>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Header from "@/components/Header.vue";
import {DeleteNodeAPI} from "@/service/node.js"

export default {
  name: "Console",
  components: {
    Header
  },
  data() {
    return {
      nodeInfoList: []
    }
  },
  created() {
    this.initWebSocket();
  },
  destroyed() {
    this.websock.close();
  },
  methods: {
    initWebSocket() {
      const wsuri = "ws://127.0.0.1:8082/api/v1/ws/query";
      this.websock = new WebSocket(wsuri);
      this.websock.onmessage = this.webSocketOnMessage;
      this.websock.onerror = this.webSocketOnError;
    },
    webSocketOnError() {
      this.initWebSocket();
    },
    webSocketOnMessage(e) {
      let data = JSON.parse(e.data);
      this.nodeInfoList = data["node_info_list"]
    },
    deleteNode(ip) {
      DeleteNodeAPI(ip).then(
          res => {
            if (200 !== res.data.code) {
              console.log(res.data.msg);
            }
          }
      ).catch(function (error) {
        console.log(error);
      });
    }
  },
};
</script>
