<template>
  <div class="list">
    <div class="list1">
      <div @click="gotoSf">
        <img
          src="@/assets/identify.png"
          style="position: relative; top: 10px; width: 35px"
        />统一身份认证
      </div>
      <a></a>
      <div @click="gotoSx">
        <img
          src="@/assets/dealevent.png"
          style="position: relative; top: 10px; width: 35px"
        />统一政务公告
      </div>
      <a></a>
      <div @click="gotoPj">
        <img
          src="@/assets/evaluate.png"
          style="position: relative; top: 10px; width: 35px"
        />统一好评差评
      </div>
      <a></a>
      <div>
        <img
          src="@/assets/serve.png"
          style="position: relative; top: 10px; width: 35px"
        />统一用户服务
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {};
  },
  methods: {
    gotoPj() {
      this.$router.push("/remark");
    },
    gotoSf() {
      this.$router.push("/identify");
    },
    gotoSx() {
      this.$router.push("/notice");
    },
  },
};
</script>

<style scoped>
.list {
  width: 100vw;
  height: 15vh;
  /* background-color: red; */
  display: flex;
  align-items: center;
  justify-content: center;
  margin-top: 250px;
}
.list1 {
  display: flex;
  /* margin:0 auto; */
  /* position: absolute; */
  /* left: 500px; */
  /* margin: 0 auto; */
}
.list1 > div {
  display: inline;
  background: rgb(255, 255, 255);
  opacity: 0.7;
  width: 150px;
  height: 75px;
  font-size: 18px;
  text-align: center;
  line-height: 75px;
  z-index: 1;
  cursor: pointer;
  user-select: none;
}
.list1 > a {
  width: 100px;
}
</style>
