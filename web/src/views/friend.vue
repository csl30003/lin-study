<template>
  <el-tabs v-model="activeName" class="el-tabs" @tab-click="handleClick" type="border-card" stretch>

    <el-tab-pane class="el-tab-pane" label="相互的朋友" name="mutualFriend">

      <el-descriptions
          class="el-descriptions"
          v-for="(friend, index) in friendList" :key="index"
          :column="3"
          :size="'large'"
          border

      >
        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <UserFilled/>
              </el-icon>
              昵称
            </div>
          </template>
          {{ friend.name }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <Tickets/>
              </el-icon>
              目标
            </div>
          </template>
          {{ friend.goal }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <CollectionTag/>
              </el-icon>
              标签
            </div>
          </template>
          {{ friend.label }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon v-if="friend.sex === '男'">
                <Male/>
              </el-icon>
              <el-icon v-else>
                <Female/>
              </el-icon>
              性别
            </div>
          </template>
          {{ friend.sex }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <Clock/>
              </el-icon>
              总专注时长
            </div>
          </template>
          {{ friend.accumulated_concentrate_time }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <More/>
              </el-icon>
              操作
            </div>
          </template>
          <el-button type="primary" @click="deepenUnderstanding(friend.id)">深入了解</el-button>
          <el-button type="info" @click="chat">传递纸条</el-button>
          <el-button type="danger" @click="breakOffRelations(friend.id)">狠心绝交</el-button>
        </el-descriptions-item>
      </el-descriptions>

    </el-tab-pane>

    <el-tab-pane label="想交的朋友" name="myFriend">

      <el-descriptions
          class="el-descriptions"
          v-for="(friend, index) in friendList" :key="index"
          :column="3"
          :size="'large'"
          border

      >
        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <UserFilled/>
              </el-icon>
              昵称
            </div>
          </template>
          {{ friend.name }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <Tickets/>
              </el-icon>
              目标
            </div>
          </template>
          {{ friend.goal }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <CollectionTag/>
              </el-icon>
              标签
            </div>
          </template>
          {{ friend.label }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon v-if="friend.sex === '男'">
                <Male/>
              </el-icon>
              <el-icon v-else>
                <Female/>
              </el-icon>
              性别
            </div>
          </template>
          {{ friend.sex }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <Clock/>
              </el-icon>
              总专注时长
            </div>
          </template>
          {{ friend.accumulated_concentrate_time }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <More/>
              </el-icon>
              操作
            </div>
          </template>
          <el-button type="primary">深入了解</el-button>
          <el-button type="danger" @click="breakOffRelations(friend.id)">狠心绝交</el-button>
        </el-descriptions-item>
      </el-descriptions>

    </el-tab-pane>

    <el-tab-pane label="想认识我的人" name="noMyFriend">

      <el-descriptions
          class="el-descriptions"
          v-for="(friend, index) in friendList" :key="index"
          :column="3"
          :size="'large'"
          border

      >
        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <UserFilled/>
              </el-icon>
              昵称
            </div>
          </template>
          {{ friend.name }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <Tickets/>
              </el-icon>
              目标
            </div>
          </template>
          {{ friend.goal }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <CollectionTag/>
              </el-icon>
              标签
            </div>
          </template>
          {{ friend.label }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon v-if="friend.sex === '男'">
                <Male/>
              </el-icon>
              <el-icon v-else>
                <Female/>
              </el-icon>
              性别
            </div>
          </template>
          {{ friend.sex }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <Clock/>
              </el-icon>
              总专注时长
            </div>
          </template>
          {{ friend.accumulated_concentrate_time }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <More/>
              </el-icon>
              操作
            </div>
          </template>
          <el-button type="primary">深入了解</el-button>
          <el-button type="success" @click="makeFriend(friend.id)">交个朋友</el-button>
        </el-descriptions-item>
      </el-descriptions>

    </el-tab-pane>

    <el-tab-pane label="搜索" name="search">
      <el-row>
        <el-col :offset="7">
          <el-input class="search-input" v-model="search" placeholder="输入用户名" clearable/>
          <el-button :icon="Search" circle @click="searchStudent"/>
        </el-col>
      </el-row>

      <el-descriptions
          class="el-descriptions"
          v-for="(friend, index) in friendList" :key="index"
          :column="3"
          :size="'large'"
          border

      >
        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <UserFilled/>
              </el-icon>
              昵称
            </div>
          </template>
          {{ friend.name }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <Tickets/>
              </el-icon>
              目标
            </div>
          </template>
          {{ friend.goal }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <CollectionTag/>
              </el-icon>
              标签
            </div>
          </template>
          {{ friend.label }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon v-if="friend.sex === '男'">
                <Male/>
              </el-icon>
              <el-icon v-else>
                <Female/>
              </el-icon>
              性别
            </div>
          </template>
          {{ friend.sex }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <Clock/>
              </el-icon>
              总专注时长
            </div>
          </template>
          {{ friend.accumulated_concentrate_time }}
        </el-descriptions-item>

        <el-descriptions-item>
          <template #label>
            <div class="cell-item">
              <el-icon>
                <More/>
              </el-icon>
              操作
            </div>
          </template>
          <el-button type="primary">深入了解</el-button>
          <el-button type="success" @click="makeFriend(friend.id)">交个朋友</el-button>
        </el-descriptions-item>
      </el-descriptions>

    </el-tab-pane>

  </el-tabs>

</template>

<script setup>
import {nextTick, reactive, ref} from 'vue'
import instance from "@/axios";
import {ElMessage} from "element-plus";
import {
  UserFilled,
  Male,
  Female,
  Clock,
  More,
  CollectionTag,
  Tickets,
  Search
} from '@element-plus/icons-vue'
import {CircleCloseFilled} from '@element-plus/icons-vue'
import router from "@/router";

const activeName = ref('')

let friendList = ref([])

const search = ref('')

let student = reactive({
  id: 0,
  name: "",
  label: "",
  sex: 0,
  area: "",
  stage: "",
  school: "",
  goal: "",
  total_study_day: 0,
  continuous_study_day: 0,
  accumulated_concentrate_time: 0,
  inspired: 0,
  inspire: 0,
  status: 0
})

const handleClick = (tab) => {
  friendList.value = []

  if (tab.paneName === "mutualFriend") {
    instance.get('http://localhost:8080/index/mutualFriend').then(res => {
      if (res.data.code === 200 && res.data.data != null) {
        for (const datum of res.data.data) {
          if (datum.sex === 0) {
            datum.sex = "男"
          } else {
            datum.sex = "女"
          }
          friendList.value.push(datum)
        }
      } else {
        ElMessage.error(res.data.message)
      }
    })
  } else if (tab.paneName === "myFriend") {
    instance.get('http://localhost:8080/index/myFriend').then(res => {
      if (res.data.code === 200 && res.data.data != null) {
        for (const datum of res.data.data) {
          if (datum.sex === 0) {
            datum.sex = "男"
          } else {
            datum.sex = "女"
          }
          friendList.value.push(datum)
        }
      } else {
        ElMessage.error(res.data.message)
      }
    })
  } else if (tab.paneName === "noMyFriend") {
    instance.get('http://localhost:8080/index/noMyFriend').then(res => {
      if (res.data.code === 200 && res.data.data != null) {
        for (const datum of res.data.data) {
          if (datum.sex === 0) {
            datum.sex = "男"
          } else {
            datum.sex = "女"
          }
          friendList.value.push(datum)
        }
      } else {
        ElMessage.error(res.data.message)
      }
    })
  }
}

const searchStudent = async () => {
  friendList.value = []
  instance.get('http://localhost:8080/index/searchStudent', {
    params: {
      searchContext: search.value
    }
  }).then(res => {
    if (res.data.code === 200) {
      for (const datum of res.data.data) {
        if (datum.sex === 0) {
          datum.sex = "男"
        } else {
          datum.sex = "女"
        }
        friendList.value.push(datum)
      }
    } else {
      ElMessage.error(res.data.message)
    }
  })
}

//  暂时无法解决
const deepenUnderstanding = async (id) => {
  student = null

  instance.get('http://localhost:8080/index/student', {
    params: {
      id: id
    }
  }).then(res => {
    if (res.data.code === 200) {
      student = res.data.data
    } else {
      ElMessage.error(res.data.message)
    }
  })
}

const makeFriend = async (id) => {
  instance.post('http://localhost:8080/index/friend', {
    object_id: id
  }).then(res => {
    if (res.data.code === 200) {
      ElMessage.success(res.data.message)
    } else {
      ElMessage.error(res.data.message)
    }
  })
  await router.push('/home/map')
  await router.push('/home/friend')
}

const breakOffRelations = async (id) => {
  instance.delete('http://localhost:8080/index/friend', {
    data: {object_id: id}
  }).then(res => {
    if (res.data.code === 200) {
      ElMessage.success(res.data.message)
    } else {
      ElMessage.error(res.data.message)
    }
  })
  await router.push('/home/map')
  await router.push('/home/friend')
}

const chat = async () => {
  await router.push('/home/chat')
}
</script>

<style scoped>
.el-descriptions {
  margin-top: 10px;
}

.search-input {
  width: 40%;
  height: 50px;
  margin-right: 5px;
}
</style>