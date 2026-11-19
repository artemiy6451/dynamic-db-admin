import { createRouter, createWebHashHistory } from 'vue-router'

import MainLayout from '../layouts/MainLayout.vue'

const router = createRouter({
    history: createWebHashHistory(),

    routes: [
        {
            path: '/',
            component: MainLayout,
            children: [
                {
                    path: '',
                    name: 'Dashboard',
                    component: () => import('../pages/DashboardPage.vue'),
                    meta: {
                        title: 'Главная',
                        description: 'Обзор приложения',
                    },
                },
                {
                    path: 'connections',
                    name: 'Connections',
                    component: () => import('../pages/ConnectionsPage.vue'),
                    meta: {
                        title: 'Подключения',
                        description: 'Управление подключениями PostgreSQL',
                    },
                },
                {
                    path: 'settings',
                    name: 'Settings',
                    component: () => import('../pages/SettingsPage.vue'),
                    meta: {
                        title: 'Настройки',
                        description: 'Настройки приложения',
                    },
                },
            ],
        },
    ],
})

export default router
