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
              <div class="col-xs-6 left">用户名</div>
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
              <div id="cpuChart" style="width: 100%; height: 500px;"></div>
            </div>
            <h3>内存使用率</h3>
            <div class="row">
              <div id="memChart" style="width: 100%; height: 500px;"></div>
            </div>
            <h3>磁盘使用率</h3>
            <div class="row">
              <div id="diskChart" style="width: 100%; height: 500px;"></div>
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
      nodeInfo: this.$route.params.nodeInfo
    }
  },
  created() {
      // let array = ['cpuChart', 'memChart', 'diskChart'];
      // array.forEach(name=>{
      //   console.log(name);
      // });
  },
  methods:{
    plot(name){
      // 基于准备好的dom，初始化echarts实例
      var chart = this.$echarts.init(document.getElementById(name));

      // 指定图表的配置项和数据
      let option = {
        xAxis: {
          type: 'category',
          data: ['60秒', '50', '40', '30', '20', '10', '0'],
          boundaryGap: false,
        },
        yAxis: {
          type: 'value',
          splitLine: {
            show: false
          }
        },
        series: [{
          data: [0, 0, 0, 0, 0, 0, 0],
          type: 'line',
          smooth: true
        }]
      };


      // 使用刚指定的配置项和数据显示图表。
      chart.setOption(option);
    }
  },
  mounted() {
    var array = ['cpuChart', 'memChart', 'diskChart'];
    array.forEach(name => {
      this.plot(name);
    });
  }
};
</script>
