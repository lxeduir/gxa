<template xmlns="">
  <div class="dashboard-container">
    <PanelGroup @handleSetLineChartData="handleSetLineChartData" />
    <!--    <div class="dashboard-text">首页</div>-->
    <el-row style="background:#fff;margin-bottom:32px;height:200px; margin-top: 50px;">
      <el-col :xs="12" :sm="12" :lg="11" class="notice" style="margin-right:109px">
        <h3 class="title">公告</h3>
        <p style="padding:10px 5px;text-indent: 2em;">
          根据市委、市政府工作安排，我市积极推进国家园林城市建设工作。

          为体现公开、公平、公正原则，现将拟命名的2022年宜宾市园林村和园林式单位(小区)名单进行公示，接受社会各界监督。

          如有意见和建议，可在公示期内以书面形式将有关情况寄至住房和城乡建设局。

          单位反映情况要加盖公章;
          
          个人反映情况要签署真实姓名，并留下真实的联系电话、地址及邮编。</p>
        <div class="pubinfo">
          <p>发布者:<span>黄sir</span></p>
          <p>所属部门:<span>住房和城乡建设局</span></p>
          <p>发布时间:<span>2023.1.23</span></p>
        </div>
      </el-col>
      <el-col :xs="12" :sm="12" :lg="11" class="notice">
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import PanelGroup from '@/views/dashboard/components/PanelGroup'
export default {
  name: 'Dashboard',
  components: {
    PanelGroup
  },
  computed: {
    ...mapGetters([
      'name'
    ])
  }
}
</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 10px;
  }

  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}

.title {
  position: relative;
  width: 50px;
  margin: 0px auto;
  padding: 5px;
  color: red;
}

.pubinfo {
  font-size: 10px;
  position: relative;
  width: 150px;
  left: 75%;
  font-weight: 800;
  // height: 100px;
  // border: 1px solid red;
}

.notice {
  background-color: rgb(255, 255, 255);
  overflow: hidden;
  box-shadow: 4px 4px 40px rgba(0, 0, 0, .05);
  border-color: rgba(0, 0, 0, .05);
  height: 400px;
}
</style>
