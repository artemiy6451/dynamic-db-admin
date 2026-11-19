<script setup lang="ts">
import { ref } from 'vue'
import { useGreeting } from '../composables/useGreeting';

const name = ref('Lokach')

const {
    greeting,
    loading,
    error,
    greet,
} = useGreeting()
</script>

<template>
    <section>
        <h1>Backend communication</h1>

        <input v-model="name" placeholder="Enter your name" />

        <button :disabled="loading" @click="greet(name)">
            {{ loading ? 'Loading...' : 'Greet' }}
        </button>

        <p v-if="greeting">{{ greeting }}</p>
        <p v-if="error">{{ error }}</p>
    </section>
</template>
