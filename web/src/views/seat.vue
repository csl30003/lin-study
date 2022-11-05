<template>
  <el-image :src="require('@/assets/seat.png')"
            style="margin: 0; position:absolute; height: 100%; width: 92%;"></el-image>
  <div class="div">
    <div id="seat1" @click="">

    </div>
    <div id="seat2" @click="">

    </div>
    <div id="seat3" @click="">

    </div>
    <div id="seat4" @click="">

    </div>
    <div id="seat5" @click="">

    </div>
    <div id="seat6" @click="">

    </div>
    <div id="seat7" @click="">

    </div>
    <div id="seat8" @click="">

    </div>
    <div id="seat9" @click="">

    </div>
    <div id="seat10" @click="">

    </div>
  </div>
</template>

<script setup>

import {onMounted, ref} from "vue";

const classroomID = ref()
const seatList = ref([])

onMounted(() => {
  getClassroomID()
})

const getClassroomID = async () => {

}

const seat = async () => {
//  判断有没有人，没人就入座，有人就判断是不是自己，是就离座，不是就提示没法入座
//  可以用element plus的抽屉做选择专注目标
}

</script>

<style scoped>
.div {
  height: 100%;
}

#seat1 {
  position: absolute;
  top: 31%;
  left: 25%;
  height: 10%;
  width: 10%;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 1);
  text-align: center;
}
</style>