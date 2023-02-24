<template>
    <div class="container">
        <div class="account">
            <el-descriptions class="margin-top" title="账号信息" :column="4" border="true">
                <el-descriptions-item label="用户名">{{ accountInfo.username }}</el-descriptions-item>
                <el-descriptions-item label="邮箱">{{ accountInfo.email }}</el-descriptions-item>
                <el-descriptions-item label="密码" :span="3">************
                    <el-button @click="dialogPasswdVisible = true" type="info" size="mini">修改</el-button>
                </el-descriptions-item>
                <el-descriptions-item label="注册时间" :span="2">{{ accountInfo.signtime }}</el-descriptions-item>
            </el-descriptions>
            <br><br><br><br><br><br>
            <el-descriptions class="margin-top" title="身份信息" :column="4" border="true">
                <el-descriptions-item label="姓名">{{ identityInfo.Name }}</el-descriptions-item>
                <el-descriptions-item label="性别" v-if="this.identityInfo.Sex == 1">男</el-descriptions-item>
                <el-descriptions-item label="性别" v-if="this.identityInfo.Sex == 2">女</el-descriptions-item>
                <el-descriptions-item label="民族">{{ identityInfo.Ethnic }}</el-descriptions-item>
                <el-descriptions-item label="籍贯">{{ identityInfo.Origin }}</el-descriptions-item>
                <el-descriptions-item label="政治面貌">{{ identityInfo.Polstatus }}</el-descriptions-item>
                <el-descriptions-item label="证件">{{ identityInfo.Value }}</el-descriptions-item>
                <el-descriptions-item label="证件号" :span="2">{{ identityInfo.Docunum }}</el-descriptions-item>
                <el-descriptions-item label="生日" :span="2">{{ identityInfo.Birthdate }}</el-descriptions-item>
                <el-descriptions-item label="手机号" :span="2">{{ identityInfo.Phonenum }}</el-descriptions-item>

                <el-descriptions-item label="备注">
                    <!-- <el-tag size="small">学校</el-tag> -->
                </el-descriptions-item>


            </el-descriptions>
            <el-button type="info" class="updatebtn" icon="el-icon-edit" @click="handleEidt()">修改</el-button>
            <el-button type="info" class="upbtn" icon="el-icon-upload el-icon--right"
                @click="toIdentity()">上传</el-button>
            <!-- 弹出框 -->
            <el-dialog title="更改信息" :visible.sync="dialogFormVisible">
                <el-form :model="form">
                    <el-form-item label="姓名" :label-width="formLabelWidth">
                        <el-input :placeholder="form.Name" v-model="form.Name" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="性别" :label-width="formLabelWidth">
                        <el-radio v-model="form.Sex" label="1">男</el-radio>
                        <el-radio v-model="form.Sex" label="2">女</el-radio>
                    </el-form-item>
                    <el-form-item label="民族" :label-width="formLabelWidth">
                        <el-input :placeholder="form.Ethnic" v-model="form.Ethnic" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="籍贯" :label-width="formLabelWidth">
                        <el-input :placeholder="form.Origin" v-model="form.Origin" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="政治面貌" :label-width="formLabelWidth">
                        <el-radio v-model="form.Polstatus" label="党员">党员</el-radio>
                        <el-radio v-model="form.Polstatus" label="团员">团员</el-radio>
                        <el-radio v-model="form.Polstatus" label="其他">其他</el-radio>
                    </el-form-item>
                    <el-form-item label="证件" :label-width="formLabelWidth">
                        <el-select v-model="form.Value" placeholder="请选择">
                            <el-option v-for="item in docutypes" :key="item.value" :label="item.label"
                                :value="item.label">
                            </el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="证件号" :label-width="formLabelWidth">
                        <el-input :placeholder="form.Docunum" v-model="form.Docunum" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="生日" :label-width="formLabelWidth">
                        <el-date-picker :placeholder="选择日期" v-model="form.Birthdate" type="date">
                        </el-date-picker> </el-form-item>
                    <el-form-item label="手机" :label-width="formLabelWidth">

                        <el-input :placeholder="form.Phonenum" v-model="form.Phonenum" autocomplete="off"></el-input>

                    </el-form-item>
                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="dialogFormVisible = false">取 消</el-button>
                    <el-button type="primary" @click="updateIdentity()">确 定</el-button>
                </div>
            </el-dialog>
            <el-dialog title="更改密码" :visible.sync="dialogPasswdVisible">
                <el-form :model="loginForm">
                    <el-form-item label="邮箱" :label-width="formLabelWidth">
                        <el-input :placeholder="loginForm.email" v-model="loginForm.email"
                            autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="新密码" :label-width="formLabelWidth">
                        <el-input type="password" v-model="loginForm.passwd" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="验证码" :label-width="formLabelWidth">
                        <el-input v-model="loginForm.captcha" autocomplete="off"></el-input>
                        <button class="yzm" @click="getyzm">获取验证码</button>
                    </el-form-item>

                </el-form>
                <div slot="footer" class="dialog-footer">
                    <el-button @click="dialogFormVisible = false">取 消</el-button>
                    <el-button type="primary" @click="resetPasswd()">确 定</el-button>
                </div>
            </el-dialog>
        </div>
    </div>
</template>
<script>
import { reqInfo, reqForgetPasswd } from '@/api/user'
import store from '@/store'
import { getUid } from '@/utils/auth'
export default {
    name: 'index',
    data() {
        return {
            loginForm: {
                email: '',
                passwd: '',
                captcha: '',
            },
            identityInfo: {},
            accountInfo: {
                username: '',
                email: '',
                signtime: ''
            },
            lpersonalInfo: {
            },
            dialogPasswdVisible: false,
            dialogFormVisible: false,
            formLabelWidth: '120px',
            form:
            {
                // Name: '',
                Sex: '0',
                // Ethnic: '',
                // Origin: '',
                // Polstatus: '',
                // Value: '',
                // Docunum: '',
                // Birthdate: '',
                // Phonenum: ''
            },
            docutypes: [{
                value: '选项1',
                label: '居民身份证'
            }, {
                value: '选项2',
                label: '护照'
            }, {
                value: '选项3',
                label: '士官证'
            }, {
                value: '选项4',
                label: '港澳通行证'
            }],
        }

    },
    created() {
        this.$nextTick(() => {
            this.accountInfo = store.getters.info,
            this.getidentity()
        })

    },
    methods: {
        // async getinfo() {
        //     try {
        //         const res = await reqInfo(getUid())
        //         this.accountInfo = {
        //             username: res.data.username,
        //             email: res.data.email,
        //             signtime: res.data.signtime

        //         }
        //     }
        //     catch (error) {
        //         console.log(error.message)
        //     }
        // },
        async getidentity() {
            try {
                const res = await this.$store.dispatch('user/reqIdentityInfo')
                res[0].Sex = String(res[0].Sex)
                this.identityInfo = res[0]
                this.form = res[0]
            }
            catch (error) {
                console.log(error.message)
            }

        },
      
        toIdentity() {
            this.$router.push({ name: 'identify' })
        },
        handleEidt() {
            this.dialogFormVisible = true
        },
        async updateIdentity() {
            try {
                const sex = this.form.Sex

                this.form.Sex = Number(sex)
                console.log(this.form)
                await this.$store.dispatch('user/updateIdentityInfo', this.form)
                this.dialogFormVisible = false
            }
            catch (error) {
                console.log(error.message)
            }
            window.location.reload()
        },
        async resetPasswd() {
            try {
                this.dialogPasswdVisible = false
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
            } catch (error) {
                console.log(error.message)
            }

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
    }
}
</script>
<style>
.container {
    height: 700px;
    position: relative;
}

.upbtn {
    position: absolute;
    top: 500px;
    left: 500px;
}

.yzm {
    position: absolute;
    top: 10px;
    right: 60px;
}

.updatebtn {
    position: absolute;
    top: 500px;
    left: 400px;
}

.account {
    /* border: 1px solid red; */

}
</style>