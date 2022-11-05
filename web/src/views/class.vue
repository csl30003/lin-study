<template>
  <el-image :src="require('@/assets/class.png')"
            style="margin: 0; position:absolute; height: 100%; width: 92%"></el-image>
  <div class="div">
    <el-container class="el-container">
      <el-header class="el-header">
        <h1 class="h1">
          {{ layer }}
        </h1>
      </el-header>

      <el-main class="el-main">
        <el-table class="el-table" size="large" :data="classList" :row-style="{height: '60px'}"
                  :cell-style="{padding:'0px'}">
          <el-table-column prop="i" min-width="20%"/>
          <el-table-column prop="name" min-width="40%"/>
          <el-table-column min-width="40%">
            <template #default="scope">
              <el-button
                  link
                  type="success"
                  size="large"
                  @click.prevent="getSeat(scope.row)"
                  style="font-size: 30px"
              >
                进入
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-main>
    </el-container>
  </div>
</template>

<script setup>
import {useRouter} from "vue-router"
import {onMounted, ref} from 'vue'
import instance from "@/axios"
import {ElMessage} from "element-plus"

const router = useRouter();

const floor = router.currentRoute.value.params.floor
const layer = router.currentRoute.value.params.layer

const classList = ref([])

onMounted(() => {
  getClass()
})

const getClass = async () => {
  instance.get('http://localhost:8080/index/' + floor + '/' + layer).then(res => {
    if (res.data.code === 200) {
      let temp
      temp = res.data.data
      for (let i = 0; i < temp.length; i++) {
        classList.value.push({
          'i': i + 1,
          name: temp[i]
        })
      }
    } else {
      ElMessage.error('无法获取楼层')
    }
  })
}

const getSeat = async (row) => {
  await router.push('/home/map/' + floor + '/' + layer + '/' + row.name)
}


</script>

<style scoped>
.div {
  height: 100%;
}

.el-container {
  height: 100%;
}

.el-header {
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

.el-table {
  width: 60%;
  margin: auto;
  font-size: 30px;
}

/*最外层透明*/
::v-deep .el-table,
::v-deep .el-table__expanded-cell {
  background-color: transparent !important;
}

/* 表格内背景颜色 */
::v-deep .el-table th,
::v-deep .el-table tr,
::v-deep .el-table td {
  background-color: transparent !important;
  color: white;
}

::v-deep .el-table .cell {
  line-height: 2vw;
}
</style>