<script setup lang="ts">
import { computed, onMounted, watch, useTemplateRef } from 'vue'
import { useI18n } from 'vue-i18n'
import { useToast } from 'primevue/usetoast'
import { useOptionsStore } from '../stores/options'
import { useTTSStore } from '../stores/tts'
import { useLogStore } from '../stores/log'
import SpeakerSelector from '../components/options/SpeakerSelector.vue'
import AudioDeviceSelector from '../components/options/AudioDeviceSelector.vue'

const { t } = useI18n()
const toast = useToast()
const optionsStore = useOptionsStore()
const ttsStore = useTTSStore()
const logStore = useLogStore()
const logBoxRef = useTemplateRef<HTMLPreElement>('logBox')

function scrollToBottom() {
  const el = logBoxRef.value
  if (el) el.scrollTop = el.scrollHeight
}

watch(() => logStore.content, scrollToBottom, { flush: 'post' })

onMounted(() => {
  logStore.fetchLog()
  ttsStore.checkIsRunning()
})

const argsText = computed({
  get: () => optionsStore.config?.tts_option.args?.join(' ') ?? '',
  set: (val: string) => {
    if (optionsStore.config) {
      optionsStore.config.tts_option.args = val.trim() ? val.trim().split(/\s+/) : []
    }
  },
})

async function loadTTS() {
  if (!optionsStore.config) return
  if (!optionsStore.config.tts) {
    ttsStore.currentSpeakerId = -1
    ttsStore.selectedDeviceId = ''
    ttsStore.speakers = []
    ttsStore.devices = []
    ttsStore.styleNames = []
    return
  }
  try {
    await ttsStore.loadSpeakers()
    await Promise.all([
      ttsStore.loadCurrentSpeaker(),
      ttsStore.loadDevices(optionsStore.config.tts_option.device),
    ])
  } catch (e) {
    console.error('[loadTTS] error:', e)
  }
}

watch(() => optionsStore.config, loadTTS, { immediate: true })

async function save() {
  try {
    await optionsStore.saveConfig().then(() => {
      toast.add({ severity: 'success', summary: t('options.toast.savedSummary'), detail: t('options.toast.savedDetail'), life: 3000 })
    })
    await loadTTS()
    toast.add({ severity: 'warn', summary: "注意", detail: "一部の変更は再起動する必要があります。", life: 3000 })
  } catch (e) {
    toast.add({ severity: 'error', summary: t('options.toast.errorSummary'), detail: t('options.toast.errorDetail'), life: 5000 })
  }
}
</script>

<template>
  <div class="tts-view" v-if="optionsStore.config">

    <Card>
      <template #title>{{ $t('tts.sections.voiceSettings') }}</template>
      <template #content>
        <template v-if="optionsStore.config.tts">
          <SpeakerSelector />
          <AudioDeviceSelector style="margin-top: 0.75rem;" />
        </template>
        <p v-else class="disabled-notice">{{ $t('tts.disabledNotice') }}</p>
      </template>
    </Card>

    <Card>
      <template #title>{{ $t('tts.sections.engineSettings') }}</template>
      <template #content>
        <Message style="margin-bottom: 0.75rem;" v-if="ttsStore.isRunning" severity="success" :closable="false"><i class="pi pi-check"></i> {{ $t('tts.engine.statusRunning') }}</Message>
        <Message style="margin-bottom: 0.75rem;" v-else severity="warn" :closable="false"><i class="pi pi-times"></i> {{ $t('tts.engine.statusStopped') }}</Message>
        <div class="field-group" >
          <div class="field">
            <label>{{ $t('tts.engine.baseurl') }}</label>
            <InputText v-model="optionsStore.config.tts_option.Baseurl" fluid />
          </div>
          <div class="field">
            <label>{{ $t('tts.engine.path') }}</label>
            <InputText v-model="optionsStore.config.tts_option.path" fluid />
          </div>
          <div class="field">
            <label>{{ $t('tts.engine.args') }}</label>
            <InputText v-model="argsText" fluid />
          </div>
          <div class="field">
            <label>{{ $t('tts.engine.log') }}</label>
            <InputText v-model="optionsStore.config.tts_option.log" fluid />
          </div>
          <div class="toggle-row">
            <label>{{ $t('tts.engine.run') }}</label>
            <ToggleSwitch v-model="optionsStore.config.tts_option.run" />
          </div>
        </div>
      </template>
    </Card>

    <Card>
      <template #title>{{ $t('tts.sections.cacheSettings') }}</template>
      <template #content>
        <div class="field-group">
          <div class="toggle-row">
            <label>{{ $t('tts.cache.enable') }}</label>
            <ToggleSwitch v-model="optionsStore.config.tts_option.cache" />
          </div>
          <div class="field" v-if="optionsStore.config.tts_option.cache">
            <label>{{ $t('tts.cache.path') }}</label>
            <InputText v-model="optionsStore.config.tts_option.cache_path" fluid />
          </div>
        </div>
      </template>
    </Card>

    <Card>
      <template #title>
        <div class="log-header">
          <span>{{ $t('tts.sections.log') }}</span>
          <Button
            icon="pi pi-refresh"
            text
            rounded
            :loading="logStore.isLoading"
            :disabled="logStore.isLoading"
            v-tooltip.left="$t('log.refresh')"
            @click="logStore.fetchLog()"
          />
        </div>
      </template>
      <template #content>
        <div class="log-content-wrapper">
          <ProgressSpinner v-if="logStore.isLoading && !logStore.content" class="log-spinner" />
          <pre ref="logBox" class="log-box">{{ logStore.content }}</pre>
        </div>
      </template>
    </Card>

    <div class="save-row">
      <Button :label="$t('options.saveButton')" icon="pi pi-save" @click="save" />
    </div>

  </div>
</template>

<style scoped>
.tts-view {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.field-group {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.field {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.field label {
  font-size: 0.9rem;
}

.toggle-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.toggle-row label {
  font-size: 0.9rem;
}

.save-row {
  display: flex;
  justify-content: flex-end;
  padding-top: 0.25rem;
}

.disabled-notice {
  color: var(--p-text-muted-color);
  font-size: 0.9rem;
}

.log-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
}

.log-content-wrapper {
  position: relative;
}

.log-spinner {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 2rem;
  height: 2rem;
}

.log-box {
  font-family: 'Consolas', 'Menlo', 'Courier New', monospace;
  font-size: 0.8rem;
  line-height: 1.5;
  white-space: pre-wrap;
  word-break: break-all;
  overflow-y: auto;
  max-height: 55vh;
  margin: 0;
  padding: 0.5rem;
  background: #ffffff;
  color: #ff69b4;
  border-radius: var(--p-border-radius-sm, 4px);
  border: 1px solid var(--p-inputtext-border-color, #a0a0a0);
  outline: none;
  transition: border-color 0.2s, box-shadow 0.2s;
}
</style>
