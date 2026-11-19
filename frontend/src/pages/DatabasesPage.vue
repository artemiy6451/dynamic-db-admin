<script setup lang="ts">
import { Database, FolderOpen, Plus, HardDrive } from 'lucide-vue-next'
</script>

<template>
    <div class="space-y-8">
        <section>
            <h1 class="text-2xl font-semibold text-white">
                Базы данных
            </h1>
            <p class="mt-2 text-sm text-gray-400">
                Открывайте SQLite-файлы и управляйте базами данных в рамках одного рабочего пространства.
            </p>
        </section>

        <section class="rounded-xl border border-white/10 bg-[#0b0d17] px-6 py-14 text-center">
            <div class="mx-auto flex h-14 w-14 items-center justify-center rounded-2xl bg-emerald-500/10">
                <Database :size="28" class="text-emerald-400" />
            </div>

            <h2 class="mt-5 text-base font-semibold text-white">
                Пока нет открытых баз данных
            </h2>

            <p class="mx-auto mt-2 max-w-lg text-sm leading-6 text-gray-400">
                Здесь будут отображаться открытые SQLite-файлы. Вы сможете переключаться между базами,
                просматривать их структуру и редактировать записи.
            </p>

            <div class="mt-6 flex flex-wrap justify-center gap-3">
                <div class="inline-flex items-center gap-2 rounded-lg border border-white/10 bg-white/[0.03] px-4 py-2.5 text-sm text-gray-400">
                    <FolderOpen :size="16" />
                    Открытие файла
                </div>
                <div class="inline-flex items-center gap-2 rounded-lg border border-white/10 bg-white/[0.03] px-4 py-2.5 text-sm text-gray-400">
                    <Plus :size="16" />
                    Создание базы
                </div>
            </div>

            <p class="mt-6 text-xs text-gray-500">
                Управление SQLite-файлами будет добавлено в Sprint 2.
            </p>
        </section>

        <section class="grid gap-4 md:grid-cols-2">
            <div class="rounded-xl border border-white/10 bg-[#0b0d17] p-5">
                <div class="flex items-center gap-3">
                    <HardDrive :size="18" class="text-emerald-400" />
                    <h3 class="text-sm font-medium text-gray-200">Локальные файлы</h3>
                </div>
                <p class="mt-2 text-sm text-gray-500">
                    Поддерживаемые расширения: .db, .sqlite и .sqlite3.
                </p>
            </div>

            <div class="rounded-xl border border-white/10 bg-[#0b0d17] p-5">
                <div class="flex items-center gap-3">
                    <Database :size="18" class="text-emerald-400" />
                    <h3 class="text-sm font-medium text-gray-200">Несколько баз</h3>
                </div>
                <p class="mt-2 text-sm text-gray-500">
                    В дальнейшем можно будет держать открытыми несколько файлов и быстро переключаться между ними.
                </p>
            </div>
        </section>
    </div>
</template>
