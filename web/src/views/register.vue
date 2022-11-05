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
              <el-form-item class="el-form-item">
                <el-input class="el-input" maxlength="30" show-word-limit placeholder="用户名" v-model="form.name"/>
              </el-form-item>

              <el-form-item class="el-form-item">
                <el-input class="el-input" maxlength="30" show-word-limit placeholder="密码" show-password v-model="form.password" type="password"/>
              </el-form-item>

              <el-form-item class="el-form-item">
                <el-input class="el-input" maxlength="30" show-word-limit placeholder="标签" v-model="form.label"/>
              </el-form-item>

              <el-form-item class="el-form-item">
                <el-input class="el-input" maxlength="10" show-word-limit placeholder="地区" v-model="form.area"/>
              </el-form-item>

              <el-form-item class="el-form-item">
                <el-input class="el-input" maxlength="10" show-word-limit placeholder="阶段" v-model="form.stage"/>
              </el-form-item>

              <el-form-item class="el-form-item">
                <el-input class="el-input" maxlength="20" show-word-limit placeholder="学校" v-model="form.school"/>
              </el-form-item>

              <el-form-item class="el-form-item">
                <el-input class="el-input" maxlength="30" show-word-limit placeholder="目标" v-model="form.goal"/>
              </el-form-item>

              <el-form-item class="el-form-item">
                <el-switch
                    v-model="form.sex"
                    style="--el-switch-on-color: #467dc7; --el-switch-off-color: #83e112"
                    active-value=0
                    inactive-value=1
                    active-text="男"
                    inactive-text="女"
                />
              </el-form-item>

              <el-row>
                <el-col :span="6" :offset="6">
                  <el-button type="primary" class="button-class" round @click="commit">提交</el-button>
                </el-col>
                <el-col :span="6">
                  <el-button type="success" class="button-class" round @click="back">返回</el-button>
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
import instance from "@/axios";
import { ElMessage } from 'element-plus'

const form = reactive({
  name: '',
  password: '',
  label: '',
  sex: '',
  area: '',
  stage: '',
  school: '',
  goal: '',
})

const router = useRouter();

const commit = async () => {
  instance.post('http://localhost:8080/register', {
    name: form.name,
    password: form.password,
    label: form.label,
    sex: parseInt(form.sex),
    area: form.area,
    stage: form.stage,
    school: form.school,
    goal: form.goal,
    'image': '/default.png',
  }).then(res => {
    if (res.data.code === 200) {
      ElMessage({
        message: res.data.message,
        type: 'success',
      })
      router.push('/');
    } else {
      ElMessage.error(res.data.message)
    }
  })
}
const back = () => {
  router.push('/');
}


</script>

<style scoped>
.h1 {
  text-align: center;
  font-size: 100px;
  transform: translateY(-50%);
  color: #ffbc37;
  font-family: STLiti, serif;
}

.radius {
  height: 85%;
  width: 400px;
  border: 1px solid var(--el-border-color);
  border-radius: 20px;
  margin-top: 100px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.5);
  text-align: center;
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

.el-form-item{
  display: inline-block;
  width: 80%;
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