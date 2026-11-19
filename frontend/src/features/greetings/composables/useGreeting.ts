import { ref } from "vue";
import { Greet } from "../../../../bindings/artemiy6451/dynamic-db-admin/internal/adapters/wails/appservice";


export function useGreeting() {
    const greeting = ref('')
    const loading = ref(false)
    const error = ref<string | null>(null)

    async function greet(name: string) {
        loading.value = true
        error.value = null

        try {
            greeting.value = await Greet(name)
        } catch (err) {
            error.value =
                err instanceof Error
                    ? err.message
                    : 'Failed to call backend'
        } finally {
            loading.value = false
        }
    }

    return {
        greeting,
        loading,
        error,
        greet,
    }
}
