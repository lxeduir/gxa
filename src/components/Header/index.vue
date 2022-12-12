<template>
  <div class="header1">
    <el-menu :default-active="activeIndex2" class="el-menu-demo" mode="horizontal" @select="handleSelect"
      background-color="#fff" text-color="black" active-text-color="#1482F1">
      <img src="@/assets/headlogo.png" border="none" style="max-width:none; float:left;">
      <el-menu-item index="1" label="首页" @click="handleClick(1)">首页</el-menu-item>
      <el-menu-item index="2" labe="个人服务" @click="handleClick(2)">个人服务</el-menu-item>
      <el-menu-item index="3" label="法人服务" @click="handleClick(3)">法人服务</el-menu-item>
    </el-menu>
    <div class="selector">
      <input class="select" type="text" placeholder="请输入您需要办理的事项"><span>检索</span>
    </div>
    <div class="login" v-if="!username">
      <div><a @click="tologin">登录</a></div>
      <div><span class="register"></span><a @click="toregister">注册</a></div>
    </div>
    <div class="personal" v-else >
      <a class="name">{{ username }}</a>
      <a class="logout" @click="logout">退出登录</a>
    </div>
    <div class="nav">
      当前位置：{{ navigation }}
    </div>
  </div>
</template>

<script>
import {getUid} from "@/utils/auth"

export default {
  data() {
    return {
      navigation: '首页',
      str: ['首页', '个人服务', '法人服务'],
      activeIndex: '1',
      activeIndex2: '1'
    };
  },

  methods: {
    handleSelect(key, keyPath) {
    },
    handleClick(idx) {
      this.navigation = this.str[idx - 1]
      if (this.navigation == '个人服务') this.$router.push({ name: 'personal' })
      if (this.navigation == '首页') this.$router.push({ name: 'home' })
      if (this.navigation == '法人服务') this.$router.push({ name: 'company' })
    },
    toregister() {
      var self = this;
      console.log("register")
      self.$router.push({ name: 'register' })
    },
    tologin() {
      this.$router.push({ name: 'login' })
    },
    logout(){
      this.$store.dispatch('user/logout')
    }
  },
  computed:{
    username(){
      return this.$store.state.user.username
    }
  },
}
</script>

<style scoped>
.header1 {
  width: 100vw;
  height: 100px;
}

.el-menu-item {
  font-size: 20px !important;
}

.personal {
  /* border: 1px solid red; */
  width: 200px;
  position: relative;
  right: -200px;
  float: right;
  top: -45px;
  display: flex;
}
.personal .name{
  position: relative;
  top: 3px;
}
.personal .logout{
  position: relative;
  width: 70px;
  left: 5px;
}
.login {
  width: 100px;
  position: relative;
  right: -200px;
  float: right;
  top: -45px;
  display: flex;
}

.login>div {
  cursor: pointer;
  margin-left: 10px;
}

.login a:hover {
  background-color: rgb(200, 239, 244);
}

.register {
  width: 50px;
  border-left: 2px solid #c1bbb7;
  padding-left: 10px;
}

.nav {
  position: absolute;
  top: 70px;
  left: -45vw;
  font-size: 18px;
}

.selector {
  position: relative;
  right: 300px;
  float: right;
  top: -50px;
  display: flex;
}

.select {
  height: 30px;
  border: 0;
  outline: none;
  background-color: rgb(240, 240, 240);
}

.selector span {
  margin-left: 10px;
  padding: 5px;
  background-color: rgb(244, 183, 102);
}
</style>
