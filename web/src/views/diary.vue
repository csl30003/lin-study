<template>
  <h2 class="h2">日记簿</h2>
  <el-table
      class="el-table"
      :data="diaryList"
      :row-style="{height: '60px'}"
      :cell-style="{padding:'0px'}"
      max-height="500px"
      :default-sort="{ prop: 'date', order: 'descending' }"
      stripe
  >
    <el-table-column prop="created_at" label="日期" min-width="20%" sortable/>
    <el-table-column prop="diary_content" label="日记内容" min-width="50%"/>
    <el-table-column prop="likes" label="点赞数" min-width="10%"/>
    <el-table-column label="操作" min-width="20%">
      <template #default="scope">
        <el-button type="success" round size="default" @click.prevent="likeDiary(scope.row)">点赞</el-button>
        <el-button type="warning" round size="default" @click.prevent="changeDiary(scope.row)">修改</el-button>
        <el-button type="danger" round size="default" @click.prevent="eraseTheDiary(scope.row)">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
  <el-button class="mt-4" style="width: 100%; height: 5%;" @click="keepDiary">写日记</el-button>
</template>

<script setup>
import {onMounted, ref} from "vue";
import instance from "@/axios";
import {ElMessage, ElMessageBox} from "element-plus";

let diaryList = ref([])

onMounted(() => {
  getDiary()
})

const getDiary = async () => {
  instance.get('http://localhost:8080/index/diary').then(res => {
    if (res.data.code === 200) {
      for (const datum of res.data.data) {
        datum.created_at = datum.created_at.replace(/T/g, " ").substr(0, 19)
        diaryList.value.push(datum)
      }
    } else {
      ElMessage.error('无法获取日记')
    }
  })
}

const keepDiary = async () => {
  ElMessageBox.prompt('输入你的想法', '写日记', {
    confirmButtonText: '提交',
    cancelButtonText: '取消',
    center: true,
    type: 'info',
  })
      .then(({value}) => {
        instance.post('http://localhost:8080/index/diary', {
          diary_content: value
        }).then(res => {
          if (res.data.code === 200) {
            let temp = res.data.data
            temp.created_at = temp.created_at.replace(/T/g, " ").substr(0, 19)
            diaryList.value.unshift(temp)

            ElMessage({
              message: res.data.message,
              type: 'success',
            })
          } else {
            ElMessage.error(res.data.message)
          }
        })
      })
      .catch(() => {
      })
}

const likeDiary = async (row) => {
  instance.patch('http://localhost:8080/index/likeDiary', {
    id: row.id
  }).then(res => {
    if (res.data.code === 200) {
      row.likes += 1

      ElMessage({
        message: res.data.message,
        type: 'success',
      })
    } else {
      ElMessage.error(res.data.message)
    }
  })
}

const changeDiary = async (row) => {
  ElMessageBox.prompt('重新组织语言', '修改日记', {
    confirmButtonText: '提交',
    cancelButtonText: '取消',
    center: true,
    inputValue: row.diary_content,
    type: 'warning',
  })
      .then(({value}) => {
        instance.patch('http://localhost:8080/index/diary', {
          id: row.id,
          diary_content: value,
        }).then(res => {
          if (res.data.code === 200) {
            for (const datum of diaryList.value) {
              if (datum.id === row.id) {
                datum.diary_content = value
                break
              }
            }

            ElMessage({
              message: res.data.message,
              type: 'success',
            })
          } else {
            ElMessage.error(res.data.message)
          }
        })
      })
      .catch(() => {
      })
}

const eraseTheDiary = async (row) => {
  ElMessageBox.confirm(
      '确定要删除这篇日记吗?',
      '警告',
      {
        confirmButtonText: '删除',
        cancelButtonText: '取消',
        type: 'error',
        center: true,
      }
  )
      .then(() => {
        instance.delete('http://localhost:8080/index/diary', {
          data: {
            id: row.id
          }
        }).then(res => {
          if (res.data.code === 200) {
            ElMessage({
              message: res.data.message,
              type: 'success',
            })
          } else {
            ElMessage.error(res.data.message)
          }
        })
      })
      .catch(() => {
      })
}

</script>

<style scoped>
.h2 {
  text-align: center;
  font-size: 70px;
  transform: translateY(-50%);
  color: #ffbc37;
  font-family: STLiti, serif;
}

.el-table {
  width: 80%;
  margin: auto;
  font-size: 100%;
}

.el-table .warning-row {
  --el-table-tr-bg-color: var(--el-color-warning-light-9);
}

.el-table .success-row {
  --el-table-tr-bg-color: var(--el-color-success-light-9);
}
</style>