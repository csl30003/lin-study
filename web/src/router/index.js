import {createRouter, createWebHistory} from 'vue-router'

const routes = [
    {
        path: '/',
        component: () => import('../views/index')
    },
    {
        path: '/register',
        component: () => import('../views/register')
    },
    {
        path: '/home',
        component: () => import('../views/home'),
        children: [
            {
                path: 'map',
                component: () => import('../views/floor'),
            },
            {
                path: 'map/:floor',
                component: () => import('../views/layer')
            },
            {
                path: 'map/:floor/:layer',
                component: () => import('../views/class')
            },
            {
                path: 'map/:floor/:layer/:class',
                component: () => import('../views/seat')
            },
            {
                path: 'chat',
                component: () => import('../views/chat')
            }
        ]
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router
