import { reqInfo, reqLogin, reqRegister, reqRegisterCode } from '@/api/user'
import { getToken, setToken, removeToken,setUser,getUser,removeUser} from '@/utils/auth'
import { resetRouter } from '@/router'
import router from '@/router'
import { Message } from 'element-ui'
const getDefaultState = () => {
    return {
        Uid: '',
        Info: [],
        token: null,
        username:''||getUser()
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
    SET_Username:(state,username)=>{
        state.username=username
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
            setUser(res.username)
            commit('SET_TOKEN', res.token)
            commit('SET_Username',res.username)
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
        let result = await reqInfo(state.token)
        const res = result.data
        if (result.status == 200) {
            commit('SET_INFO', res.list)
        } else {
            return Promise.reject(new Error('falie'))
        }
    },
    async getyzm({commit},info) {
        const {email,type}=info
        let result = await reqRegisterCode(info)
        Message({
            message: "获取成功!",
            type: 'success',
            duration: 3 * 1000
        })
        if (result.status != 200) return Promise.reject(new Error('falie'))

    },
    logout({ commit }) {
        removeUser()
        removeToken()
        commit('RESET_STATE')
    },
    resetToken({ commit }) {
        return new Promise(resolve => {
            removeToken()
            commit('RESET_STATE')
            resolve()
        })
    }
}
export default {
    namespaced: true,
    state,
    mutations,
    actions,
}

