import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import axios from 'axios'
import VueAxios from 'vue-axios'

import App from './App.vue'
import router from "./router"

const app = createApp(App)

app.use(ElementPlus)
app.use(VueAxios)
app.use(router)
app.mount('#app')

// D:
// cd goProjects\lin_study\web
// vue serve