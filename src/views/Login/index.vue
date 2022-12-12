<template>
  <div class="content">
    <div class="login-content">
      <a href="" class="banner-box">宜宾政务服务平台</a>
      <p>测试邮箱:3021955171@qq.com 密码:111111</p>
      <div class="login-box">
        <h1>账号登录</h1>
        <input type="text" placeholder="邮箱" v-model="loginForm.email">
        <input type="password" placeholder="密码" v-model="loginForm.passwd">
        <button class="login-btn" id="login-btn"><span @click="userlogin">登录</span></button>
      </div>
    </div>
  </div>
</template>
<script>
import { Message } from 'element-ui';
export default {
  name: 'Login',
  data() {
    return {
      loginForm: {
        email: '',
        passwd: ''
      }
    }
  },
  methods: {
    async userlogin() {
      if (!this.loginForm.email) {
        Message({
          message: "请输入用户名",
          type: 'error',
          duration: 3 * 1000
        })
        return;
      } else if (!this.loginForm.passwd) {
        Message({
          message: "请输入密码",
          type: 'error',
          duration: 3 * 1000
        })
        return
      }
      await this.$store.dispatch("user/login", this.loginForm).then(() => {this.$router.push('/')}
      ).catch((error) => { console.log(error.message) })
    }
  }
}
</script>
<style scoped>
.x {
  margin-left: 2px solid white;
}

a {
  text-decoration: none;
}

.content {
  /* border: 1px solid red; */
  width: 100vw;
  height: 100vh;
  background-size: cover;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: rgb(221, 236, 241);
}

.login-content {
  position: relative;
  width: 800px;
  height: 360px;
  background-color: rgb(173, 236, 242);
  border-radius: 40px;
  padding: 30px;
}
.banner-box {
  display: inline-block;
  width: 500px;
  height: 100%;
  background-size: cover;
  border-radius: 20px 0 0 20px;
  text-align: center;
  font-size: 80px;
  color: rgb(93, 93, 93);
  position: relative;
  /* border: 1px solid black; */
  right: 150px;
}

.login-box {
  width: 330px;
  height: 100%;
  position: absolute;
  top: 0;
  right: 0;
  border-radius: 0 30px 30px 0;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.login-box>h1 {
  font-size: 36px;
  margin: 60px 0;
  color: rgb(140, 131, 131);
}

.login-box>input {
  margin-bottom: 20px;
  width: 200px;
  height: 36px;
  padding: 0 20px;
  border-radius: 36px;
  border: none;
  outline: none;
  font-size: 18px;
}

.login-btn {
  margin-top: 20px;
  width: 200px;
  height: 40px;
  background-color: rgb(146, 141, 149);
  border: none;
  border-radius: 40px;
  color: #fff;
  font-size: 18px;
}
.login-btn:hover {
  background-color: rgb(82, 87, 90);
}
@media screen and (max-width: 900px) {
  .login-content {
    margin: 0 30px;
  }
  .banner-box {
    border-radius: 20px;
  }
}
@media screen and (max-width: 600px) {
  .banner-box {
    display: none;
  }
  .login-box {
    position: initial;
    margin: -40px auto;
    background-color: transparent;
  }
}
</style>
    