import { ref } from 'vue'
import { defineStore } from 'pinia'
import { IsLoaded } from '../../wailsjs/go/pages/App'

export const useAppStore = defineStore('app', () => {
  const isReady = ref(false)
  let timer: ReturnType<typeof setInterval> | null = null

  async function startPolling() {
    if (isReady.value) return
    timer = setInterval(async () => {
      try {
        const loaded = await IsLoaded()
        if (loaded) {
          isReady.value = true
          stopPolling()
        }
      } catch {
        // backend not ready yet
      }
    }, 500)
  }

  function stopPolling() {
    if (timer) {
      clearInterval(timer)
      timer = null
    }
  }

  return { isReady, startPolling, stopPolling }
})
