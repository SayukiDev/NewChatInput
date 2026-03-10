<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import { useToast } from 'primevue/usetoast'
import { useOptionsStore } from '../stores/options'
import PortSettings from '../components/options/PortSettings.vue'

const { t } = useI18n()
const toast = useToast()
const optionsStore = useOptionsStore()

async function save() {
  try {
    await optionsStore.saveConfig().then(()=>{
      toast.add({ severity: 'success', summary: t('options.toast.savedSummary'), detail: t('options.toast.savedDetail'), life: 3000 })
    })
    toast.add({ severity: 'warn', summary: "注意", detail: "一部の変更は再起動する必要があります。", life: 3000 })
  } catch (e) {
    toast.add({ severity: 'error', summary: t('options.toast.errorSummary'), detail: t('options.toast.errorDetail'), life: 5000 })
  }
}
</script>

<template>
  <div class="options-view" v-if="optionsStore.config">

    <Card>
      <template #title>{{ $t('options.sections.oscPort') }}</template>
      <template #content>
        <PortSettings
          v-model:sendPort="optionsStore.config.send_port"
          v-model:recvPort="optionsStore.config.recv_port"
        />
      </template>
    </Card>

    <Card>
      <template #title>{{ $t('options.sections.settings') }}</template>
      <template #content>
        <div class="toggle-group">
          <div class="toggle-row">
            <label>{{ $t('options.toggles.typingIndicator') }}</label>
            <ToggleSwitch v-model="optionsStore.config.enable_typing_msg" />
          </div>
          <div class="toggle-row">
            <label>{{ $t('options.toggles.realtime') }}</label>
            <ToggleSwitch v-model="optionsStore.config.realtime" />
          </div>
          <div class="toggle-row">
            <label>{{ $t('options.toggles.messageKeeping') }}</label>
            <ToggleSwitch v-model="optionsStore.config.msg_keeping" />
          </div>
          <div class="toggle-row">
            <label>{{ $t('options.toggles.voiceControl') }}</label>
            <ToggleSwitch v-model="optionsStore.config.voice_control" />
          </div>
        </div>
      </template>
    </Card>

    <Card>
      <template #title>{{ $t('options.sections.tts') }}</template>
      <template #content>
        <div class="toggle-row">
          <label>{{ $t('options.toggles.enableTts') }}</label>
          <ToggleSwitch v-model="optionsStore.config.tts" />
        </div>
      </template>
    </Card>

    <div class="save-row">
      <Button :label="$t('options.saveButton')" icon="pi pi-save" @click="save" />
    </div>

  </div>
</template>

<style scoped>
.options-view {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.toggle-group {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
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
</style>
