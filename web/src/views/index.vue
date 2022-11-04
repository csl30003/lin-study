<template>
  <div>
    <el-main class="el-main">
      <el-image :src="require('@/assets/back.png')"
                style="margin: 0;position:absolute; height: 100%; width: 100%"></el-image>
      <el-row>
        <el-col align="center">
          <h1 class="h1">LinStudy</h1>
        </el-col>
      </el-row>
      <el-row class="el-row">
        <el-col align="center">
          <div class="radius">
            <el-form :model="form" label-width="auto" size="large">
              <br>
              <br>
              <el-form-item style="display: inline-block; width: 80%">
                <el-input class="el-input" maxlength="30" show-word-limit placeholder="用户名" v-model="form.name"/>
              </el-form-item>
              <el-form-item style="display: inline-block; width: 80%">
                <el-input class="el-input" maxlength="30" show-word-limit placeholder="密码" show-password
                          v-model="form.password" type="password"/>
              </el-form-item>

              <el-row>
                <el-col :span="6" :offset="6">
                  <el-button type="primary" class="button-class" round @click="login">登录</el-button>
                </el-col>
                <el-col :span="6">
                  <el-button type="success" class="button-class" round @click="register">注册</el-button>
                </el-col>
              </el-row>
            </el-form>
          </div>
        </el-col>
      </el-row>
    </el-main>
  </div>
</template>

<script setup>
import {useRouter} from "vue-router";
import {reactive} from "vue";

const router = useRouter();

const login = () => {
  router.push('/login');
}
const register = () => {
  router.push('/register');
}
const form = reactive({
  name: '',
  password: '',
})

</script>

<style scoped>
.h1 {
  text-align: center;
  font-size: 100px;
  transform: translateY(-50%);
  color: #ffbc37;
  font-family: STLiti;
}

.radius {
  height: 70%;
  width: 400px;
  border: 1px solid var(--el-border-color);
  border-radius: 20px;
  margin-top: 100px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.5);
  text-align:center;
}

.el-main {
  height: 100%;
  width: 100%;
  margin: 0;
  padding: 0;
}

.el-row {
  margin-top: 10px;
  margin-bottom: 10px;
}

.el-input {
  text-align: center;
  vertical-align: middle;
}

.button-class {
  width: 70px;
  height: 50px;
  text-align: center;
  font-size: 30px;
}
</style>