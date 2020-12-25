<template>
  <div class="home">
    <Header current_page="monitor"/>
    <div class="row">
      <div class="col-xs-3 col-md-offset-1">
        <div class="fh5co-blog">
          <div class="blog-text">
            <h4>详情</h4>
            <div class="row">
              <div class="col-xs-6 left">IP</div>
              <div class="col-xs-6 text-right">{{ nodeInfo.ip }}  </div>
            </div>
            <div class="row">
              <div class="col-xs-6 left">主机名</div>
              <div class="col-xs-6 text-right">{{ nodeInfo.host_name }}  </div>
            </div>
            <div class="row">
              <div class="col-xs-6 left">用户名</div>
              <div class="col-xs-6 text-right">{{ nodeInfo.ssh_username }}</div>
            </div>
            <div class="row">
              <div class="col-xs-6 left">密码</div>
              <div class="col-xs-6 text-right">{{ nodeInfo.ssh_password }}</div>
            </div>
            <div class="row">
              <div class="col-xs-6 left">状态</div>
              <div v-if="nodeInfo.status" class="col-xs-6 posted_on text-right" style="color: green;">运行中</div>
              <div v-else class="col-xs-6 posted_on text-right" style="color: gray;">已停止</div>
            </div>
          </div>
        </div>
      </div>
      <div class="col-xs-7">
        <div class="fh5co-blog">
          <div class="blog-text">
            <h3>CPU使用率</h3>
            <div class="row">
              <div id="cpu" style="width: 100%; height: 500px;"></div>
            </div>
            <h3>内存使用率</h3>
            <div class="row">
              <div id="mem" style="width: 100%; height: 500px;"></div>
            </div>
            <h3>磁盘使用率</h3>
            <div class="row">
              <div id="disk" style="width: 100%; height: 500px;"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Header from "@/components/Header.vue";

export default {
  name: "Monitor",
  components: {
    Header,
  },
  data() {
    return {
      nodeInfo: '',
      chartData: {
        cpu: [0, 0, 0, 0, 0, 0, 0],
        disk: [0, 0, 0, 0, 0, 0, 0],
        mem: [0, 0, 0, 0, 0, 0, 0]
      },
      ws: ''
    }
  },
  created() {

  },
  methods: {
    plot (name) {
      var chart = this.$echarts.init(document.getElementById(name));
      var option = {
        xAxis: {
          type: 'category',
          data: ['60秒', '50', '40', '30', '20', '10', '0'],
          boundaryGap: false,
        },
        yAxis: {
          type: 'value',
          splitNumber: 11,
          in: 0,
          max: 100,
          splitLine: {
            show: false
          }
        },
        series: [{
          data: this.chartData[name],
          type: 'line',
          smooth: true
        }]
      };
      chart.setOption(option);
    }
  },
  mounted() {


    const wsuri = "ws://127.0.0.1:8080/api/v1/ws/query";
    this.ws = new WebSocket(wsuri);
    this.ws.onmessage = (e) => {
      let data = JSON.parse(e.data);
      for (let i in data['node_info_list']) {
        if (data['node_info_list'][i]['ip'] === this.$route.params.ip) {
          this.nodeInfo = data['node_info_list'][i];
          this.chartData['cpu'].shift()
          this.chartData['cpu'].push(data['node_info_list'][i]['cpu_percent'])
          this.chartData['mem'].shift()
          this.chartData['mem'].push(data['node_info_list'][i]['memory_percent'])
          this.chartData['disk'].shift()
          this.chartData['disk'].push(data['node_info_list'][i]['disk_percent'])
          this.plot('mem');
          this.plot('cpu');
          this.plot('disk');
        }
      }
    }
    this.plot('mem');
    this.plot('cpu');
    this.plot('disk');
  },
  beforeDestroy() {
    this.ws.close();
  }
};
</script>
