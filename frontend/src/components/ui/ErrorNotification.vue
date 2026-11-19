<script setup lang="ts">
import { AlertCircle, X } from 'lucide-vue-next'

import BaseButton from './BaseButton.vue';

withDefaults(
    defineProps<{
        title?: string
        message: string
        dismissible?: boolean
    }>(),
    {
        title: 'Произошла ошибка',
        dismissible: true,
    },
)

const emit = defineEmits<{
    dismiss: []
}>()
</script>

<template>
    <div class="flex gap-3 rounded-lg border border-red-500/20 bg-red-500/10 p-4" role="alert">
        <AlertCircle :size="20" class="mt-0.5 shrink-0 text-red-400" />

        <div class="min-w-0 flex-1">
            <p class="text-sm font-medium text-red-300">
                {{ title }}
            </p>

            <p class="mt-1 break-words text-sm text-red-400/80">
                {{ message }}
            </p>
        </div>

        <BaseButton v-if="dismissible" variant="ghost" size="sm" aria-label="Закрыть уведомление"
            @click="emit('dismiss')">
            <X :size="16" />
        </BaseButton>
    </div>
</template>
