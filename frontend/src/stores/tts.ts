import { ref } from 'vue'
import { defineStore } from 'pinia'
import {
  GetAudioDevice,
  GetSpackers,
  InstallModel,
  SaveAudioDevice,
  SaveSpacker,
  SelectedSpacker,
} from '../../wailsjs/go/pages/TTS'
import type { aivis, tts } from '../../wailsjs/go/models'
import {useOptionsStore} from "./options.ts";
import {IsRunning} from "../../wailsjs/go/pages/TTS";

export const useTTSStore = defineStore('tts', () => {
  const speakers = ref<aivis.Speaker[]>([])
  const currentSpeakerId = ref<number>(-1)
  const styleNames = ref<string[]>([])
  const devices = ref<tts.Device[]>([])
  const selectedDeviceId = ref<string>('')
  const isRunning = ref(false)
  const isInstalling = ref(false)

  async function checkIsRunning() {
    try {
      isRunning.value = await IsRunning()
    } catch {
      isRunning.value = false
    }
  }

  async function loadSpeakers() {
    speakers.value = await GetSpackers()
    console.log('[tts] loadSpeakers:', speakers.value.length)
  }

  async function loadCurrentSpeaker() {
    const rsp = await SelectedSpacker()
    console.log('[tts] loadCurrentSpeaker rsp:', rsp)
    currentSpeakerId.value = rsp.SpackerId
    const speaker = speakers.value.find(s => s.styles.some(st => st.id === rsp.SpackerId))
    styleNames.value = speaker ? speaker.styles.map(s => s.name) : []
  }

  async function loadDevices(currentDeviceId: string) {
    devices.value = await GetAudioDevice()
    console.log('[tts] loadDevices:', devices.value.length, 'selected:', currentDeviceId)
    selectedDeviceId.value = currentDeviceId
  }

  async function selectSpeaker(speaker: aivis.Speaker) {
    const firstStyleId = speaker.styles[0]?.id ?? 0
    await SaveSpacker(firstStyleId)
    currentSpeakerId.value = firstStyleId
    styleNames.value = speaker.styles.map(s => s.name)
    await useOptionsStore().loadConfig()
  }

  async function selectStyle(styleId: number) {
    await SaveSpacker(styleId)
    currentSpeakerId.value = styleId
    await useOptionsStore().loadConfig()
  }

  async function selectDevice(id: string) {
    await SaveAudioDevice(id)
    selectedDeviceId.value = id
  }

  async function installModel() {
    isInstalling.value = true
    try {
      await InstallModel()
      await loadSpeakers()
    } finally {
      isInstalling.value = false
    }
  }

  return {
    speakers,
    currentSpeakerId,
    styleNames,
    devices,
    selectedDeviceId,
    isRunning,
    isInstalling,
    loadSpeakers,
    loadCurrentSpeaker,
    loadDevices,
    selectSpeaker,
    selectStyle,
    selectDevice,
    installModel,
    checkIsRunning,
  }
})
