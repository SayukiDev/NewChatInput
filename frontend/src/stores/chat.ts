import { ref, watch } from 'vue'
import { defineStore } from 'pinia'
import { watchDebounced } from '@vueuse/core'
import {SendMessage, SetFullInputMode, SetTyping} from '../../wailsjs/go/pages/Input'

export const MAX_LENGTH = 144

export const useChatStore = defineStore('chat', () => {
  const message = ref('')
  const isSending = ref(false)
  const isTyping = ref(false)
  const isFullInput = ref(false)

  async function sendMessage() {
    const text = message.value.trim()
    if (isSending.value) return
    isSending.value = true
    try {
      message.value = ''
      await SendMessage(text)
    } finally {
      isSending.value = false
    }
  }

  async function reset() {
    if (isSending.value)
      isSending.value = false
    if (message.value!='') {
      message.value = ''
    }
    else{
      await sendMessage()
    }
  }

  async function toggleFullInput(enter : boolean = false) {
    await SetFullInputMode(enter)
    isFullInput.value = enter
  }

  // Typing indicator: notify backend when user is typing, debounce to avoid spam
  watchDebounced(
    message,
    async (val) => {
      const typing = val.length > 0
      if (typing !== isTyping.value) {
        isTyping.value = typing
        try {
          await SetTyping(typing)
        } catch {
          // ignore
        }
      }
    },
    { debounce: 300 },
  )

  // Auto-stop typing after 2 seconds of no input changes
  let typingTimeout: ReturnType<typeof setTimeout> | null = null
  watch(message, () => {
    if (typingTimeout) clearTimeout(typingTimeout)
    typingTimeout = setTimeout(async () => {
      if (isTyping.value) {
        isTyping.value = false
        try {
          await SetTyping(false)
        } catch {
          // ignore
        }
      }
    }, 2000)
  })

  return { message, isSending, sendMessage, reset,isFullInput, toggleFullInput }
})
