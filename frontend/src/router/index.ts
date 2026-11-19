import { createRouter, createWebHashHistory } from 'vue-router'

import DashboardPage from '../pages/DashboardPage.vue'
import ConnectionsPage from '../pages/ConnectionsPage.vue'
import SettingsPage from '../pages/SettingsPage.vue'

const router = createRouter({
    history: createWebHashHistory(),

    routes: [
        {
            path: '/',
            name: 'dashboard',
            component: DashboardPage,
        },
        {
            path: '/connections',
            name: 'connections',
            component: ConnectionsPage,
        },
        {
            path: '/settings',
            name: 'settings',
            component: SettingsPage,
        },
    ],
})

export default router
