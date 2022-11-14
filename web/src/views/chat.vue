<template>
  <div class="common-layout">
    <el-container class="el-container">
      <el-aside width="20%">
        <el-table class="el-table" :data="friendList" size="large" :row-style="{height: '100%'}"
                  :cell-style="{padding:'0px'}" @row-click="getChat">
          <el-table-column prop="name" lable="用户名" min-width="100%" align="center"/>
        </el-table>
      </el-aside>

      <el-main class="el-main">

      </el-main>
    </el-container>
  </div>
</template>

<script setup>

import {getCurrentInstance, onMounted, reactive, ref} from "vue";
import instance from "@/axios";
import {ElMessage} from "element-plus";
import {JSEncrypt} from 'jsencrypt';

let friendList = ref([])
let serverPublicKey = ref()
let chatList = ref([])
// 无奈之举，.env配置文件里不能读取换行的字符串，加上\n又不能识别到，这是最笨的解决方法了，尽力了，前端不是很会。
let PRIVATE_KEY = process.env.VUE_APP_PRIVATE_KEY1 + "\n" +
    process.env.VUE_APP_PRIVATE_KEY2 + "\n" +
    process.env.VUE_APP_PRIVATE_KEY3 + "\n" +
    process.env.VUE_APP_PRIVATE_KEY4 + "\n" +
    process.env.VUE_APP_PRIVATE_KEY5 + "\n" +
    process.env.VUE_APP_PRIVATE_KEY6 + "\n" +
    process.env.VUE_APP_PRIVATE_KEY7 + "\n" +
    process.env.VUE_APP_PRIVATE_KEY8 + "\n" +
    process.env.VUE_APP_PRIVATE_KEY9 + "\n" +
    process.env.VUE_APP_PRIVATE_KEY10 + "\n" +
    process.env.VUE_APP_PRIVATE_KEY11 + "\n" +
    process.env.VUE_APP_PRIVATE_KEY12 + "\n" +
    process.env.VUE_APP_PRIVATE_KEY13 + "\n" +
    process.env.VUE_APP_PRIVATE_KEY14 + "\n" +
    process.env.VUE_APP_PRIVATE_KEY15 + "\n" +
    process.env.VUE_APP_PRIVATE_KEY16
let PUBLIC_KEY = process.env.VUE_APP_PUBLIC_KEY1 + "\n" +
    process.env.VUE_APP_PUBLIC_KEY2 + "\n" +
    process.env.VUE_APP_PUBLIC_KEY3 + "\n" +
    process.env.VUE_APP_PUBLIC_KEY4 + "\n" +
    process.env.VUE_APP_PUBLIC_KEY5 + "\n" +
    process.env.VUE_APP_PUBLIC_KEY6

const form = reactive({
  recipient_id: 0,
  chat_content: '',
})
const {appContext} = getCurrentInstance()

onMounted(() => {
  getFriend()
  getPublicKey()
})

const getFriend = async () => {
  instance.get('http://localhost:8080/index/mutualFriend').then(res => {
    if (res.data.code === 200) {
      for (const datum of res.data.data) {
        friendList.value.push(datum)
      }
    } else {
      ElMessage.error('无法获取好友')
    }
  })
}

const getPublicKey = async () => {
  instance.get('http://localhost:8080/index/getPublicKey').then(res => {
    if (res.data.code === 200) {
      serverPublicKey = res.data.data
    } else {
      ElMessage.error('无法获取公钥')
    }
  })
}

const getChat = async (row) => {
  instance.post('http://localhost:8080/index/getChat', {
    recipient_id: row.id,
    public_key: PUBLIC_KEY,
  }).then(res => {
    if (res.data.code === 200) {
      chatList.value = []
      for (const datum of res.data.data) {
        datum.chat_content = RSADecrypt(datum.chat_content)
        chatList.value.push(datum)
      }
      console.log(chatList)
    } else {
      ElMessage.error(res.data.message)
    }
  })
}

const sendMessageTest = async () => {
  // let chat_content = RSAEncrypt(form.chat_content)

  let chat_content = RSAEncrypt("hi")

  instance.post('http://localhost:8080/index/chat', {
    // recipient_id: parseInt(form.recipient_id),
    // chat_content: chat_content,
    recipient_id: 2,
    chat_content: chat_content,
  }).then(res => {
    if (res.data.code === 200) {
      ElMessage.error(res.data.message)
    } else {
      ElMessage.error(res.data.message)
    }
  })
}

function RSAEncrypt(plainText) {
  let jse = new JSEncrypt();
  jse.setPublicKey(serverPublicKey); // 设置 加密公钥
  return jse.encrypt(plainText); // 进行加密
}

function RSADecrypt(cipherText) {
  let jse = new JSEncrypt();
  jse.setPrivateKey(PRIVATE_KEY);
  return jse.decrypt(cipherText);
}
</script>

<style scoped>
.el-container {
  height: 100%;
}

.el-main {
  height: 100%;
  width: 100%;
  margin: 0;
  padding: 0;
}

.el-table {
  width: 100%;
  margin: auto;
  font-size: 30px;
}

::v-deep .el-table .cell {
  line-height: 5vw;
}
</style>