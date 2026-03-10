import { ref } from 'vue'
import { defineStore } from 'pinia'
import { ReadLog } from '../../wailsjs/go/pages/TTS'

export const useLogStore = defineStore('log', () => {
  const content = ref('')
  const isLoading = ref(false)

  async function fetchLog() {
    isLoading.value = true
    try {
      content.value = await ReadLog()
    } finally {
      isLoading.value = false
    }
  }

  return { content, isLoading, fetchLog }
})
