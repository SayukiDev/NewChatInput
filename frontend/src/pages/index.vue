<template>
  <v-container class="px-0">
    <!-- Typing Card -->
    <v-card
        @keydown="handlePageKeyDown"
        height="100%"
        width="100%"
        class="mb-4"
    >
      <v-card-title class="pb-2" v-if="!appStore.isFullInputMode">
        <v-icon class="mr-2">mdi-keyboard</v-icon>
        Typing
      </v-card-title>
      <v-card-text>
        <v-textarea
          width="100%"
          ref="inputRef"
          v-model="inputText"
          label="Enter your message"
          variant="outlined"
          clearable
          autofocus
          :rows="appStore.isFullInputMode ? 10 : 5"
          no-resize
          density="comfortable"
          placeholder="Type your message here..."
          @keydown="handleInputKeydown"
        />
      </v-card-text>
    </v-card>
    
    <!-- Actions Card -->
    <v-card v-if="!appStore.isFullInputMode">
      <v-card-title class="pb-2">
        <v-icon class="mr-2">mdi-gesture-tap-button</v-icon>
        Actions
      </v-card-title>
      <v-card-text>
        <v-row>
          <v-col cols="6">
            <v-btn
              color="primary"
              variant="elevated"
              block
              size="large"
              @click="handleSend"
              :disabled="!InputTextNotNull"
            >
              <v-icon left>mdi-send</v-icon>
              Send
            </v-btn>
          </v-col>
          
          <v-col cols="6">
            <v-btn
              color="warning"
              variant="outlined"
              block
              size="large"
              @click="handleClear"
            >
              <v-icon left>mdi-eraser</v-icon>
              Clear
            </v-btn>
          </v-col>
        </v-row>

        <v-row class="mt-2">
          <v-col cols="12">
            <v-btn
              color="info"
              variant="tonal"
              block
              size="large"
              @click="enterFullInputMode"
              :disabled="appStore.isFullInputMode"
            >
              <v-icon class="mr-2">mdi-fullscreen</v-icon>
              Full Input Mode
            </v-btn>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script setup lang="ts">
import { ref, computed,useTemplateRef } from 'vue'
import { useMessagesStore } from '@/stores/messages';
import { useAppStore } from '@/stores/app';
import {SendMessage, SetFullInputMode} from "../../wailsjs/go/pages/Input";

const messagesStore = useMessagesStore()
const appStore = useAppStore()
const inputText = ref('')
const inputRef=ref<HTMLElement>()

const InputTextNotNull=computed(function() {
  return inputText?.value!==""
})

const handleSend = () => {
  if (!inputText?.value.trim()) {
    messagesStore.addWarning('Please enter a message')
    return
  }

  SendMessage(inputText.value).then(() => {
  }).catch((error) => {
    messagesStore.addError(`Failed to send message: ${error.message}`)
  }).finally(()=>{
    messagesStore.addSuccess('Message sent successfully')
  })
  inputText.value = ''
}

const handleClear = () => {
  if (inputText?.value.trim()) {
    inputText.value = ''
    messagesStore.addSuccess('Input cleared')
  }else{
    SendMessage('').then(()=>{
      messagesStore.addInfo('ChatBox cleared')
    }).catch((error) => {
      messagesStore.addError(`Failed to clear input: ${error.message}`)
    })
  }
}

const handleInputKeydown = (event: KeyboardEvent) => {
  switch (event.key) {
    case 'Enter':
      if (!event.shiftKey) {
        handleSend()
      }
      break
    case 'Delete':
      handleClear()
      break
    default:
      return
  }
  event.preventDefault()
}

function handlePageKeyDown(event: KeyboardEvent) {
  if (event.key === 'Escape'&&appStore.isFullInputMode) {
    appStore.setInputMode(false)
    SetFullInputMode(false)
    focusInput()
  }
}

// Function to enter full input mode
const enterFullInputMode = () => {
  appStore.setInputMode(true)
  SetFullInputMode(true).then(()=>{
    messagesStore.addInfo('Entered full input mode - press ESC to exit')
    focusInput()
  })
}

function focusInput(){
  (inputRef.value as HTMLElement)?.focus()
}
</script>