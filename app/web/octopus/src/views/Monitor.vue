<template>
  <div class="home">
    <Header current_page="monitor"/>
    <div class="row">
      <div class="col-md-10 col-md-offset-1">
        <div class="fh5co-blog">
          <div class="blog-text">
            <h2> CPU </h2>
            <div class="row">
              <div id="cpuChart" style="width: 100%; height: 500px;"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col-md-10 col-md-offset-1">
        <div class="fh5co-blog">
          <div class="blog-text">
            <h2> 内存 </h2>
            <div class="row">
              <div id="memChart" style="width: 100%; height: 500px;"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col-md-10 col-md-offset-1">
        <div class="fh5co-blog">
          <div class="blog-text">
            <h2> 磁盘 </h2>
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
      date: [],
      diskData: [],
    }
  },
  created() {
    let now = new Date();
    for (let i = 0; i < 24; i++) {
      this.diskData.push(10);
      this.date.push(now.getHours() + ":" + now.getMinutes());
    }
  },
  methods: {
    memChart() {
      const memChart = this.$echarts.init(document.getElementById('memChart'));
      let option = {
        xAxis: {
          type: 'category',
          data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
        },
        yAxis: {
          type: 'value'
        },
        series: [{
          data: [820, 932, 901, 934, 1290, 1330, 1320],
          type: 'line',
          smooth: true
        }]
      };
      memChart.setOption(option);
    },
    diskChart() {
      const diskChart = this.$echarts.init(document.getElementById('diskChart'));
      let option = {
        grid : {
          top : 40,
          bottom: 80
        },
        tooltip: {
          trigger: 'axis',
        },
        dataZoom: [
          {
            type: 'slider',
            start: 0, // 起始0
            end: 100, // 终止100
            minSpan: 1,
            bottom: '0%',
            dataBackground: {
              lineStyle: {
                color: '#F0F2F5'
              },
              areaStyle: {
                color: '#F0F2F5',
                opacity: 1,
              }
            }
          },
          {
            type: 'inside'//使鼠标在图表中时滚轮可用
          }
        ],
        xAxis: {
          type: 'category',
          splitNumber: 24,
          boundaryGap: false,
          data: this.date,
          splitLine: {
            show: true
          },
        },
        yAxis: {
          type: 'value',
          boundaryGap: [0, '100%'],
          splitLine: {
            show: false
          }
        },
        series: [ // 曲线设置
          {
            data: this.diskData,
            type: 'line',
            smooth: true,
            symbol: 'circle',
            areaStyle: { // 区域设置
              color: {
                type: 'linear',
                x: 0,
                y: 0,
                x2: 0,
                y2: 1,
                colorStops: [{
                  offset: 0, color: "red" // 100% 处的颜色
                }, {
                  offset: 1, color: "white"  // 0% 处的颜色
                }]
              }
            },
          },
        ],
      };
      diskChart.setOption(option);
    },
    cpuChart() {
      const cpuChart = this.$echarts.init(document.getElementById('cpuChart'));
      let option = {
        xAxis: {
          type: 'category',
          data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
        },
        yAxis: {
          type: 'value'
        },
        series: [{
          data: [820, 932, 901, 934, 1290, 1330, 1320],
          type: 'line',
          smooth: true
        }]
      };
      cpuChart.setOption(option);
    }
  },
  mounted() {
    this.diskChart();
    this.cpuChart();
    this.memChart();
  }
};
</script>
