<template>
    <v-form @submit.prevent>
      <v-card class="mb-4">
        <v-card-title>
          <v-icon class="mr-2">mdi-console-network</v-icon>
          {{ t('options.osc') }}
        </v-card-title>
        <v-card-text>
          <v-row>
            <v-col cols="6">
              <v-number-input
                  :label="t('options.sendPort')"
                  v-model="opt.send_port"
                  variant="outlined"
                  max.number="65535"
                  min.number="1"
                  control-variant="split"
              >
              </v-number-input>
            </v-col>
            <v-col cols="6">
              <v-number-input
                  :label="t('options.recvPort')"
                  v-model="opt.recv_port"
                  variant="outlined"
                  max.number="65535"
                  min.number="1"
                  control-variant="split"
              >
              </v-number-input>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
      <v-card>
        <v-card-title>
          <v-icon class="mr-2">mdi-keyboard</v-icon>
          {{ t('options.typing') }}
        </v-card-title>
        <v-card-text>
          <v-row dense>
            <v-col cols="12">
              <v-switch
                  :label="t('options.enableTypingMsg')"
                  density="compact"
                  v-model="opt.enable_typing_msg"
              ></v-switch>
              <v-switch
                  :label="t('options.realtimeSend')"
                  density="compact"
                  v-model="opt.realtime"
              ></v-switch>
              <v-switch
                  :label="t('options.enableTTS')"
                  density="compact"
                  v-model="opt.tts"
              ></v-switch>
              <v-switch
                  :label="t('options.msgKeeping')"
                  density="compact"
                  v-model="opt.msg_keeping"
              ></v-switch>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
      <VoiceVoxOptions ref="vvRef" v-model:voicevox="opt.voicevox" :loaded="loaded" />
    </v-form>
    <v-btn
        color="info"
        @click="handleSave"
        :text="t('general.save')"
        width="100%"
        class="mt-4"
        :disabled="!loaded"
    ></v-btn>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { options } from "../../wailsjs/go/models";
import { Load, Save } from "../../wailsjs/go/pages/Options";
import { useMessagesStore } from "@/stores/messages";
import VoiceVoxOptions from "@/components/VoiceVoxOptions.vue";

const { t } = useI18n()
const msg = useMessagesStore()
let loaded = ref(false)
let vvRef = ref<any>(null)

let opt = ref<options.Config>(new options.Config({
  send_port: 9000,
  recv_port: 9001,
  enable_typing_msg: false,
  realtime: false,
  msg_keeping: false,
  tts: false,
  voice_control: false,
  voicevox: {
    auto_start: false,
    path: "",
    line_limit: 50,
    selected: 0,
    args: []
  }
}));

onMounted(() => {
  Load().then((options) => {
    opt.value = options
    loaded.value = true
  })
})

function handleSave() {
  Save(opt.value)
    .then(() => {
      msg.addSuccess(t('messages.saved'))
      // notify child to update its internal state post-save
      vvRef.value?.onSaved?.()
    })
    .catch((error) => {
      msg.addError(t('messages.failedToSave', { error: error.message }))
    })
}
</script>
