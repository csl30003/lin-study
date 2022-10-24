import {createRouter, createWebHistory} from "vue-router";

import index from './views/index.vue'
import login from './views/login.vue'
import register from './views/register.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            component: index,
        },
        {
            path: '/login',
            component: login
        },
        {
            path: '/register',
            component: register
        }
    ]
})
export default router