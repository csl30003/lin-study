<template>
  <h2 class="h2">通知</h2>
  <el-table
      class="el-table"
      :data="noticeList"
      :row-style="{height: '60px'}"
      :cell-style="{padding:'0px'}"
      max-height="500px"
      :default-sort="{ prop: 'date', order: 'descending' }"
      stripe
  >
    <el-table-column prop="created_at" label="日期" min-width="15%" sortable/>
    <el-table-column prop="notice_man" label="通知人身份" min-width="15%"/>
    <el-table-column prop="notice_name" label="通知人名称" min-width="15%"/>
    <el-table-column prop="notice_type" label="通知类型" min-width="15%"/>
    <el-table-column prop="notice_content" label="通知内容" min-width="40%"/>
  </el-table>
</template>

<script setup>
import {onMounted, ref} from "vue";
import instance from "@/axios";
import {ElMessage} from "element-plus";

const noticeList = ref([])

onMounted(() => {
  getNotice()
})

const getNotice = async () => {
  instance.get('http://localhost:8080/index/notice').then(res => {
    if (res.data.code === 200) {
      for (const datum of res.data.data) {
        datum.created_at = datum.created_at.replace(/T/g, " ").substr(0, 19)

        let temp = datum.notice_content.split('-')
        datum.notice_man = temp[0]
        datum.notice_name = temp[1]
        datum.notice_type = temp[2]
        datum.notice_content = temp[3]
        noticeList.value.push(datum)
      }
    } else {
      ElMessage.error('无法获取日记')
    }
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
  width: 90%;
  margin: auto;
  font-size: 100%;
}
</style>