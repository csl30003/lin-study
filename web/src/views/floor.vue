<template>
  <el-image :src="require('@/assets/floor.png')"
            style="margin: 0; position:relative; height: 100%; width: 100%"></el-image>

  <div class="radius" @click="getLayer('大学部')">
    <h1>大学部</h1>
  </div>


</template>

<script setup>
import {useRouter} from "vue-router"

const router = useRouter()

const getLayer = async (floor) => {
  await router.push('/home/map/' + floor)
}
</script>

<style scoped>
.radius {
  position: absolute;
  top: 37%;
  left: 68.5%;
  height: 24%;
  width: 21%;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 1);
  text-align: center;
}
</style>