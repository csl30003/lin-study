<template>
  <el-image :src="require('@/assets/seat.png')"
            style="margin: 0; position:absolute; height: 100%; width: 92%;"></el-image>
  <div class="div">
    <div id="seat1" @click="seat(seatInfo.seat1, 1)">
      <h2 v-if="seatInfo.seat1 === '_'"></h2>
      <h2 v-else>{{ seatInfo.seat1 }}</h2>
    </div>
    <div id="seat2" @click="seat(seatInfo.seat2,2)">
      <h2 v-if="seatInfo.seat2 === '_'"></h2>
      <h2 v-else>{{ seatInfo.seat2 }}</h2>
    </div>
    <div id="seat3" @click="seat(seatInfo.seat3,3)">
      <h2 v-if="seatInfo.seat3 === '_'"></h2>
      <h2 v-else>{{ seatInfo.seat3 }}</h2>
    </div>
    <div id="seat4" @click="seat(seatInfo.seat4,4)">
      <h2 v-if="seatInfo.seat4 === '_'"></h2>
      <h2 v-else>{{ seatInfo.seat4 }}</h2>
    </div>
    <div id="seat5" @click="seat(seatInfo.seat5,5)">
      <h2 v-if="seatInfo.seat5 === '_'"></h2>
      <h2 v-else>{{ seatInfo.seat5 }}</h2>
    </div>
    <div id="seat6" @click="seat(seatInfo.seat6,6)">
      <h2 v-if="seatInfo.seat6 === '_'"></h2>
      <h2 v-else>{{ seatInfo.seat6 }}</h2>
    </div>
    <div id="seat7" @click="seat(seatInfo.seat7,7)">
      <h2 v-if="seatInfo.seat7 === '_'"></h2>
      <h2 v-else>{{ seatInfo.seat7 }}</h2>
    </div>
    <div id="seat8" @click="seat(seatInfo.seat8,8)">
      <h2 v-if="seatInfo.seat8 === '_'"></h2>
      <h2 v-else>{{ seatInfo.seat8 }}</h2>
    </div>
    <div id="seat9" @click="seat(seatInfo.seat9,9)">
      <h2 v-if="seatInfo.seat9 === '_'"></h2>
      <h2 v-else>{{ seatInfo.seat9 }}</h2>
    </div>
    <div id="seat10" @click="seat(seatInfo.seat10,10)">
      <h2 v-if="seatInfo.seat10 === '_'"></h2>
      <h2 v-else>{{ seatInfo.seat10 }}</h2>
    </div>
  </div>

  <el-drawer
      v-model="target"
      title="请选择本次专注的目标和专注的时长"
      direction="rtl"
      size="40%"
      :before-close="chooseTargetClose"
  >
    <h3>专注目标</h3>
    <br>

    <el-check-tag
        v-for="tag in targetTags"
        :key="tag"
        size="large"
        round
        class="mx-1"
        closable
        :disable-transitions="false"
        effect="dark"
        :checked="tag.checked"
        @click="choose(tag.id, tag.checked)"
    >
      {{ tag.target }}
    </el-check-tag>

    <el-input
        v-if="addInputVisible"
        ref="addInputRef"
        v-model="addInputValue"
        class="ml-1 w-20"
        size="large"
        @keyup.enter="addTarget"
        @blur="addTarget"
    />
    <el-button v-else class="button-new-tag ml-1" type="success" size="default" @click="addShowInput">
      + 添加专注目标
    </el-button>

    <el-input
        v-if="deleteInputVisible"
        ref="deleteInputRef"
        v-model="deleteInputValue"
        class="ml-1 w-20"
        size="large"
        @keyup.enter="deleteTarget"
        @blur="deleteTarget"
    />
    <el-button v-else class="button-new-tag ml-1" type="danger" size="default" @click="deleteShowInput">
      + 删除专注目标
    </el-button>

    <br>
    <br>
    <br>
    <br>
    <br>

    <h3>专注时长(分钟)</h3>
    <br>

    <el-input
        v-model="concentrateTime"
        size="large"
        style="width: 50%"
    />
    <br>
    <br>
    <el-button type="primary" @click="beginConcentrate">开始专注！！！</el-button>
  </el-drawer>

  <el-dialog
      v-model="centerDialogVisible"
      title="离开"
      width="30%"
      align-center
  >
    <span>确定要打退堂鼓吗？现在离开的话，是没有专注记录的喔！</span>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="manualUnseat">
          还是离开
        </el-button>
        <el-button type="primary" @click="centerDialogVisible = false">
          继续学习
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>

import {nextTick, onMounted, onUnmounted, reactive, ref} from "vue";
import instance from "@/axios";
import {ElMessage, ElButton, ElDrawer, ElInput} from "element-plus";
import {useRouter} from "vue-router";
import axios from "axios";

let student = reactive({
  id: '',
  name: '',
})
const router = useRouter()
const floor = router.currentRoute.value.params.floor
const layer = router.currentRoute.value.params.layer
const class_ = router.currentRoute.value.params.class

const classroomID = ref()

const seatInfo = reactive({
  seat1: '',
  seat2: '',
  seat3: '',
  seat4: '',
  seat5: '',
  seat6: '',
  seat7: '',
  seat8: '',
  seat9: '',
  seat10: '',
})

//  为了给离座方法传参
const seatTemp = ref(0)

const target = ref(false)
const addInputValue = ref('')
const targetTags = ref([])
const addInputVisible = ref(false)
const addInputRef = ref()
const chooseTargetID = ref(0)
const concentrateTime = ref('')
const deleteInputVisible = ref(false)
const deleteInputValue = ref('')
const deleteInputRef = ref()

const centerDialogVisible = ref(false)

onMounted(() => {
  getStudent()
  getClassroomID()
  getSeatInfo()
  getConcentrateTarget()
})

onUnmounted(() => {
  unseatButNoMessage(student.name, seatTemp.value)
})

const getStudent = async () => {
  //  我知道重新获取个人信息很捞，明明在父组件就已经获取过了，但是我现在不知道怎么传值，暂时这样
  instance.get('http://localhost:8080/index/getStudentInfo').then(res => {
    if (res.data.code === 200) {
      student.id = res.data.data.id
      student.name = res.data.data.name
    } else {
      ElMessage.error('无法获取你的消息')
    }
  })
}

const getClassroomID = async () => {
  instance.get('http://localhost:8080/index/' + floor + '/' + layer + '/' + class_ + '/getClassroomID').then(res => {
    if (res.data.code === 200) {
      classroomID.value = res.data.data
    } else {
      ElMessage.error('无法获取教室信息')
    }
  })
}

const getSeatInfo = async () => {
  instance.get('http://localhost:8080/index/' + floor + '/' + layer + '/' + class_).then(res => {
    if (res.data.code === 200) {
      seatInfo.seat1 = res.data.data.seat1
      seatInfo.seat2 = res.data.data.seat2
      seatInfo.seat3 = res.data.data.seat3
      seatInfo.seat4 = res.data.data.seat4
      seatInfo.seat5 = res.data.data.seat5
      seatInfo.seat6 = res.data.data.seat6
      seatInfo.seat7 = res.data.data.seat7
      seatInfo.seat8 = res.data.data.seat8
      seatInfo.seat9 = res.data.data.seat9
      seatInfo.seat10 = res.data.data.seat10
    } else {
      ElMessage.error('无法获取座位信息')
    }
  })
}

const getConcentrateTarget = async () => {
  targetTags.value = []
  instance.get('http://localhost:8080/index/concentrateTarget').then(res => {
    if (res.data.code === 200) {
      for (const datum of res.data.data) {
        targetTags.value.push({
          id: datum.id,
          target: datum.target,
          checked: false
        })
      }
    } else {
      ElMessage.error('无法获取你的好友')
    }
  })
}

const seat = async (name, seat) => {
//  判断有没有人，没人就入座，有人就判断是不是自己，是就离座，不是就查看对方的信息
//  可以用element plus的抽屉做选择专注目标
  if (name === '_') {
    //  入座
    const instanceXWWWForm = axios.create({
      timeout: 6000,
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      withCredentials: true
    })

    let formData = new FormData()
    formData.append('seat', seat.toString())
    formData.append('classroom_id', classroomID.value.toString())

    seatTemp.value = seat

    instanceXWWWForm.patch('http://localhost:8080/index/classroom/seat', formData).then(async res => {
      if (res.data.code === 200) {
        await getSeatInfo()

        //  开始选择专注目标
        target.value = true
      } else {
        ElMessage.error(res.data.message)
      }
    })
  } else if (name === student.name) {
    //  离座
    centerDialogVisible.value = true
  } else {
    //  查看对方的信息
  }
}

const unseat = async (name, seat) => {
  //  离座
  const instanceXWWWForm = axios.create({
    timeout: 6000,
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    withCredentials: true
  })

  let formData = new FormData()
  formData.append('seat', seat.toString())
  formData.append('classroom_id', classroomID.value.toString())

  seatTemp.value = 0
  console.log(seat.toString(), classroomID.value.toString())

  instanceXWWWForm.patch('http://localhost:8080/index/classroom/unseat', formData).then(async res => {
    if (res.data.code === 200) {
      await getSeatInfo()
      ElMessage.success('已离开座位')
    } else {
      ElMessage.error(res.data.message)
    }
  })
}

const unseatButNoMessage = async (name, seat) => {
  //  离座但是没有消息反馈
  const instanceXWWWForm = axios.create({
    timeout: 6000,
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    withCredentials: true
  })

  let formData = new FormData()
  formData.append('seat', seat.toString())
  formData.append('classroom_id', classroomID.value.toString())

  seatTemp.value = 0
  console.log(seat.toString(), classroomID.value.toString())

  instanceXWWWForm.patch('http://localhost:8080/index/classroom/unseat', formData).then(async res => {
    if (res.data.code === 200) {
      await getSeatInfo()
    }
  })
}

const manualUnseat = async () => {
  //  手动离座
  centerDialogVisible.value = false
  await unseat(student.name, seatTemp.value)
}

const chooseTargetClose = async () => {
  target.value = false
  await unseat(student.name, seatTemp.value)
}

const deleteShowInput = async () => {
  deleteInputVisible.value = true
  await nextTick(() => {
    deleteInputRef.value.input.focus()
  })
}

const deleteTarget = async () => {
  if (deleteInputValue.value) {
    //  删除专注目标
    let tag = deleteInputValue.value

    instance.delete('http://localhost:8080/index/concentrateTarget', {
      data: {
        target: deleteInputValue.value
      }
    }).then(res => {
      if (res.data.code === 200) {
        ElMessage({
          message: res.data.message,
          type: 'success',
        })

        let id1 = targetTags.value.findIndex(item => {
          if (item.target === tag) {
            return true
          }
        })

        targetTags.value.splice(id1, 1)
      } else {
        ElMessage.error(res.data.message)
      }
    })
  }
  deleteInputVisible.value = false
  deleteInputValue.value = ''
}

const addShowInput = async () => {
  addInputVisible.value = true
  await nextTick(() => {
    addInputRef.value.input.focus()
  })
}

const addTarget = async () => {
  if (addInputValue.value) {
    //  添加专注目标
    //  我不知道为什么post之后inputValue.value值就没了，离谱
    let tag = addInputValue.value

    instance.post('http://localhost:8080/index/concentrateTarget', {
      target: addInputValue.value
    }).then(res => {
      if (res.data.code === 200) {
        ElMessage({
          message: res.data.message,
          type: 'success',
        })

        targetTags.value.push({
          id: res.data.data,
          target: tag,
          checked: false
        })
      } else {
        ElMessage.error(res.data.message)
      }
    })
  }
  addInputVisible.value = false
  addInputValue.value = ''
}

const choose = async (id, checked) => {
  for (const tag of targetTags.value) {
    if (tag.id !== id) {
      tag.checked = false
    } else {
      tag.checked = !checked
      if (tag.checked === true) {
        chooseTargetID.value = id
      } else {
        chooseTargetID.value = 0
      }
    }
  }
}

const beginConcentrate = () => {
  ElMessage.success(chooseTargetID.value + "  " + concentrateTime.value)
}
</script>

<style scoped>
.div {
  height: 100%;
}

h2 {
  font-family: FZShuTi, serif;
  font-size: large;
  vertical-align: middle;
  transform: translateY(70%);
  display: inline-block;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;

}

#seat1 {
  position: absolute;
  top: 31%;
  left: 25%;
  height: 10%;
  width: 10%;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 1);
  text-align: center;
  margin: 0 auto;
}

#seat2 {
  position: absolute;
  top: 31%;
  left: 37%;
  height: 10%;
  width: 10%;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 1);
  text-align: center;
  margin: 0 auto;
}

#seat3 {
  position: absolute;
  top: 31%;
  left: 49%;
  height: 10%;
  width: 10%;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 1);
  text-align: center;
}

#seat4 {
  position: absolute;
  top: 31%;
  left: 61%;
  height: 10%;
  width: 10%;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 1);
  text-align: center;
}

#seat5 {
  position: absolute;
  top: 31%;
  left: 73%;
  height: 10%;
  width: 10%;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 1);
  text-align: center;
}

#seat6 {
  position: absolute;
  top: 57%;
  left: 25%;
  height: 10%;
  width: 10%;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 1);
  text-align: center;
}

#seat7 {
  position: absolute;
  top: 57%;
  left: 37%;
  height: 10%;
  width: 10%;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 1);
  text-align: center;
}

#seat8 {
  position: absolute;
  top: 57%;
  left: 49%;
  height: 10%;
  width: 10%;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 1);
  text-align: center;
}

#seat9 {
  position: absolute;
  top: 57%;
  left: 61%;
  height: 10%;
  width: 10%;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 1);
  text-align: center;
}

#seat10 {
  position: absolute;
  top: 57%;
  left: 73%;
  height: 10%;
  width: 10%;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 1);
  text-align: center;
}

.mx-1 {
  margin-right: 8px;
}

.dialog-footer button:first-child {
  margin-right: 10px;
}
</style>