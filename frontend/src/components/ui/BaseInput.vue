<script setup lang="ts">
import { computed } from 'vue'

const props = withDefaults(
    defineProps<{
        modelValue: string
        label?: string
        placeholder?: string
        type?: string
        error?: string
        disabled?: boolean
        id?: string
    }>(),
    {
        label: '',
        placeholder: '',
        type: 'text',
        error: '',
        disabled: false,
        id: undefined,
    },
)

const emit = defineEmits<{
    'update:modelValue': [value: string]
}>()

const inputId = computed(
    () => props.id ?? `input-${Math.random().toString(36).slice(2, 9)}`,
)
</script>

<template>
    <div class="w-full space-y-2">
        <label v-if="label" :for="inputId" class="block text-sm font-medium text-gray-300">
            {{ label }}
        </label>

        <input :id="inputId" :type="type" :value="modelValue" :placeholder="placeholder" :disabled="disabled"
            :aria-invalid="Boolean(error)" :aria-describedby="error ? `${inputId}-error` : undefined"
            class="w-full rounded-lg border bg-[#0b0d17] px-3 py-2.5 text-sm text-white outline-none transition-colors placeholder:text-gray-600 focus:ring-2 disabled:cursor-not-allowed disabled:opacity-50"
            :class="error
                    ? 'border-red-500/50 focus:border-red-500 focus:ring-red-500/20'
                    : 'border-white/10 focus:border-emerald-500/60 focus:ring-emerald-500/20'
                " @input="emit('update:modelValue', ($event.target as HTMLInputElement).value)" />

        <p v-if="error" :id="`${inputId}-error`" class="text-xs text-red-400">
            {{ error }}
        </p>
    </div>
</template>
