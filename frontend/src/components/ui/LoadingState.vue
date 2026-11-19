<script setup lang="ts">
import { LoaderCircle } from 'lucide-vue-next'

withDefaults(
    defineProps<{
        label?: string
        size?: 'sm' | 'md' | 'lg'
    }>(),
    {
        label: 'Загрузка...',
        size: 'md',
    },
)
</script>

<template>
    <div class="flex items-center justify-center gap-3 text-gray-400" role="status" aria-live="polite">
        <LoaderCircle class="animate-spin text-emerald-400" :size="size === 'sm' ? 16 : size === 'lg' ? 28 : 22" />

        <span class="text-sm">
            {{ label }}
        </span>
    </div>
</template>
