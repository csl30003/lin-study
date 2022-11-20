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
</template>

<script setup>

import {onMounted, reactive, ref} from "vue";
import instance from "@/axios";
import {ElMessage} from "element-plus";
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

onMounted(() => {
  getStudent()
  getClassroomID()
  getSeatInfo()
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

    instanceXWWWForm.patch('http://localhost:8080/index/classroom/seat', formData).then(async res => {
      if (res.data.code === 200) {
        await getSeatInfo()
      } else {
        ElMessage.error(res.data.message)
      }
    })
  } else if (name === student.name) {
    //  离座
  } else {
    //  查看对方的信息
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
</style>