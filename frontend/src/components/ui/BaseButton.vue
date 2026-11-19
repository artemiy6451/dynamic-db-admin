<script setup lang="ts">
import { LoaderCircle } from 'lucide-vue-next'

withDefaults(
    defineProps<{
        variant?: 'primary' | 'secondary' | 'danger' | 'ghost'
        size?: 'sm' | 'md' | 'lg'
        type?: 'button' | 'submit' | 'reset'
        loading?: boolean
        disabled?: boolean
    }>(),
    {
        variant: 'primary',
        size: 'md',
        type: 'button',
        loading: false,
        disabled: false,
    },
)
</script>

<template>
    <button :type="type" :disabled="disabled || loading"
        class="inline-flex items-center justify-center gap-2 rounded-lg font-medium transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-emerald-500/60 disabled:pointer-events-none disabled:opacity-50"
        :class="[
            {
                'bg-emerald-500 text-white hover:bg-emerald-400': variant === 'primary',
                'border border-white/10 bg-white/5 text-gray-200 hover:bg-white/10': variant === 'secondary',
                'bg-red-500/10 text-red-400 hover:bg-red-500/20': variant === 'danger',
                'text-gray-400 hover:bg-white/5 hover:text-white': variant === 'ghost',
            },
            {
                'px-3 py-1.5 text-xs': size === 'sm',
                'px-4 py-2 text-sm': size === 'md',
                'px-5 py-2.5 text-sm': size === 'lg',
            },
        ]">
        <LoaderCircle v-if="loading" :size="16" class="animate-spin" />

        <slot />
    </button>
</template>
