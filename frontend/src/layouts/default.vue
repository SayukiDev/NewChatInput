<template>
    <v-container fluid class="pa-4">
      <!-- Tabs Card with same width as content -->
      <v-card v-if="!appStore.isFullInputMode" class="tabs-card">
        <v-row align="center" no-gutters>
          <v-col>
            <v-tabs v-model="currentTab" grow>
              <v-tab value="input" @click="navigateToTab('input')">{{ t('tabs.input') }}</v-tab>
              <v-tab value="options" @click="navigateToTab('options')">{{ t('tabs.options') }}</v-tab>
            </v-tabs>
          </v-col>
          <v-col cols="auto" class="px-2">
          </v-col>
        </v-row>
      </v-card>

      <!-- Content Area -->
      <v-container class="pb-0 px-0" width="100%">
        <router-view/>
      </v-container>
    </v-container>
</template>

<script setup lang="ts">
import {ref, watch} from 'vue'
import {useRoute, useRouter} from 'vue-router'
import {useAppStore} from '@/stores/app'

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const appStore = useAppStore()
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
    {immediate: true}
)
</script>