<template>
  <div>
    <!-- Typing Card -->
    <v-card class="mb-4">
      <v-card-title class="pb-2">
        <v-icon class="mr-2">mdi-keyboard</v-icon>
        Typing
      </v-card-title>
      <v-card-text>
        <v-textarea
          v-model="inputText"
          label="Enter your message"
          variant="outlined"
          clearable
          rows="5"
          no-resize
          density="comfortable"
          placeholder="Type your message here..."
          @keydown="handleKeydown"
        />
      </v-card-text>
    </v-card>
    
    <!-- Actions Card -->
    <v-card>
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
              color="secondary"
              variant="outlined"
              block
              size="large"
              @click="handleClear"
              :disabled="!InputTextNotNull"
            >
              <v-icon left>mdi-eraser</v-icon>
              Clear
            </v-btn>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useMessagesStore } from '@/stores/messages'
import {SendMessage} from "../../wailsjs/go/pages/Input";

const messagesStore = useMessagesStore()
const inputText = ref('')

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
  inputText.value = ''
  SendMessage('').then(()=>{
    messagesStore.addInfo('Input cleared')
  }).catch((error) => {
    messagesStore.addError(`Failed to clear input: ${error.message}`)
  })
}

const handleKeydown = (event: KeyboardEvent) => {
  if (event.key === 'Enter' && !event.shiftKey) {
    event.preventDefault()
    handleSend()
  }
  // Shift+Enter allows default behavior (new line)
}
</script>

<route lang="yaml">
meta:
  layout: default
</route>