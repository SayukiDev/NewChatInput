<script setup lang="ts">
import { computed, ref, watch, nextTick } from 'vue'
import { MAX_LENGTH } from '../../stores/chat'

const props = defineProps<{
  modelValue: string
  isSending: boolean
  isFullInput: boolean
}>()

const emit = defineEmits<{
  'update:modelValue': [value: string]
  send: []
  reset: []
  enterFullInput: []
  exitFullInput: []
}>()

const charCount = computed(() => props.modelValue.length)
const counterClass = computed(() => {
  if (charCount.value > MAX_LENGTH) return 'counter-error'
  if (charCount.value > 120) return 'counter-warn'
  return ''
})
const canSend = computed(
  () => props.modelValue.trim().length > 0 && !props.isSending && charCount.value <= MAX_LENGTH,
)
const canReset=computed(()=>!props.isSending)

function onKeydown(e: KeyboardEvent) {
  switch (e.key) {
    case 'Enter':
      if (e.shiftKey){
        return;
      }
      e.preventDefault()
      if (canSend.value) emit('send')
      break
    case 'Delete':
      e.preventDefault()
      if (canReset) emit('reset')
      break
    case 'Escape':
      if (props.isFullInput) {
        e.preventDefault()
        emit('exitFullInput')
      }
      break
  }
}

function onInput(val: string) {
  emit('update:modelValue', val)
}

const textareaRef = ref<{ $el: HTMLElement } | null>(null)

watch(
  () => props.isSending,
  (sending) => {
    if (!sending) {
      nextTick(() => {
        const el = textareaRef.value?.$el?.querySelector('textarea') ?? textareaRef.value?.$el
        el?.focus()
      })
    }
  },
)
</script>

<template>
  <div class="message-input">
    <Textarea
      ref="textareaRef"
      :modelValue="modelValue"
      @update:modelValue="onInput"
      :placeholder="$t('chat.input.placeholder')"
      :rows="isFullInput ? 1 : 5"
      class="input-area"
      :class="{ 'full-input-textarea': isFullInput }"
      :autofocus="isFullInput"
      @keydown="onKeydown"
    />
    <div v-if="!isFullInput" class="input-footer">
      <span :class="['char-counter', counterClass]">
        {{ charCount }}/{{ MAX_LENGTH }}
      </span>
      <div class="input-actions">
        <Button
          icon="pi pi-expand"
          severity="secondary"
          text
          rounded
          v-tooltip.left="$t('chat.fullInput.toastSummary')"
          @click="$emit('enterFullInput')"
        />
        <Button
          :label="$t('chat.input.send')"
          :disabled="!canSend"
          :loading="isSending"
          @click="$emit('send')"
        />
      </div>
    </div>
  </div>
</template>

<style scoped>
.message-input {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  flex: 1;
  height: 100%;
}

.input-area {
  width: 100%;
  resize: none;
  flex: 1;
  overflow-y: auto;
  min-height: 0;
}

.input-area :deep(textarea) {
  height: 100%;
}

.full-input-textarea :deep(textarea) {
  height: 100%;
}

.input-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.char-counter {
  font-size: 0.8rem;
  color: var(--p-text-muted-color);
}

.counter-warn {
  color: var(--p-orange-500);
}

.counter-error {
  color: var(--p-red-500);
  font-weight: 600;
}

.input-actions {
  display: flex;
  align-items: center;
  gap: 0.25rem;
}
</style>
