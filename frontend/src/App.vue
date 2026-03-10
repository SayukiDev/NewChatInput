<script setup lang="ts">
import { onMounted, onUnmounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAppStore } from './stores/app'
import { useOptionsStore } from './stores/options'
import {useChatStore} from "./stores/chat.ts";
import { GetSizeRatio } from '../wailsjs/go/pages/App'
import {WindowGetSize, WindowSetSize} from '../wailsjs/runtime/runtime'

const router = useRouter()
const route = useRoute()
const appStore = useAppStore()
const optionsStore = useOptionsStore()
const chatStore = useChatStore()
let resizeTimer: ReturnType<typeof setTimeout> | null = null

function onResize() {
  if (resizeTimer) clearTimeout(resizeTimer)
  resizeTimer = setTimeout(() => {
    window.removeEventListener('resize', onResize)
    GetSizeRatio().then(ratio => {
      WindowGetSize().then(({w, }) => {
        const width = w
        const height = Math.round(width / ratio)
        WindowSetSize(width, height)
        setTimeout(() => window.addEventListener('resize', onResize), 100)
      })
    })
  }, 100)
}

onMounted(async () => {
  window.addEventListener('resize', onResize)
  appStore.startPolling()
})

onUnmounted(() => {
  window.removeEventListener('resize', onResize)
  if (resizeTimer) clearTimeout(resizeTimer)
})

watch(
  () => appStore.isReady,
  async (ready) => {
    if (ready) {
      await optionsStore.loadConfig()
    }
  },
)

function onTabChange(path: string) {
  router.push(path)
}
</script>

<template>
  <div v-if="!appStore.isReady" class="loading">
    <ProgressSpinner />
    <p>{{ $t('app.connecting') }}</p>
  </div>
  <template v-else>
    <Card class="app-card">
      <template #header v-if="!chatStore.isFullInput">
        <Tabs :value="route.path">
          <TabList>
            <Tab value="/chat" @click="onTabChange('/chat')">{{ $t('app.tabs.chat') }}</Tab>
            <Tab value="/options" @click="onTabChange('/options')">{{ $t('app.tabs.options') }}</Tab>
            <Tab value="/tts" @click="onTabChange('/tts')">{{ $t('app.tabs.tts') }}</Tab>
          </TabList>
        </Tabs>
      </template>
      <template #content>
        <router-view />
      </template>
    </Card>
    <!--<StatusBar />-->
    <Toast />
  </template>
</template>

<style scoped>
.loading {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  gap: 1rem;
}

.loading p {
  color: var(--p-text-muted-color);
}

.app-card {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

:deep(.p-card-body) {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  padding: 1rem;
}

:deep(.p-card-content) {
  flex: 1;
  overflow-y: auto;
  padding: 0.5rem;
}

:deep(.p-tablist-tab-list) {
  justify-content: center;
}
</style>
