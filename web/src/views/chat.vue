<template>
  <div class="common-layout">
    <el-container class="el-container">
      <el-aside width="20%">
        <el-table class="el-table" :data="friendList" size="large" :row-style="{height: '100%'}"
                  :cell-style="{padding:'0px'}" @row-click="getChatByRow">
          <el-table-column prop="name" lable="用户名" min-width="100%" align="center"/>
        </el-table>
      </el-aside>

      <el-main class="el-main">
        <div class="chat-content" id="chat-content">

          <div v-for="(chat, index) in chatList" :key="index">
            <!-- 对方 -->
            <div class="your-word" v-if="chat.recipient_id === student.id">
              <div class="your-info">
                <div class="your-info-content">{{ chat.chat_content }}</div>
                <div class="your-time">
                  {{ chat.created_at }}
                </div>
              </div>
            </div>

            <!-- 我的 -->
            <div class="my-word" v-else>
              <div class="my-info">
                <div class="my-info-content">{{ chat.chat_content }}</div>
                <div class="my-time">
                  {{ chat.created_at }}
                </div>
              </div>
            </div>
          </div>

        </div>
        <el-input
            class="el-input"
            v-model="form.chat_content"
            :autosize="{ minRows: 5, maxRows: 10 }"
            type="textarea"
            placeholder="请输入"
            clearable
        />
        <el-button class="el-button" type="success" @click="sendMessageTest">发送</el-button>
      </el-main>
    </el-container>
  </div>
</template>

<script setup>

import {nextTick, onMounted, onUpdated, reactive, ref} from "vue";
import instance from "@/axios";
import {ElMessage} from "element-plus";
import {JSEncrypt} from 'jsencrypt';

let friendList = ref([])
let serverPublicKey = ref()
let chatList = ref([])
let student = reactive({
  id: '',
  name: '',
})

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

onMounted(() => {
  getFriend()
  getPublicKey()
  getStudent()
})

const getFriend = async () => {
  instance.get('http://localhost:8080/index/mutualFriend').then(res => {
    if (res.data.code === 200) {
      for (const datum of res.data.data) {
        friendList.value.push(datum)
      }
    } else {
      ElMessage.error('无法获取你的信息')
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

const getStudent = async () => {
  instance.get('http://localhost:8080/index/getStudentInfo').then(res => {
    if (res.data.code === 200) {
      student.id = res.data.data.id
      student.name = res.data.data.name
    } else {
      ElMessage.error('无法获取好友')
    }
  })
}

const getChatByRow = async (row) => {
  form.recipient_id = row.id
  instance.post('http://localhost:8080/index/getChat', {
    recipient_id: row.id,
    public_key: PUBLIC_KEY,
  }).then(async res => {
    if (res.data.code === 200) {
      //  清空聊天记录
      chatList.value = []

      //  装填聊天记录
      for (const datum of res.data.data) {
        datum.chat_content = RSADecrypt(datum.chat_content)
        chatList.value.push(datum)
      }
    } else {
      chatList.value = []
    }

    //  跳到聊天页面最下边
    await nextTick()
    const div = document.getElementById('chat-content');
    div.scrollTop = div.scrollHeight
  })
}

const getChatByID = async (id) => {
  form.recipient_id = id
  instance.post('http://localhost:8080/index/getChat', {
    recipient_id: id,
    public_key: PUBLIC_KEY,
  }).then(async res => {
    if (res.data.code === 200) {
      //  清空聊天记录
      chatList.value = []

      //  装填聊天记录
      for (const datum of res.data.data) {
        datum.chat_content = RSADecrypt(datum.chat_content)
        chatList.value.push(datum)
      }
    } else {
      chatList.value = []
    }

    //  跳到聊天页面最下边
    await nextTick()
    const div = document.getElementById('chat-content');
    div.scrollTop = div.scrollHeight
  })
}

const sendMessageTest = async () => {
  let chat_content = RSAEncrypt(form.chat_content)

  instance.post('http://localhost:8080/index/chat', {
    recipient_id: parseInt(form.recipient_id),
    chat_content: chat_content,
  }).then(res => {
    if (res.data.code === 200) {
      form.chat_content = ''

      ElMessage({
        message: res.data.message,
        type: 'success',
      })
    } else {
      ElMessage.error(res.data.message)
    }

    //  刷新一下聊天记录，其实应该用websocket的，但是时间所剩无几
    getChatByID(parseInt(form.recipient_id))
  })
}

function RSAEncrypt(plainText) {
  let jse = new JSEncrypt();
  jse.setPublicKey(serverPublicKey);
  return jse.encrypt(plainText);
}

function RSADecrypt(cipherText) {
  let jse = new JSEncrypt();
  jse.setPrivateKey(PRIVATE_KEY);
  return jse.decrypt(cipherText);
}

</script>

<style scoped>
.el-container {
  width: 100%;
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

.chat-content {
  margin-top: 3%;
  margin-left: 3%;
  margin-right: 3%;
  width: 94%;
  padding: 0;
  background-color: #f5f5f5;
  height: 67%;
  overflow: auto;
}

.your-word {
  display: flex;
  margin-bottom: 60px;
}

.your-info {
  width: 47%;
  margin-left: 10px;
}

.your-time {
  font-size: 12px;
  color: rgba(51, 51, 51, 0.8);
  margin: 0;
  height: 20px;
  line-height: 20px;
  /*margin-top: -5px;*/
  /*margin-top: 5px;*/
}

.your-info-content {
  word-break: break-all;
  /*max-width: 45 %;*/
  display: inline-block;
  padding: 10px;
  font-size: 17px;
  background: #fff;
  position: relative;
  margin-top: 8px;
  /*background: #dbdbdb;*/
  border-radius: 4px;
}

/*小三角形*/
.your-info-content::before {
  position: absolute;
  left: -8px;
  top: 8px;
  content: "";
  border-right: 10px solid #fff;
  border-top: 8px solid transparent;
  border-bottom: 8px solid transparent;
}


.my-word {
  display: flex;
  justify-content: flex-end;
  margin-bottom: 60px;
}

.my-info {
  width: 90%;
  /*margin-left: 10 px;*/
  text-align: right;
  /*position: relative;*/
  display: flex;
  align-items: flex-end;
  flex-wrap: wrap;
  flex-direction: column;
}


.my-info-content {
  word-break: break-all;
  max-width: 45%;
  padding: 10px;
  font-size: 17px;
  /*float: right;*/
  margin-right: 10px;
  position: relative;
  margin-top: 8px;
  background: #95ec69;
  text-align: left;
  border-radius: 4px;
}

.my-time {
  padding-right: 12px;
  padding-top: 5px;
  font-size: 12px;
  color: rgba(51, 51, 51, 0.8);
  margin: 0;
  height: 20px;
}


.my-info-content::after {
  position: absolute;
  right: -8px;
  top: 8px;
  content: "";
  border-left: 10px solid #95ec69;
  border-top: 8px solid transparent;
  border-bottom: 8px solid transparent;
}

.el-input {
  margin-left: 3%;
  margin-right: 3%;
  width: 94%;
}

.el-button {
  margin-left: 87%;
  width: 10%;
  height: 5%;
}
</style>