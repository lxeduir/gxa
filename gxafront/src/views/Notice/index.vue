<template>
  <div id="Notice" v-if="isreload">
    <div class="picture">
      <a class="text0">政务公告</a>
      <a class="text1">让权力在阳光下进行</a>
      <a class="text2">提升政府工作效能，优化营商环境，建设人民满意的服务型政府</a>
    </div>
    <div class="detail">
      <Message></Message>
    </div>
  </div>
</template>

<script>
import Message from "../../components/Message";

export default {
  components: {
    Message,
  },
  data() {
    return {
      isreload: false,
    };
  },
  
  created() {
    this.getNotice();
    // console.log(this.notice)
  },

  methods: {
    async getNotice() {
      try {
        const res = await this.$store.dispatch("user/reqPublicInfo");
        this.isreload = true;
        // console.log("getNotice")
      } catch (error) {
        console.log(error.message);
      }
      // window.location.reload()
      // this.$router.go(0)
    },
  },
  watch: {
    $router(to, from) {
      window.location.reload();
    },
  },
};
</script>

<style scoped>
.picture {
  height: 300px;
  width: 1540px;
  /* background-color: #6db6ff; */
  background-image: url("@/assets/zhuhai.jpg");
}
.text0 {
  position: absolute;
  top: 150px;
  left: 100px;
  font-size: 30px;
  color: white;
}
.text1 {
  position: absolute;
  top: 200px;
  left: 100px;
  font-size: 20px;
  color: white;
}
.text2 {
  position: absolute;
  top: 230px;
  left: 100px;
  font-size: 20px;
  color: white;
}
.detail {
  background-color: white;
  position: relative;
  top: -120px;
  left: 170px;
  width: 1200px;
  /*height: 350px;*/
  border-radius: 10px;
  border: 1px solid #b3d8fc;
}
.detail div {
  display: inline;
  margin: 100px;
  position: relative;
  top: 50px;
  font-size: 30px;
}
.detail a {
  font-size: 26px;
  position: relative;
  top: 50px;
  left: -90px;
}
.toremark {
  position: relative;
  top: -130px;
  left: 100px;
  background-color: #1482f1;
  color: white;
  border-radius: 5px;
  width: 120px;
  height: 40px;
  line-height: 40px;
}
</style>
