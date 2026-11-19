<script setup lang="ts">
import { useRouter } from 'vue-router'
import {
    Database,
    Table2,
    Clock3,
    Plus,
    ArrowUpRight,
    Activity,
} from 'lucide-vue-next'

import BaseButton from '../components/ui/BaseButton.vue'

const router = useRouter()

const statistics = [
    {
        label: 'Открытые базы',
        value: '0',
        description: 'SQLite-файлов открыто сейчас',
        icon: Database,
    },
    {
        label: 'Таблицы',
        value: '0',
        description: 'В открытых базах данных',
        icon: Table2,
    },
    {
        label: 'Недавние базы',
        value: '0',
        description: 'Файлов в истории приложения',
        icon: Clock3,
    },
]

function openDatabases() {
    router.push({ name: 'Databases' })
}
</script>

<template>
    <div class="space-y-8">
        <section>
            <h1 class="text-2xl font-semibold text-white">
                Добро пожаловать!
            </h1>

            <p class="mt-2 text-sm text-gray-400">
                Создавайте и открывайте SQLite-базы, изучайте их структуру
                и управляйте данными в одном приложении.
            </p>
        </section>

        <section class="grid grid-cols-1 gap-4 md:grid-cols-3">
            <div
                v-for="item in statistics"
                :key="item.label"
                class="rounded-xl border border-white/10 bg-[#0b0d17] p-5"
            >
                <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-400">
                        {{ item.label }}
                    </span>

                    <component :is="item.icon" :size="19" class="text-emerald-400" />
                </div>

                <p class="mt-5 text-3xl font-semibold text-white">
                    {{ item.value }}
                </p>

                <p class="mt-1 text-xs text-gray-500">
                    {{ item.description }}
                </p>
            </div>
        </section>

        <section class="rounded-xl border border-white/10 bg-[#0b0d17] p-6">
            <div class="flex flex-col gap-5 sm:flex-row sm:items-center sm:justify-between">
                <div>
                    <h2 class="text-base font-semibold text-white">
                        Начало работы
                    </h2>

                    <p class="mt-2 max-w-xl text-sm text-gray-400">
                        Откройте существующий файл SQLite или создайте новую базу.
                        После этого здесь появятся доступные инструменты работы
                        со структурой и данными.
                    </p>
                </div>

                <BaseButton @click="openDatabases">
                    <Plus :size="17" />
                    Открыть базу данных
                    <ArrowUpRight :size="15" />
                </BaseButton>
            </div>
        </section>

        <section>
            <h2 class="mb-4 text-base font-semibold text-white">
                Последняя активность
            </h2>

            <div class="rounded-xl border border-white/10 bg-[#0b0d17] p-8 text-center">
                <Activity :size="28" class="mx-auto text-gray-600" />

                <p class="mt-3 text-sm font-medium text-gray-300">
                    Пока нет активности
                </p>

                <p class="mt-1 text-xs text-gray-500">
                    Здесь будут отображаться последние действия с базами данных.
                </p>
            </div>
        </section>
    </div>
</template>
