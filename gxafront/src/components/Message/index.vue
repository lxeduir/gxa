<template>
  <div>
    <!-- <div class="hot">
      <p>热点</p>
      <div>
        <img src="">
        <div class="hot_content">{{notice.content}}</div>
      </div>
    </div> -->
    <div class="notice">
      <p>通知公告</p>
      <div style="border: 1px solid black"></div>
      <div v-for="(item,index) in notice" :key="index">
        <div class="notice_content">
          <p>{{ item.title }}</p>
          <a>{{ item.content }}</a>
        </div>
        <button class="more" @click="getitem(index)">了解更多</button>
        <div class="notice_time">
          <p>{{ item.begin_time }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import store from "@/store";
// import MarkDown from "markdown";

export default {
  // components: { MarkDown },
  data() {
    return {
      notice: {
        title: "",
        content: "",
        begin_time: "",
        end_time: "",
        order: "",
      },
    };
  },

  created() {
    this.notice = this.$store.getters.notice;
  },

  methods: {
    getitem(index) {
      this.$router.push({ name: 'notdetail' ,params: {index: index}});
      // console.log(index)
    }
  },
};
</script>

<style lang="less" scoped>
.hot > p {
  position: relative;
  top: -60px;
  left: -530px;
  font-size: 20px;
}
.hot > div {
  border-top: 2px solid black;
  position: relative;
  top: -50px;
  left: 25px;
  width: 1150px;
}
.hot > img {
  float: left;
}
.hot_content {
  float: right;
}
.notice > p {
  position: relative;
  top: -60px;
  left: -530px;
  font-size: 20px;
}
.notice > div {
  /*border-top: 2px solid black;*/
  position: relative;
  top: -50px;
  left: 25px;
  width: 1150px;
}
.notice_content {
  border-right: 2px solid #1482f1;
  width: 1000px;
  position: relative;
  margin: 10px;
  margin-top: 30px;
  /*height: 20px;*/
  p {
    /*border: 1px solid red;*/
    color: #1482f1;
    font-size: 25px;
    text-align: left;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }
  a {
    /*border: 1px solid green;*/
    font-size: 20px;
    display: block;
    text-align: left;
    word-wrap: break-word;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }
}
.notice_time {
  /*border: 2px solid blue;*/
  position: relative;
  top: -25px;
  width: 140px;
  /*height: 20px;*/
  float: right;
  font-size: 10px;
}
.more {
  position: relative;
  left: 580px;
  top: -70px;
  width: 100px;
  height: 30px;
  color: white;
  background-color: #1482f1;
  border: none;
}
</style>
