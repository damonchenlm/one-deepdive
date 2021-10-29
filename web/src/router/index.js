import {createRouter, createWebHashHistory} from 'vue-router'

const routes = [{
    path: '/',
    component: () => import('@/views/home/index')
},
    {
        path: '/news',
        component: () => import('@/views/news/index')
    },
    {
        path: '/about',
        component: () => import('@/views/about/index')
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router