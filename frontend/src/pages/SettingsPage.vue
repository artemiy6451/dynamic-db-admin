<script setup lang="ts">
import { Settings } from 'lucide-vue-next'
</script>

<template>
    <section class="space-y-6">
        <div>
            <h1 class="text-xl font-semibold text-white">
                Настройки
            </h1>

            <p class="mt-2 text-sm text-gray-400">
                Настройка параметров приложения.
            </p>
        </div>

        <div class="rounded-xl border border-white/10 bg-[#0b0d17] px-6 py-12 text-center">
            <Settings :size="32" class="mx-auto text-gray-600" />

            <h2 class="mt-4 text-sm font-medium text-gray-300">
                Настройки пока недоступны
            </h2>

            <p class="mt-2 text-xs text-gray-500">
                Параметры приложения появятся на этом экране в следующих версиях.
            </p>
        </div>
    </section>
</template>
