<script setup lang="ts">
import { Database } from 'lucide-vue-next'
</script>

<template>
  <section class="space-y-6">
    <div>
      <h1 class="text-xl font-semibold text-white">
        Подключения
      </h1>

      <p class="mt-2 text-sm text-gray-400">
        Управление подключениями к серверам PostgreSQL.
      </p>
    </div>

    <div class="rounded-xl border border-white/10 bg-[#0b0d17] px-6 py-12 text-center">
      <Database
        :size="32"
        class="mx-auto text-gray-600"
      />

      <h2 class="mt-4 text-sm font-medium text-gray-300">
        Нет сохранённых подключений
      </h2>

      <p class="mt-2 text-xs text-gray-500">
        Добавленные подключения будут отображаться здесь.
      </p>
    </div>
  </section>
</template>
