<template>
  <div>
    <el-container class="el-container">
      <el-aside width="8%">
        <el-radio-group v-model="isCollapse" style="margin-bottom: 20px">
          <el-radio-button :label="false">详细</el-radio-button>
          <el-radio-button :label="true">缩小</el-radio-button>
        </el-radio-group>

        <el-menu
            class="el-menu-vertical"
            :collapse="isCollapse"
        >
          <el-menu-item index="1" @click="floor">
            <el-icon><School /></el-icon>
            <template #title>地图</template>
          </el-menu-item>

          <el-menu-item index="2" @click="friend">
            <el-icon><Sugar /></el-icon>
            <template #title>同学</template>
          </el-menu-item>

          <el-menu-item index="3" @click="chat">
            <el-icon><ChatLineSquare /></el-icon>
            <template #title>纸条</template>
          </el-menu-item>

          <el-menu-item index="4" @click="diary">
            <el-icon><Notebook /></el-icon>
            <template #title>日记</template>
          </el-menu-item>

          <el-menu-item index="5" @click="notice">
            <el-icon><BellFilled /></el-icon>
            <template #title>通知</template>
          </el-menu-item>

          <el-menu-item index="6">
            <el-icon><User /></el-icon>
            <template #title>个人</template>
          </el-menu-item>

          <el-menu-item index="7" @click="logout">
            <el-icon><CloseBold /></el-icon>
            <template #title>退出登录</template>
          </el-menu-item>
        </el-menu>
      </el-aside>

      <el-main class="el-main">
        <router-view style="height: 100%; width: 100%"></router-view>
      </el-main>
    </el-container>
  </div>
</template>

<script setup>
import {useRouter} from "vue-router"
import {ref, onMounted, reactive} from 'vue'
import {ElMessage} from "element-plus"
import instance from "@/axios"
import {
  School,
  Sugar,
  ChatLineSquare,
  Notebook,
  BellFilled,
  User,
  CloseBold,
} from '@element-plus/icons-vue'

const router = useRouter();

const isCollapse = ref(false)

const student = reactive({
  id: '',
  name: '',
})

onMounted(() => {
  getStudentInfo()
})

const getStudentInfo = async ()  => {
  instance.get('http://localhost:8080/index/getStudentInfo').then(res => {
    if (res.data.code === 200) {
      student.id = res.data.data.id
      student.name = res.data.data.name
      ElMessage({
        message: '欢迎' + res.data.data.name + '同学😘',
        type: 'success',
      })
    } else {
      ElMessage.error('请这位同学先报道哦😘')
      router.push('/')
    }
  })
}

const floor = async () => {
  await router.push('/home/map')
}

const friend = async () => {
  await router.push('/home/friend')
}

const chat = async () => {
  await router.push('/home/chat')
}

const diary = async () => {
  await router.push('/home/diary')
}

const notice = async () => {
  await router.push('/home/notice')
}

const me = async () => {
  await router.push('/home/me')
}

const logout = async () => {
  instance.get('http://localhost:8080/index/logout').then(res => {
    if (res.data.code === 200) {
      ElMessage({
        message: '拜拜咯👋期待我们的下一次相遇',
        type: 'success',
      })
      router.push('/')
    } else {
      ElMessage.error(res.data.message)
    }

  })
}

</script>

<style scoped>
.el-container{
  height: 100%;
}

.el-main {
  height: 100%;
  width: 100%;
  margin: 0;
  padding: 0;
}

.el-menu-vertical:not(.el-menu--collapse) {
  width: 100%;
  min-height: 500px;
}
</style>