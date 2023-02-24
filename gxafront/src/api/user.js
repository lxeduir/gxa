import request from "@/utils/request";

export const reqLogin=(data)=>{
    return request({url:'/login',method:'post',data})
}
export const reqRegister=(data)=>{
    return request({url:'/signup',method:'post',data})
}
export const reqInfo=(uid,token)=>{
    return request({url:`/userinfo?uid=${uid}`,method:'get', headers: {'Authorization': token}})
}
//使用邮箱获取验证码
export const reqRegisterCode=(data)=>{
    return request({url:'/captcha',method:'post',data})
}

export const reqIdentity=()=>{
    return request({url:`/identity`,method:'get'})
}
export const upIdentity=(data)=>{
    return request({url:'/identity',method:'post',data})
}
export const updateIdentity=(data)=>{
    return request({url:'identity',method:'put',data})
}
export const reqForgetPasswd=(data)=>{
    return request({url:'/forgetpasswd',method:'post',data})
}

export const reqPublic = (data) => {
    return request({url:'/notice',method:'get'})
}