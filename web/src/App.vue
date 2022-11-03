<template>
  <el-container height="100%">
    <el-header class="el-header" height="100px">
      <h1 class="h1">LinStudy</h1>
    </el-header>
    <router-view></router-view>
  </el-container>
</template>

<style scoped>
.el-header{
  display: table-cell;
  vertical-align: middle;
  background-color: #ffe05d;
}
.h1 {
  text-align: center;
  font-size: 60px;
  color: white;
  transform: translateY(-50%);
}
</style>