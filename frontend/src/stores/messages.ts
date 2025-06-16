import { defineStore } from 'pinia'
import { ref } from 'vue'

export interface Message {
  id: string
  text: string
  color?: string
  timeout?: number
  visible: boolean
}

export const useMessagesStore = defineStore('messages', () => {
  const messages = ref<Message[]>([])

  const addMessage = (message: Omit<Message, 'id' | 'visible'>) => {
    const newMessage: Message = {
      ...message,
      id: Date.now().toString() + Math.random().toString(36).slice(2, 11),
      visible: true
    }
    messages.value.push(newMessage)
  }

  const removeMessage = (id: string) => {
    const index = messages.value.findIndex(msg => msg.id === id)
    if (index > -1) {
      messages.value.splice(index, 1)
    }
  }

  const addError = (text: string) => {
    addMessage({
      text,
      color: 'error',
      timeout: 6000
    })
  }

  const addWarning = (text: string) => {
    addMessage({
      text,
      color: 'warning',
      timeout: 5000
    })
  }

  const addInfo = (text: string) => {
    addMessage({
      text,
      color: 'info',
      timeout: 4000
    })
  }

  const addSuccess = (text: string) => {
    addMessage({
      text,
      color: 'success',
      timeout: 4000
    })
  }

  const clearMessages = () => {
    messages.value = []
  }

  return {
    messages,
    addMessage,
    removeMessage,
    addError,
    addWarning,
    addInfo,
    addSuccess,
    clearMessages
  }
})