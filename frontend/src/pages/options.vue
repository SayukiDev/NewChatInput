<template>
  <div>
    <v-form @submit.prevent>
      <v-card class="mb-4">
        <v-card-title>
          <v-icon class="mr-2">mdi-console-network</v-icon>
          OSC
        </v-card-title>
        <v-card-text>
          <v-row>
            <v-col cols="6">
              <v-number-input
                  label="SendPort"
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
                  label="SendPort"
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
          Typing
        </v-card-title>
        <v-card-text>
          <v-row dense>
            <v-col cols="12">
              <v-switch 
                label="EnableTypingMsg" 
                density="compact"
                v-model="opt.enable_typing_msg"
              ></v-switch>
              <v-switch 
                label="RealtimeSend" 
                density="compact"
                v-model="opt.realtime"
              ></v-switch>
              <v-switch 
                label="EnableTTS" 
                density="compact"
                v-model="opt.tts"
              ></v-switch>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
      <v-card class="mt-4">
        <v-card-title>
          <v-icon class="mr-2">mdi-microphone</v-icon>
          VoiceVox
        </v-card-title>
        <v-card-text>
          <v-row>
            <v-col cols="12">
              <v-switch
                label="AutoStart"
                density="compact"
                v-model="opt.voicevox.auto_start"
                :disabled="!loaded"
              ></v-switch>
            </v-col>
            <v-col cols="12">
              <v-text-field
                label="実行パス"
                v-model="opt.voicevox.path"
                :disabled="!loaded"
                variant="outlined"
                placeholder="VoiceVoxの実行ファイルパスを入力"
                prepend-inner-icon="mdi-folder"
              ></v-text-field>
            </v-col>
            <v-col cols="6">
              <v-number-input
                label="LineLimit"
                :disabled="!loaded"
                v-model="opt.voicevox.line_limit"
                variant="outlined"
                min.number="1"
                control-variant="split"
              ></v-number-input>
            </v-col>
            <v-col cols="6">
              <v-number-input
                label="Spacker"
                :disabled="!loaded"
                v-model="opt.voicevox.selected"
                variant="outlined"
                min.number="0"
                control-variant="split"
              ></v-number-input>
            </v-col>
            <v-col cols="12">
              <v-textarea
                label="Args"
                :disabled="!loaded"
                v-model="opt.voicevox.args"
                variant="outlined"
                rows="1"
                persistent-hint
              ></v-textarea>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
    </v-form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import {options} from "../../wailsjs/go/models";
import {Load} from "../../wailsjs/go/pages/Options";
let loaded = ref(false)
let opt = ref<options.Options>(new options.Options({
  send_port: 9000,
  recv_port: 9001,
  enable_typing_msg: false,
  realtime: false,
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

onMounted(()=>{
  Load().then((options)=>{
    opt.value=options
    loaded.value=true
  })
})


</script>