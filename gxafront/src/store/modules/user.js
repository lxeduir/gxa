import { reqInfo, reqLogin, reqRegister, reqRegisterCode, upIdentity, reqIdentity, updateIdentity,reqPublic } from '@/api/user'
import { getUid, getToken, setToken, removeToken, setUser, getUser, removeUser, setUid, removeUid } from '@/utils/auth'
import { resetRouter } from '@/router'
import router from '@/router'
import { Message } from 'element-ui'
const getDefaultState = () => {
    return {
        Uid: '',
        Info: {},
        token: null,
        username: null,
        identity: null,
        notice: null
    }
}
const state = getDefaultState()
const mutations = {
    RESET_STATE: (state) => {
        Object.assign(state, getDefaultState())
    },
    SET_TOKEN: (state, token) => {
        state.token = token
    },
    SET_UID: (state, Uid) => {
        state.Uid = Uid
    },
    SET_INFO: (state, Info) => {
        state.Info = Info
    },
    SET_Username: (state, username) => {
        state.username = username
    },
    SET_identity: (state, identity) => {
        state.identity = identity
    },
    SET_notice: (state, notice) => {
        state.notice = notice
    }
}
const actions = {
    //user register
    async register({ commit }, registerFrom) {
        console.log(registerFrom)
        const { username, passwd, email, captcha } = registerFrom
        let result = await reqRegister({ username: username.trim(), email: email.trim(), passwd: passwd.trim(), captcha: captcha.trim() })
        console.log(result)
        if (result.status === 200) {
            Message({
                message: "注册成功!",
                type: 'success',
                duration: 3 * 1000
            })
        } else {
            return Promise.reject(new Error('faile'))
        }
    },
    // user login
    async login({ commit }, userInfo) {
        const { email, passwd } = userInfo
        let result = await reqLogin({ email: email.trim(), passwd: passwd })
        const res = result.data
        if (result.status === 200) {
            setToken(res.token)
            setUid(res.uid)
            commit('SET_TOKEN', res.token)
            commit('SET_UID', res.uid)

            Message({
                message: "登录成功!",
                type: 'success',
                duration: 3 * 1000
            })
        } else {
            return Promise.reject(new Error('faile'))
        }
    },
    async getInfo({ commit, state }) {
        let result = await reqInfo(getUid())
        const res = result.data
        if (result.status == 200) {
            setUser(res.username)
            commit('SET_Username', res.username)
            commit('SET_INFO', res)
        } else {
            return Promise.reject(new Error('falie'))
        }
    },
    async getyzm({ commit }, info) {
        let result = await reqRegisterCode(info)
        if (result.status == 200)
            Message({
                message: "验证码已发送请注意查收！!",
                type: 'success',
                duration: 3 * 1000
            })
        else return Promise.reject(new Error('falie'))

    },
    logout({ commit }) {
        removeUser()
        removeUid()
        removeToken()
        commit('RESET_STATE')
        router.push('/')
     
    },
    resetToken({ commit }) {
        return new Promise(resolve => {
            removeToken()
            commit('RESET_STATE')
            resolve()
        })
    },
    async uploadIdentityInfo({ commit }, info) {
        let result = await upIdentity(JSON.stringify(info))
        if (result.status == 200) {
            commit('SET_identity', info)
        } else {
            return Promise.reject(new Error('falie'))
        }
    },
    async reqIdentityInfo({ commit }) {
        let result = await reqIdentity()
        const res = result.data
        if (result.status == 200) {
            commit('SET_identity', res.list)
            return res.list
        } else {
            return Promise.reject(new Error('falie'))
        }
    },
    async updateIdentityInfo({ commit }, from) {
        let result = await updateIdentity(JSON.stringify(from))
        if (result.status == 200) {
            Message({
                message: "修改成功!",
                type: 'success',
                duration: 3 * 1000
            })
            return "ok"
        } else {
            return Promise.reject(new Error('falie'))
        }
    },
    async reqPublicInfo({ commit }) {
        let result = await reqPublic()
        const res = result.data
        if (result.status == 200) {
            commit('SET_notice', res.list)
            return res.list
            console.log("res.list")
        } else {
            return Promise.reject(new Error('falie'))
        }
    }
}
export default {
    namespaced: true,
    state,
    mutations,
    actions,
}

