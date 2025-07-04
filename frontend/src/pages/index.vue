<template>
    <!-- Typing Card -->
    <v-card
        v-if="!appStore.isFullInputMode"
        class="mb-4 align-start"
    >
      <v-card-title class="pb-2">
        <v-icon class="mr-2">mdi-keyboard</v-icon>
        {{ t('input.typing') }}
      </v-card-title>
      <v-card-text class="pb-0 mb-0">
        <v-textarea
            ref="inputRef"
            v-model="inputText"
            :label="t('input.enterMessage')"
            variant="outlined"
            clearable
            :rows="appStore.isFullInputMode ? 10 : 5"
            no-resize
            density="comfortable"
            :placeholder="t('input.typeMessage')"
            @keydown="handleInputKeydown"
        />
      </v-card-text>
    </v-card>
    <v-textarea
        ref="inputFullRef"
        v-model="inputText"
        v-if="appStore.isFullInputMode"
        :label="t('input.enterMessage')"
        variant="outlined"
        clearable
        :rows="appStore.isFullInputMode ? 10 : 5"
        no-resize
        class="ma-0 pa-0"
        density="comfortable"
        :placeholder="t('input.typeMessage')"
        @keydown="handleFullPageKeyDown"
    />
    <!-- Actions Card -->
    <v-card v-if="!appStore.isFullInputMode">
      <v-card-title class="pb-2">
        <v-icon class="mr-2">mdi-gesture-tap-button</v-icon>
        {{ t('input.actions') }}
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
              {{ t('general.send') }}
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
              {{ t('general.clear') }}
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
              {{ t('input.fullInputMode') }}
            </v-btn>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>
</template>

<script setup lang="ts">
import {ref, computed} from 'vue'
import {useMessagesStore} from '@/stores/messages';
import {useAppStore} from '@/stores/app';
import {SendMessage, SetFullInputMode} from "../../wailsjs/go/pages/Input";
import {WindowGetSize, WindowSetSize} from "../../wailsjs/runtime";
import {useI18n} from "vue-i18n";

const { t } = useI18n()
const messagesStore = useMessagesStore()
const appStore = useAppStore()
const inputText = ref('')
const inputRef = ref<HTMLElement>()
const inputFullRef = ref<HTMLElement>()

const InputTextNotNull = computed(function () {
  return inputText?.value !== ""
})

const handleSend = () => {
  if (!inputText?.value.trim()) {
    messagesStore.addWarning(t('messages.pleaseEnterMessage'))
    return
  }

  SendMessage(inputText.value).then(() => {
  }).catch((error) => {
    messagesStore.addError(t('messages.failedToSend', { error: error.message }))
  }).finally(() => {
    messagesStore.addSuccess(t('messages.messageSentSuccess'))
  })
  inputText.value = ''
}

const handleClear = () => {
  if (inputText?.value.trim()) {
    inputText.value = ''
    messagesStore.addSuccess(t('messages.inputCleared'))
  } else {
    SendMessage('').then(() => {
      messagesStore.addInfo(t('messages.chatboxCleared'))
    }).catch((error) => {
      messagesStore.addError(t('messages.failedToClear', { error: error.message }))
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

function handleFullPageKeyDown(event: KeyboardEvent) {
  if (event.key === 'Escape' && appStore.isFullInputMode) {
    appStore.setInputMode(false)
    WindowGetSize().then((size) => {
      WindowSetSize(size.w, size.h + 180);
      (inputRef.value as HTMLElement)?.focus()
    })
  }
  handleInputKeydown(event)
}

// Function to enter full input mode
const enterFullInputMode = () => {
  appStore.setInputMode(true)
  WindowGetSize().then((size) => {
    WindowSetSize(size.w, size.h - 180)
    messagesStore.addInfo(t('input.fullInputModeEntered'));
    (inputFullRef.value as HTMLElement)?.focus()
  })
}

</script>