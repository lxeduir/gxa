<template>
  <div class="title">
    <el-carousel indicator-position="none" style="width:100vw; height:500px; position:relative;">
      <el-carousel-item v-for="item in 4" :key="item">
      </el-carousel-item>
    </el-carousel>
  </div>
</template>


<script>
export default {
  data() {
    return {

    };
  },
  methods: {},
};
</script>

<style scoped>
.title {
  width: 100%;
  height: 400px;
  position: absolute;
  left: 0px;
  /* background-color: red; */

}

.el-carousel__item {
  height: 400px !important;
}

.el-carousel__item h3 {
  font-size: 0px;
  opacity: 0.75;
  line-height: 500px;
  margin: 0;
}

.el-carousel__item:nth-child(2n) {
  background: url(../../assets/Yibin1.png) no-repeat;
  background-size: 100%;

}

.el-carousel__item:nth-child(2n+1) {
  background: url(../../assets/Yibin2.png) no-repeat;
  background-size: 100%;
}

.el-carousel__item:nth-child(2n+2) {
  background: url(../../assets/Yibin3p.jpg) no-repeat;
  background-size: 100%;
}
</style>