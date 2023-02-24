<template>
  <div id="Progress">
    <div class="picture">
      <a>政务公告</a>
    </div>
    <div class="background">
      <div class="detail">
        <a>{{ notice[index].title }}</a>
        <VueMarkdown>{{ notice[index].content }}</VueMarkdown>
      </div>
    </div>
  </div>
</template>

<script>
import { VueMarkdown } from 'vue-markdown';

export default {
    components: {
        VueMarkdown
    },
  data() {
    return {
      index: "0",
      notice: {
        title: "",
        content: "",
        begin_time: "",
        end_time: "",
        order: "",
      },
    };
  },
  watch: {
    $route: function (newVal, oldVal) {
      if (newVal != oldVal) {
        this.getnotindex();
      }
    },
  },
  created() {
    this.notice = this.$store.getters.notice;
  },
  methods: {
    getnotindex() {
      var res = this.$route.params.index;
      this.index = res;
    //   console.log("detail" + res);
    },
  },
};
</script>

<style scoped>
.picture {
  height: 300px;
  width: 1540px;
  /* background-color: #6db6ff; */
  background-image: url("@/assets/李庄.png");
}
.picture > a {
  position: absolute;
  top: 200px;
  left: 100px;
  font-size: 30px;
  color: white;
}
.background {
  /*border: 1px solid rgb(190, 189, 189);*/
  position: relative;
  top: -120px;
  left: 170px;
  width: 1200px;
  background-color: white;
  box-shadow: 2px 2px 2px 2px #c7c7c7;
  overflow: hidden;
}
.detail {
  /*border: 1px solid red;*/
  position: relative;
  top: 0px;
  font-size: 24px;
  margin-top: 30px;
  margin-bottom: 30px;
  margin-left: 50px;
  margin-right: 50px;
}
</style>
