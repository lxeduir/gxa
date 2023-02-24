<template>
  <div class="content">
    <div class="login-content">
      <a href="" class="banner-box">政务服务平台</a>
      <p>测试邮箱:3021955171@qq.com 密码:666666</p>
      <div class="login-box">
        <h1>账号登录</h1> 
        <el-form :model="loginForm" :rules="rules">
          <el-form-item prop="email">
            <el-input type=" text" placeholder="邮箱" v-model="loginForm.email" clearable></el-input>
          </el-form-item>
          <el-form-item v-if="this.forget == false" prop="passwd">
            <el-input type="password" placeholder="密码" v-model="loginForm.passwd" show-password></el-input>
          </el-form-item>
          <el-form-item v-if="this.forget == true" prop="passwd">
            <el-input  type="password" placeholder="新密码" v-model="loginForm.passwd" show-password></el-input>
          </el-form-item>
        </el-form>     
        <a class="forget-btn" id="login-btn" v-if="this.forget == false" @click="changeForget">忘记密码</a>
        <a class="register-btn" v-if="this.forget == false" @click="toregister">立即注册</a>
        <el-input v-if="this.forget == true" class="captcha" type="text" placeholder="验证码" v-model="loginForm.captcha"></el-input>
        <button v-if="this.forget == true" class="yzm" @click="getyzm">获取</button>
        <button v-else class="login-btn" id="login-btn" @click="userlogin"><span>登录</span></button>
        <button v-if="this.forget == true" class="reset-btn" id="login-btn"
          @click="resetPasswd"><span>更改密码</span></button>
        <!-- <button v-else class="login-btn" id="login-btn" @click="changeForget"><span>忘记密码</span></button> -->
        <button v-if="this.forget == true" class="tologin-btn" id="login-btn"
          @click="changeForget"><span>返回登录</span></button>
      </div>

    </div>

  </div>
</template>
<script>
import { Message } from 'element-ui';
import { reqForgetPasswd } from '@/api/user'
export default {
  name: 'Login',
  data() {
    return {
      loginForm: {
        email: '',
        passwd: '',
        captcha: ''
      },
      forget: false,
      rules: {
        email: [
          {required: true, message: '请输入邮箱', trigger: 'change' },
          {
            pattern: /^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/,
            message: "请输入正确的邮箱",
            trigger: "change",
          }
        ],
        /*newpasswd: [
          { required: true, message: '请输入新密码', trigger: 'change' },
          {
            pattern: /^(?![0-9]+$)(?![a-z]+$)(?![A-Z]+$)(?!([^(0-9a-zA-Z)])+$).{6,18}$/,
            message: "密码至少包含：数字,英文,字符中的两种以上,长度6-18",
            trigger: "change",
          }  
        ],
        passwd: [
          { required: true, message: '请输入密码', trigger: 'change' },
        ]*/
      }
    }
  },
  created() {
    this.forget = false
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
      let from = {
        email: this.loginForm.email,
        passwd: this.$md5(this.loginForm.passwd.trim())
      }
      await this.$store.dispatch("user/login", from).then(() => {
        this.$router.push('/')
      }
      ).catch((error) => { console.log(error.message) })
    },
    async getyzm() {
      try {
        let info = {
          email: this.loginForm.email.trim(),
          type: "forget"
        }
        await this.$store.dispatch('user/getyzm', info)
      } catch (error) {
        console.log(error.message)
      }
    },
    async resetPasswd() {
      try {
        let from = {
          email: this.loginForm.email.trim(),
          passwd: this.$md5(this.loginForm.passwd.trim()),
          captcha: this.loginForm.captcha.trim()
        }
        await reqForgetPasswd(from)
        Message({
          message: "修改成功！",
          type: 'success',
          duration: 3 * 1000
        })
        this.forget = !this.forget
      } catch (error) {
        console.log(error.message)
      }
      this.loginForm.passwd = ''
    },
    changeForget() {
      this.forget = !this.forget,
      this.loginForm.passwd = ''
    },
    toregister() {
      this.$router.push({ name: 'register' })
    }
  }
}
</script>
<style lang="less" scoped>
/deep/.el-form-item {
  /*border: 1px solid red;*/
  height: 36px;
}

/deep/.el-form-item__error {
  top: 35px;
  left: 10px;
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
  width: 800px;
  height: 360px;
  background: rgba(255, 255, 255, 0.6);
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
  color: #ffffff;
  position: relative;
  /* border: 1px solid black; */
  right: 150px;
  top: 100px;
}

.login-box {
  /*border: 1px solid red;*/
  width: 330px;
  height: 100%;
  position: absolute;
  top: 0;
  right: 0;
  border-radius: 0 30px 30px 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  /* border: 1px solid red; */
}

.login-box>h1 {
  font-size: 36px;
  margin: 40px 0;
  color: #1482F1;
}

/deep/.el-input__inner {
  /*border: 1px solid red!important;*/
  margin-bottom: 20px;
  width: 250px;
  height: 36px;
  padding: 0 20px;
  border-radius: 36px;
  border: none;
  outline: none;
  font-size: 18px;
}

.forget-btn {
 /* border: 1px solid red;*/
  color: #1482F1;
  position: relative;
  top: -10px;
  left: -80px;
  user-select: none;
  cursor: pointer;
}

.register-btn {
  color: #1482F1;
  position: relative;
  top: -31px;
  right: -80px;
  user-select: none;
  cursor: pointer;
}
.login-btn {
  margin-top: 0px;
  width: 150px;
  height: 40px;
  background-color: rgb(121, 119, 123);
  border: none;
  border-radius: 40px;
  color: #fff;
  font-size: 18px;
}

.login-btn:hover {
  background-color: #1482F1;
}

.tologin-btn:hover {
  background-color: #1482F1;
}

.reset-btn:hover {
  background-color: #1482F1;
}
/deep/.captcha .el-input__inner {
  /* border: 1px solid red!important;*/
  width: 150px!important;
  position: relative!important;
  left: -50px!important;
}

.yzm {
  border: none;
  border-radius: 20px;
  width: 80px;
  height: 38px;;
  background-color: #1482F1;
  color: #fff;
  position: relative;
  top: -58px;
  right: -80px;
}

.reset-btn {
  position: relative;
  top: -40px;
  width: 150px;
  height: 40px;
  background-color: rgb(146, 141, 149);
  border: none;
  border-radius: 40px;
  color: #fff;
  font-size: 18px;
}

.tologin-btn {
  position: relative;
  top: -30px;
  width: 150px;
  height: 40px;
  background-color: rgb(146, 141, 149);
  border: none;
  border-radius: 40px;
  color: #fff;
  font-size: 18px;
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
