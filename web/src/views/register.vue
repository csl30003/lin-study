<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <el-row>
          <el-cow :span="12" :offset="6">
            <h1>
              添加学生
            </h1>
          </el-cow>
        </el-row>
      </el-header>

      <el-main>

        <el-row>
          <el-col :span="12" :offset="6">
            <el-form ref="formRef" :model="myForm" label-width="100px" class="demo-ruleForm">
              <el-form-item label="用户名" prop="name">
                <el-input v-model.string="myForm.name" type="text" autocomplete="off" />
              </el-form-item>

              <el-form-item>
                <el-button type="primary" @click="submitForm(formRef)">添加</el-button>
                <el-button @click="resetForm(formRef)">重置</el-button>
              </el-form-item>
            </el-form>
          </el-col>
        </el-row>

        <el-row>
          <el-col :span="12" :offset="6">
            <el-table :data="tableData" style="width: 100%">
              <el-table-column prop="name" lable="用户名" width="180" />
            </el-table>
          </el-col>
        </el-row>

      </el-main>
    </el-container>
  </div>
</template>



<script lang="ts" setup>

import { reactive, ref, onMounted } from 'vue'
import type { FormInstance } from 'element-plus'
import instance from '../axios/index'

let tableData = ref([])

const addStudents = async () => {
  instance.post('http://localhost:8080/student',{
    name: myForm.name,
  }).then(res => {
    console.log(res.data)
    getStudents()
  })
}

const getStudents = async () => {
  instance.get('http://localhost:8080/student').then(res => {
    tableData.value = res.data.data
    console.log(tableData)
  })
}

onMounted(() => {
  getStudents()
})

const formRef = ref<FormInstance>()

const myForm = reactive({
  name: '',
})

const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  addStudents()
}

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
}

</script>



<style scoped>
h1{
  text-align: center;
}
</style>