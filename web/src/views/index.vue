<template>
  <div class="common-layout">
    <el-container height="100%">
      <el-header height="300px">
        <h1 class="h1-class-1">LinStudy</h1>
      </el-header>

      <el-main>
        <el-row class="el-row">
          <el-col align="center">
            <el-button type="primary" class="button-class" round @click="login()">登录</el-button>
          </el-col>
        </el-row>
        <el-row class="el-row">
          <el-col align="center">
            <el-button type="success" class="button-class" round @click="register()">注册</el-button>
          </el-col>
        </el-row>
      </el-main>
    </el-container>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue'

export default defineComponent({
  methods: {
    login() {
      this.$router.push({
        path: '/login'
      });
    },
    register() {
      this.$router.push({
        path: '/register'
      });
    }
  }
})
</script>

<style scoped>
.h1-class-1 {
  text-align: center;
  font-size: 120px;
  color: #ffe05d;
}

.el-row {
  margin-bottom: 50px;
}

.button-class {
  width: 300px;
  height: 150px;
  text-align: center;
  font-size: 70px;
}
</style>