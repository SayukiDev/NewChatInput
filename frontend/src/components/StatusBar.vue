<script setup lang="ts">
import { ref } from 'vue'
import { useOptionsStore } from '../stores/options'
import { WindowSetAlwaysOnTop } from '../../wailsjs/runtime/runtime'

const optionsStore = useOptionsStore()
const alwaysOnTop = ref(false)

async function toggleAlwaysOnTop() {
  alwaysOnTop.value = !alwaysOnTop.value
  await WindowSetAlwaysOnTop(alwaysOnTop.value)
}
</script>

<template>
  <div class="status-bar">
    <div class="status-left">
      <span v-if="optionsStore.config?.realtime" class="realtime-badge">
        {{ $t('options.toggles.realtime') }}
      </span>
    </div>
    <div class="status-right">
      <Button
        :icon="alwaysOnTop ? 'pi pi-lock' : 'pi pi-lock-open'"
        :severity="alwaysOnTop ? 'info' : 'secondary'"
        size="small"
        text
        rounded
        @click="toggleAlwaysOnTop"
      />
    </div>
  </div>
</template>

<style scoped>
.status-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0.25rem 0.75rem;
  border-top: 1px solid #f9a8d4;
  flex-shrink: 0;
  background-color: #fff0f5;
}

.realtime-badge {
  font-size: 0.75rem;
  color: var(--p-primary-color);
  font-weight: 600;
}
</style>
