<script setup lang="ts">
import { useTTSStore } from '../../stores/tts'

const ttsStore = useTTSStore()
</script>

<template>
  <div class="audio-device-selector">
    <div class="field">
      <label>{{ $t('options.audioDevice.label') }}</label>
      <Select
        :modelValue="ttsStore.selectedDeviceId"
        :options="ttsStore.devices"
        optionLabel="Name"
        optionValue="Id"
        :placeholder="$t('options.audioDevice.placeholder')"
        @update:modelValue="ttsStore.selectDevice($event)"
      />
    </div>
  </div>
</template>

<style scoped>
.field {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.field label {
  font-size: 0.85rem;
  font-weight: 500;
}
</style>
