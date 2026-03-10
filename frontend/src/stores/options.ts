import { ref } from 'vue'
import { defineStore } from 'pinia'
import { Load, Save } from '../../wailsjs/go/pages/Options'
import { options } from '../../wailsjs/go/models'
import {useTTSStore} from "./tts.ts";

export const useOptionsStore = defineStore('options', () => {
  const config = ref<options.Config | null>(null)
  const isLoading = ref(false)

  async function loadConfig() {
    isLoading.value = true
    try {
      config.value = await Load()
    } finally {
      isLoading.value = false
    }
  }

  async function saveConfig() {
    if (!config.value) return
    let ts=useTTSStore()
    config.value.tts_option.device=ts.selectedDeviceId
    config.value.tts_option.now_spacker=ts.currentSpeakerId
    return Save(config.value)
  }

  return { config, isLoading, loadConfig, saveConfig }
})
