<script setup lang="ts">
import { computed } from 'vue'
import { RouterView, useRoute } from 'vue-router'

import Sidebar from '../components/layout/Sidebar.vue'

const route = useRoute()

const pageTitle = computed(() => route.meta.title ?? 'Dynamic DB Admin')
const pageDescription = computed(() => route.meta.description ?? '')
</script>

<template>
    <div class="flex h-screen overflow-hidden bg-[#06070f] text-gray-100">
        <Sidebar />

        <div class="flex min-w-0 flex-1 flex-col">
            <header class="flex h-16 shrink-0 items-center justify-between border-b border-white/10 px-8">
                <div>
                    <h2 class="text-lg font-semibold text-white">
                        {{ pageTitle }}
                    </h2>

                    <p v-if="pageDescription" class="text-xs text-gray-500">
                        {{ pageDescription }}
                    </p>
                </div>

                <div class="flex items-center gap-2 text-xs text-gray-400">
                    <span class="h-2 w-2 rounded-full bg-gray-500" />
                    База данных не открыта
                </div>
            </header>

            <main class="min-h-0 flex-1 overflow-y-auto p-8">
                <div class="mx-auto w-full max-w-7xl">
                    <RouterView />
                </div>
            </main>
        </div>
    </div>
</template>
