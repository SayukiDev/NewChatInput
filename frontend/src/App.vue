<script lang="ts" setup>
import { useMessagesStore } from './stores/messages'

const messagesStore = useMessagesStore()
</script>

<template>
  <v-app>
    <v-main>
      <router-view />
    </v-main>
    
    <!-- Global Messages Bar -->
    <v-snackbar
      v-for="message in messagesStore.messages"
      :key="message.id"
      v-model="message.visible"
      :color="message.color"
      :timeout="message.timeout"
      location="bottom"
      @update:model-value="(value: boolean) => !value && messagesStore.removeMessage(message.id)"
    >
      {{ message.text }}
      <template #actions>
        <v-btn
          variant="text"
          @click="messagesStore.removeMessage(message.id)"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </v-app>
</template>

