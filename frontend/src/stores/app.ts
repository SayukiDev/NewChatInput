import { defineStore } from 'pinia'
import { ref } from 'vue'

// App store for managing application state
export const useAppStore = defineStore('app', () => {
  // Controls whether full input mode is enabled
  const isFullInputMode = ref(false)

  // Toggle the input mode
  const toggleInputMode = () => {
    isFullInputMode.value = !isFullInputMode.value
  }

  // Set input mode to a specific value
  const setInputMode = (value: boolean) => {
    isFullInputMode.value = value
  }

  return {
    isFullInputMode,
    toggleInputMode,
    setInputMode
  }
})