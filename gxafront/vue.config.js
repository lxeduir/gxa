const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: false,    //关闭eslint校验工具
  devServer: {
    host: 'localhost',
    port: 8080,    //端口号
    open: true
  }
})

