<script setup lang="ts">
import {
    LayoutDashboard,
    Database,
    Settings,
} from 'lucide-vue-next'

const navigation = [
    {
        name: 'Dashboard',
        label: 'Главная',
        path: '/',
        icon: LayoutDashboard,
    },
    {
        name: 'Databases',
        label: 'Базы данных',
        path: '/databases',
        icon: Database,
    },
    {
        name: 'Settings',
        label: 'Настройки',
        path: '/settings',
        icon: Settings,
    },
]
</script>

<template>
    <aside class="flex h-screen w-64 shrink-0 flex-col border-r border-white/10 bg-[#0b0d17]">
        <div class="flex h-16 items-center gap-3 border-b border-white/10 px-5">
            <div class="flex h-9 w-9 items-center justify-center rounded-xl bg-emerald-500/15">
                <Database :size="20" class="text-emerald-400" />
            </div>

            <div class="min-w-0">
                <h1 class="text-sm font-semibold text-white">
                    Dynamic DB Admin
                </h1>
                <p class="text-xs text-gray-500">
                    SQLite Manager
                </p>
            </div>
        </div>

        <nav class="flex-1 space-y-1 px-3 py-5">
            <p class="mb-3 px-3 text-xs font-medium uppercase tracking-wider text-gray-500">
                Workspace
            </p>

            <RouterLink
                v-for="item in navigation"
                :key="item.name"
                :to="item.path"
                active-class="bg-emerald-500/10 text-emerald-400"
                class="flex items-center gap-3 rounded-lg px-3 py-2.5 text-sm text-gray-400 transition-colors hover:bg-white/5 hover:text-white"
            >
                <component :is="item.icon" :size="18" />
                <span>{{ item.label }}</span>
            </RouterLink>
        </nav>

        <div class="border-t border-white/10 p-4">
            <div class="flex items-center justify-between">
                <div>
                    <p class="text-xs font-medium text-gray-300">
                        Dynamic DB Admin
                    </p>
                    <p class="mt-1 text-xs text-gray-500">
                        Early Development
                    </p>
                </div>

                <span class="rounded-md border border-white/10 px-2 py-1 text-[10px] text-gray-500">
                    v0.1.0
                </span>
            </div>
        </div>
    </aside>
</template>
