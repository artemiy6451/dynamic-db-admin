<template>
    <h1 class="title">{{ appTitle }}</h1>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { GetConfig } from "../../bindings/artemiy6451/dynamic-db-admin/internal/config/configservice"

const appTitle = ref("")

onMounted(async () => {
    const config = await GetConfig()

    appTitle.value = config.App.Window.Title

})

</script>

<style scoped>
</style>
