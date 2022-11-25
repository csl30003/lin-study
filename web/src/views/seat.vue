<template>
  <div class="div" @mouseleave="concentrateOut">
    <el-image :src="require('@/assets/seat.png')"
              style="margin: 0; position:absolute; height: 100%; width: 92%;"></el-image>
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

    <div id="time" class="time">
      <el-icon size="large">
        <Timer/>
      </el-icon>
      <h3>倒计时: {{ min }}分钟 {{ sec }}秒</h3>
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
        title="请做出你的选择"
        width="30%"
        align-center
    >
      <span>{{ student.name }}同学想做些什么呢？</span>
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="manualUnseat">
          离开
        </el-button>
        <el-button type="primary" @click="centerDialogVisible = false">
          休息 休息一下
        </el-button>
        <el-button type="success" @click="centerDialogVisible = false;target = true">
          开始专注
        </el-button>
      </span>
      </template>
    </el-dialog>

    <el-dialog
        v-model="concentrateOverVisible"
        title="太棒啦"
        width="30%"
        align-center
    >
      <span>恭喜你完成本次专注！</span>
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="manualUnseat">
          我想回家了
        </el-button>
        <el-button type="primary" @click="concentrateOverVisible = false">
          小休一下
        </el-button>
        <el-button type="success" @click="concentrateOverVisible = false;target = true">
          继续狠狠专注
        </el-button>
      </span>
      </template>
    </el-dialog>

    <el-dialog v-model="outVisible" :show-close="false" :close-on-click-modal="false">
      <template #header="{ titleId, titleClass }">
        <div class="out-header">
          <h4 :id="titleId" :class="titleClass">遛号啦！！！！！！！！</h4>
          <el-button type="success" @click="concentrateBack">
            <el-icon size="large">
              <Position/>
            </el-icon>
            我回来了
          </el-button>
        </div>
      </template>
      <el-icon size="large">
        <Stopwatch/>
      </el-icon>
      三十秒内没回来专注记录就没有啦！QAQ
      <h3>倒计时: {{ outSec }}秒</h3>
    </el-dialog>
  </div>
</template>

<script setup>

import {nextTick, onMounted, onUnmounted, reactive, ref} from "vue";
import instance from "@/axios";
import {ElMessage, ElButton, ElDrawer, ElInput, ElNotification} from "element-plus";
import {useRouter} from "vue-router";
import axios from "axios";
import {
  Timer,
  Stopwatch,
  Position,
} from '@element-plus/icons-vue'

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
const concentrateOverVisible = ref(false)

const outVisible = ref(false)

const min = ref('_')
const sec = ref('_')
//以下两个时间变量是为了遛号保存时间服务的
const minTemp = ref('_')
const secTemp = ref('_')
//outSec是为了遛号倒计时
const outSec = ref('_')


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

  instanceXWWWForm.patch('http://localhost:8080/index/classroom/unseat', formData).then(async res => {
    if (res.data.code === 200) {
      await getSeatInfo()
      ElMessage.success('已离开座位')
    } else {
      ElMessage.error(res.data.message)
    }
  })

  await quitConcentrate()
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

  instanceXWWWForm.patch('http://localhost:8080/index/classroom/unseat', formData).then(async res => {
    if (res.data.code === 200) {
      await getSeatInfo()
    }
  })

  await quitConcentrate()
}

const manualUnseat = async () => {
  //  手动离座
  centerDialogVisible.value = false
  concentrateOverVisible.value = false
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

const beginConcentrate = async () => {
  if (chooseTargetID.value === 0) {
    ElMessage.warning('还没有选择专注目标呢！')
  } else {
    let concentrateTimeTemp
    try {
      concentrateTimeTemp = parseInt(concentrateTime.value)
    } catch (e) {
      ElMessage.warning('没有输入正确的专注时长！')
      return
    }

    instance.post('http://localhost:8080/index/classroom/concentrate', {
      concentrate_target_id: chooseTargetID.value,
      concentrate_time: concentrateTimeTemp,
    }).then(async res => {
      if (res.data.code === 200) {
        ElMessage.success(res.data.message)
        // 计时器
        min.value = '0'
        sec.value = '0'
        await countDown(concentrateTimeTemp * 60)
      } else {
        ElMessage.error(res.data.message)
      }
    })

    target.value = false
  }
}

const quitConcentrate = async () => {
  if (chooseTargetID.value !== 0) {
    instance.delete('http://localhost:8080/index/classroom/concentrate').then(res => {
      if (res.data.code === 200) {
        ElMessage({
          message: res.data.message,
          type: 'success',
        })
      } else {
        ElMessage.error(res.data.message)
      }
    })
    min.value = '_'
    sec.value = '_'
  }
}

const endConcentrate = async () => {
  if (chooseTargetID.value !== 0) {
    instance.patch('http://localhost:8080/index/classroom/concentrate').then(res => {
      if (res.data.code === 200) {
        ElMessage({
          message: res.data.message,
          type: 'success',
        })
      } else {
        ElMessage.error(res.data.message)
      }
    })
    min.value = '_'
    sec.value = '_'
  }
}

//倒计时
const countDown = async (msec) => {
  if (msec < 0) {
    return
  }
  if (min.value === '_' && sec.value === '_') {
    //强退
    await quitConcentrate()
    return
  }

  if (min.value === 'out' && sec.value === 'out') {
    //遛号中
    return
  }

  let minInt = parseInt(msec / 60 % 60)
  let secInt = parseInt(msec % 60)
  min.value = minInt.toString()
  sec.value = secInt > 9 ? secInt : '0' + secInt
  if (minInt >= 0 && secInt >= 0) {
    if (minInt === 0 && secInt === 0) {
      //倒计时结束
      await endConcentrate()

      //弹窗
      concentrateOverVisible.value = true
      return
    }
    setTimeout(function () {
      countDown(msec - 1)
    }, 1000)
  }
}

const concentrateOut = async () => {
  if (min.value !== '_' && sec.value !== '_') {
    //正在专注但却遛号
    instance.patch('http://localhost:8080/index/classroom/concentrateOut').then(async res => {
      if (res.data.code === 200) {
        ElMessage.success(res.data.message)
        outVisible.value = true

        minTemp.value = min.value
        secTemp.value = sec.value
        min.value = 'out'
        sec.value = 'out'

        //启动30秒倒计时
        outSec.value = '30'
        await outCountDown(30)
      } else {
        ElMessage.error(res.data.message)
      }
    })
  }
}

const outCountDown = async (msec) => {
  if (msec < 0) {
    outSec.value = '_'
    return
  }
  if (outSec.value === '_') {
    //已恢复
    return
  }

  let outSecInt = parseInt(msec)
  outSec.value = outSecInt.toString()
  if (outSecInt >= 0) {
    if (outSecInt === 0) {
      //倒计时结束 成功遛号
      await quitConcentrate()

      outVisible.value = false
      outSec.value = '_'

      ElNotification({
        title: '遛号啦',
        message: '专注失败惹 记录也被抹去 下次一定要认真哦😫',
        duration: 0,
      })
      return
    }
    setTimeout(function () {
      outCountDown(msec - 1)
    }, 1000)
  }
}

const concentrateBack = async () => {
  //正在遛号但却改过自新
  if (min.value === 'out' && sec.value === 'out' && outSec.value !== '_') {
    instance.patch('http://localhost:8080/index/classroom/concentrateBack').then(async res => {
      if (res.data.code === 200) {
        ElMessage.success(res.data.message)
        min.value = minTemp.value
        sec.value = secTemp.value
        minTemp.value = '_'
        secTemp.value = '_'
        outSec.value = '_'
        outVisible.value = false

        ElNotification({
          title: '回来辣',
          message: '要专心致志学习哦！',
        })

        let minInt = parseInt(min.value)
        let secInt = parseInt(sec.value)
        await countDown(minInt * 60 + secInt)
      } else {
        ElMessage.error(res.data.message)
      }
    })
  }
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

#time {
  position: absolute;
  top: 44%;
  left: 39%;
  height: 10%;
  width: 30%;
  /*box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 1);*/
  text-align: center;
}

.mx-1 {
  margin-right: 8px;
}

.dialog-footer button:first-child {
  margin-right: 10px;
}

.out-header {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
</style>