<template>
  <div class="common-layout">
    <el-container height="100%">
      <el-header height="300px">
        <h1 class="h1-class-2">登录</h1>
      </el-header>

      <el-main>
        <el-row>
          <el-col :span="8" :offset="8">
            <el-form :model="form" label-width="100px" size="large">
              <el-form-item label="用户名">
                <el-input v-model="form.name"/>
              </el-form-item>
              <el-form-item label="密码">
                <el-input v-model="form.password" type="password"/>
              </el-form-item>

              <el-form-item>
                <el-button type="primary" @click="login">登录</el-button>
                <el-button>返回</el-button>
              </el-form-item>
            </el-form>
          </el-col>
        </el-row>

      </el-main>
    </el-container>
  </div>
</template>

<script lang="ts">
import {reactive} from 'vue'
import instance from '../axios/index'
import {defineComponent} from 'vue'

export default defineComponent({
  methods: {
    home() {
      this.$router.push({
        path: '/home'
      });
    }
  }
})

const form = reactive({
  name: '',
  password: '',
})

const login = async () => {
  instance.post('http://localhost:8080/login', {
    name: form.name,
    password: form.password,
  }).then(res => {
    if (res.status) {

    }
  })
}
</script>

<style>
.h1-class-2 {
  text-align: center;
  font-size: 80px;
  color: #ffe05d;
}
</style>