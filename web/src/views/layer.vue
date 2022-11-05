<template>
  <div class="div">
    <el-container class="el-container">
      <el-header class="el-header">
        <h1 class="h1">
          {{floor}}
        </h1>
      </el-header>

      <el-main class="el-main">
        Main
      </el-main>
    </el-container>
  </div>
</template>

<script setup>
import {useRouter} from "vue-router"
import {ref, onMounted} from 'vue'
import instance from "@/axios"
import {ElMessage} from "element-plus"

const router = useRouter();

const floor = router.currentRoute.value.params.floor

const layerList = ref([])

onMounted(() => {
  getLayer()
})

const getLayer = async () => {
  instance.get('http://localhost:8080/index/' + floor).then(res => {
    if (res.data.code === 200) {
      layerList.value = res.data.data
    } else {
      ElMessage.error('无法获取楼层')
    }
  })
}
</script>

<style scoped>
.div{
  height: 100%;
}

.el-container{
  height: 100%;
}

.el-header{
  height: 30%;
}

.h1 {
  text-align: center;
  font-size: 100px;
  transform: translateY(-50%);
  color: #ffbc37;
  font-family: STLiti, serif;
}

.el-main {
  height: 70%;
  width: 100%;
  margin: 0;
  padding: 0;
}
</style>