import request from "@/utils/request";

export const reqLogin=(data)=>{
    return request({url:'/login',method:'post',data})
}
export const reqRegister=(data)=>{
    // console.log(data)
    return request({url:'/signup',method:'post',data})
}
export const reqInfo=(uid,token)=>{
    return request({url:`/userinfo?uid=${uid}`,method:'get', headers: {'Authorization': token}})
}
//使用邮箱获取验证码
export const reqRegisterCode=(data)=>{
    console.log(data)
    return request({url:'/captcha',method:'post',data})
}