<script setup lang="ts">
import { computed } from 'vue'
import { useTTSStore } from '../../stores/tts'

const ttsStore = useTTSStore()

const currentSpeakerObject = computed(() =>
  ttsStore.speakers.find(s => s.styles.some(st => st.id === ttsStore.currentSpeakerId))
)

const styleOptions = computed(() =>
  ttsStore.styleNames.map((name, i) => ({
    label: name,
    value: currentSpeakerObject.value?.styles[i]?.id ?? i,
  })),
)
</script>

<template>
  <div class="speaker-selector">
    <div class="field">
      <label>{{ $t('options.speaker.label') }}</label>
      <Select
        :modelValue="currentSpeakerObject"
        :options="ttsStore.speakers"
        optionLabel="name"
        :placeholder="$t('options.speaker.placeholder')"
        @update:modelValue="ttsStore.selectSpeaker($event)"
      />
    </div>
    <div class="field" v-if="ttsStore.styleNames.length > 0">
      <label>{{ $t('options.speaker.styleLabel') }}</label>
      <Select
        :modelValue="ttsStore.currentSpeakerId"
        :options="styleOptions"
        optionLabel="label"
        optionValue="value"
        :placeholder="$t('options.speaker.stylePlaceholder')"
        @update:modelValue="ttsStore.selectStyle($event)"
      />
    </div>
  </div>
</template>

<style scoped>
.speaker-selector {
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
  font-size: 0.85rem;
  font-weight: 500;
}
</style>
