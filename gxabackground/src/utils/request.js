import axios from 'axios'
import { MessageBox, Message } from 'element-ui'
import { getToken } from '@/utils/auth'
const service = axios.create({
  baseURL: 'https://api.edulx.xyz/backstage',
  timeout: 5000 // request timeout
})

service.interceptors.request.use(
  config => {
    const jwt=getToken()
    if (jwt) {
      config.headers['Authorization'] = jwt
    }
    return config
  },
  error => {
    console.log(error)
    return Promise.reject(error)
  }
)

service.interceptors.response.use(
  response => {
    // console.log(response)
    const res = response.data
    if (response.status!=200) {
      Message({
        message: res.message ||res.msg|| res.err||'Error',
        type: 'error',
        duration: 5 * 1000
      })
      return Promise.reject(new Error(res.message || 'Error'))
    } else {
      // console.log(res.err)
      if(res.err!=null||res.error){
        Message({
          message: res.err,
          type: 'error',
          duration: 5 * 1000
        })
      }
       else return response
    }
  },
  error => {
    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)
export default service
