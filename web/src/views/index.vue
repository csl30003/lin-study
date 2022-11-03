<template>
  <div>
    <el-main>
      <el-row class="el-row">
        <el-col align="center">
          <div class="radius">
            <el-form :model="form" label-width="auto" size="large">
              <el-form-item label="用户名">
                <el-input v-model="form.name"/>
              </el-form-item>
              <el-form-item label="密码">
                <el-input v-model="form.password" type="password"/>
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
.radius {
  height: 70%;
  width: 50%;
  border: 1px solid var(--el-border-color);
  border-radius: 20px;
  margin-top: 100px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.5);
}

.el-row {
  margin-top: 10px;
  margin-bottom: 10px;
}

.button-class {
  width: 70px;
  height: 50px;
  text-align: center;
  font-size: 30px;
}
</style>