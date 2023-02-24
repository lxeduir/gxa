<template>
  <div class="content">
    <div class="login-content">
      <a href="" class="banner-box">政务服务平台</a>
      <div class="login-box">
        <h1>账号注册</h1>
        <input type="text" placeholder="用户名" v-model="registerFrom.username">
        <input type="text" placeholder="邮箱" v-model="registerFrom.email">
        <input type="password" placeholder="密码" v-model="registerFrom.passwd" maxlength="16">
        <input class="captcha" type="text" placeholder="验证码" v-model="registerFrom.captcha"><button class="yzm"
          @click="getyzm">获取</button>
        <button class="login-btn" id="login-btn"><span @click="userRegister" class="register">注册</span><span>|</span><span class="tologin" @click="tologin">去登录</span></button>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      registerFrom: {
        username: '',
        email: '',
        passwd: '',
        captcha: '',
      }
    }
  },
  methods: {
    async userRegister() {
      try {
        let info = {
          username: this.registerFrom.username,
          email: this.registerFrom.email,
          passwd: this.$md5(this.registerFrom.passwd.trim()),
          captcha: this.registerFrom.captcha
        }
        await this.$store.dispatch('user/register', info).then(()=>{

        })
      } catch (error) {
        console.log(error.message)
      }
    },
    async getyzm() {
      try {
        let info = {
          email: this.registerFrom.email,
          type: "signup"
        }
        await this.$store.dispatch('user/getyzm', info)
      } catch (error) {
        console.log(error.message)
      }
    },
    tologin(){
      this.$router.push({name:'login'})
    }
  }
}
</script>
<style scoped>
.yzm {
  border: none;
  border-radius: 20px;
  width: 80px;
  height: 40px;;
  background-color: #1482F1;
  color: #fff;
  position: relative;
  top: -53px;
  right: -70px;
}

.captcha {
  width: 100px!important;
  position: relative!important;
  left: -50px!important; 
}
.tologin{
  /* border: 1px solid red; */
  margin-left: 8px ;
}
.register{
  margin-right: 8px;
}
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
  background-image: url(../../assets/Yibin1.png);
}

.login-content {
  position: relative;
  width: 900px;
  height: 360px;
  background-color: rgb(234, 231, 231,0.7);
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
  color: rgb(255, 255, 255);
  position: relative;
  /* border: 1px solid black; */
  right: 150px;
  top: 100px;
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
  height: 5px;
  font-size: 36px;
  margin: 60px 0;
  color: #1482F1;
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
  width: 200px;
  height: 40px;
  background-color: rgb(146, 141, 149);
  border: none;
  border-radius: 40px;
  color: #fff;
  font-size: 18px;
  position: relative;
  top: -30px;
}

.login-btn:hover {
  background-color: #1482F1;
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
  