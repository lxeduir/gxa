import axios from 'axios'
import { MessageBox, Message } from 'element-ui'
import nProgress from 'nprogress'
import 'nprogress/nprogress.css'
import { getToken } from './auth'
const service = axios.create({
  baseURL: 'https://api.edulx.xyz/front',
  timeout: 5000 // request timeout
})

service.interceptors.request.use(
  config => {
    const jwt=getToken()
    if (jwt) {
      config.headers['Authorization'] = jwt
    }
    nProgress.start()

    return config
  },
  error => {
    console.log(error)
    return Promise.reject(error)
  }
)

service.interceptors.response.use(
  response => {
    nProgress.done()
    const res = response.data
    if (response.status!==200) {
      Message({
        message: res.message || res.err||res.error||res.msg+'！请尽快进行身份认证！',
        type: 'error',
        duration: 5 * 1000
      })
      return Promise.reject(new Error(res.message || res.error))
    } else {
      if(res.err!=null||res.error!=null){
        Message({
          message: res.err||res.error||res.msg||"Error",
          type: 'error',
          duration: 5 * 1000
        })
        return Promise.reject(new Error('failure'))

      }
       else return response
    }
  },
  error => {
    nProgress.done()
    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)
export default service
