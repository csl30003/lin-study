import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import VueAxios from 'vue-axios'
import installElementPlus from './plugins/element'

const app = createApp(App)
app.use(installElementPlus)
app.use(VueAxios)
app.use(router)
app.mount('#app')
