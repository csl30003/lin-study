const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  devServer: {
    port: 8081
  },
  transpileDependencies: true,
  lintOnSave: false
})
