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
                        description: 'Обзор открытых баз данных',
                    },
                },
                {
                    path: 'databases',
                    name: 'Databases',
                    component: () => import('../pages/DatabasesPage.vue'),
                    meta: {
                        title: 'Базы данных',
                        description: 'Открытие и управление SQLite-файлами',
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
                {
                    path: ':pathMatch(.*)*',
                    redirect: '/',
                },
            ],
        },
    ],
})

export default router
