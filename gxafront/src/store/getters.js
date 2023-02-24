const getters = {
    uid:  state=>state.user.uid,
    token: state => state.user.token,
    username:state => state.user.username,
    email:state => state.user.Info.email,
    signtime:state=>state.user.Info.signtime,
    identity:state=>state.user.identity,
    info:state=>state.user.Info,
    notice:state=>state.user.notice 
  }
  export default getters
  