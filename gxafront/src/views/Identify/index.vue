<template>
    <div id="Identity">
        <div class="picture">
            <a>身份认证</a>
        </div>
        <div class="backgroud">
            <div class="message">
                <div>
                    <pre>姓    名</pre>
                    <el-input class="name" placeholder="姓名" v-model="name"></el-input>
                </div>
                <div>
                    <pre>性    别</pre>
                    <div class="sex">
                        <el-radio v-model="sex" label="1">男</el-radio>
                        <el-radio v-model="sex" label="2">女</el-radio>
                    </div>
                </div>
                <div>
                    <pre>民    族</pre>
                    <el-input class="name" placeholder="民族" v-model="ethnic"></el-input>
                </div>
                <div>
                    <pre>籍    贯</pre>
                    <el-input class="name" placeholder="籍贯" v-model="origin"></el-input>
                </div>
                <div>
                    <pre>政治面貌</pre>
                    <div class="polstatus">
                        <el-radio v-model="polstatus" label="党员">党员</el-radio>
                        <el-radio v-model="polstatus" label="团员">团员</el-radio>
                        <el-radio v-model="polstatus" label="其他">其他</el-radio>
                    </div>
                </div>
                <div>
                    <pre>证    件</pre>
                    <el-select class="name" v-model="value" placeholder="请选择">
                        <el-option v-for="item in docutypes" :key="item.value" :label="item.label" :value="item.label">
                        </el-option>
                    </el-select>
                </div>
                <div>
                    <pre>证 件 号</pre>
                    <el-input class="name" placeholder="证件号" v-model="docunum"></el-input>
                </div>
                <div>
                    <pre>生    日</pre>
                    <el-date-picker class="name" v-model="birthdate" type="date" placeholder="选择日期">
                    </el-date-picker>
                </div>
                <div>
                    <pre>手    机</pre>
                    <el-input class="name" placeholder="手机号" v-model="phonenum"></el-input>
                </div>
            </div>
            <button class="submit" @click="sendInfo">提交</button>
        </div>
    </div>
</template>

<script>
import { Message } from 'element-ui';
export default {
    data() {
        return {
            name: '',
            sex: '0',
            ethnic: '',
            polstatus: '',
            origin: '',
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
            value: '',
            docunum: '',
            birthdate: '',
            phonenum: ''
        };
    },
    methods: {
        async sendInfo() {
            try {
                const { name, sex, ethnic, polstatus, origin, value, docunum, birthdate, phonenum } = this
                const info = {
                    name, sex: Number(sex), ethnic, polstatus, origin, value, docunum, birthdate, phonenum
                }
                if (name && sex && ethnic && polstatus && origin && value && docunum && birthdate && phonenum) {
                    await this.$store.dispatch('user/uploadIdentityInfo', info)
                    setTimeout(() => {
                        Message({
                            message: '上传成功！',
                            type: 'success',
                            duration: 3 * 1000
                        }),
                            this.$router.push('/account')
                            this.$router.go(0)
                    }, 1000);
                } else {
                    Message({
                        message: '请检查信息是否填完整！',
                        type: 'error',
                        duration: 3 * 1000
                    })
                }
            }
            catch (error) {
                console.log(error.message)
            }
        }
    },
};
</script>

<style scoped>
.picture {
    height: 300px;
    width: 1540px;
    background-color: #6db6ff;

}

.picture>a {
    position: absolute;
    top: 200px;
    left: 100px;
    font-size: 30px;
    color: white;

}

.backgroud {
    border: 1px solid rgb(190, 189, 189);
    position: relative;
    top: -120px;
    left: 170px;
    width: 1200px;
    height: 1000px;
    background-color: white;
    box-shadow: 3px 3px 3px 3px #c7c7c7;
}

.message {
    position: relative;
    left: 200px;
    width: 800px;

}

.message pre {
    /*border: 1px solid red;*/
    line-height: 60px!important;
    font-size: 20px;
    font-style: normal;
}

.message div {
    /*border: 1px solid green;*/
    display: flex;
    margin-top: 20px;
}

.name {
    position: relative;
    top: -10px;
    left: 100px;
    width: 500px;
}

.sex {
    position: relative;
    top: 3px;
    left: 100px;
}

.polstatus {
    position: relative;
    top: 3px;
    left: 100px;
}

.submit {
    position: absolute;
    top: 800px;
    left: 900px;
    width: 100px;
    height: 40px;
    line-height: 40px;
    background-color: #6db6ff;
    color: white;
    border-radius: 5px;
    border: none;
}
</style>