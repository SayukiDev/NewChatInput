<script setup lang="ts">
import {onUnmounted } from 'vue'
import { useToast } from 'primevue/usetoast'
import { useI18n } from 'vue-i18n'
import { useChatStore } from '../stores/chat'
import MessageInput from '../components/chat/MessageInput.vue'
import { SetTyping } from '../../wailsjs/go/pages/Input'
import {WindowGetSize, WindowSetSize} from "../../wailsjs/runtime";

const chatStore = useChatStore()
const toast = useToast()
const { t } = useI18n()

function enterFullInput() {
  chatStore.toggleFullInput(true).then(()=>{
    WindowGetSize().then(size => {
      WindowSetSize(size.w,size.h-1)
      toast.add({
        severity: 'info',
        summary: t('chat.fullInput.toastSummary'),
        detail: t('chat.fullInput.toastDetail'),
        life: 3000,
      })
  })
  })
}

function exitFullInput() {
  chatStore.toggleFullInput(false).then(()=>{
    WindowGetSize().then(size => {
      WindowSetSize(size.w,size.h+1)
    })
  })
}

onUnmounted(async () => {
  try {
    await SetTyping(false)
  } catch {
    // ignore
  }
})
</script>

<template>
  <!-- Normal mode -->
  <div class="chat-view">
    <Card v-if="!chatStore.isFullInput" >
      <template #title >{{ $t('chat.title') }}</template>
      <template #content>
        <MessageInput
          v-model="chatStore.message"
          :isSending="chatStore.isSending"
          :isFullInput="chatStore.isFullInput"
          @send="chatStore.sendMessage"
          @reset="chatStore.reset"
          @enterFullInput="enterFullInput"
          @exitFullInput="exitFullInput"
        />
      </template>
    </Card>
    <MessageInput
        v-else
        v-model="chatStore.message"
        :isSending="chatStore.isSending"
        :isFullInput="chatStore.isFullInput"
        @send="chatStore.sendMessage"
        @reset="chatStore.reset"
        @enterFullInput="enterFullInput"
        @exitFullInput="exitFullInput"
    />
  </div>
</template>

<style scoped>
.chat-view {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.chat-view :deep(.p-card) {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
}

.chat-view :deep(.p-card-body) {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.chat-view :deep(.p-card-content) {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.full-input-overlay {
  position: fixed;
  inset: 0;
  z-index: 9999;
  display: flex;
  flex-direction: column;
  background-color: var(--p-surface-ground);
}
</style>
