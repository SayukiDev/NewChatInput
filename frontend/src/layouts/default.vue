<template>
  <div class="d-flex flex-column min-height-layout">
    <v-container fluid class="pa-4">
          <!-- Tabs Card with same width as content -->
          <v-card class="tabs-card mb-4">
            <v-tabs v-model="currentTab" grow class="flex-shrink-0">
              <v-tab value="input" @click="navigateToTab('input')">Input</v-tab>
              <v-tab value="options" @click="navigateToTab('options')">Options</v-tab>
            </v-tabs>
          </v-card>

          <!-- Content Area -->
          <div class="content-area">
            <router-view />
          </div>
    </v-container>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()
const currentTab = ref('input')

const navigateToTab = (tab: string) => {
  if (tab === 'input') {
    router.push('/')
  } else if (tab === 'options') {
    router.push('/options')
  }
}

// Sync route with tab
watch(
  () => route.path,
  (newPath) => {
    if (newPath === '/') {
      currentTab.value = 'input'
    } else if (newPath === '/options') {
      currentTab.value = 'options'
    }
  },
  { immediate: true }
)
</script>