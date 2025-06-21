<template>
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
              ></v-switch>
            </v-col>
            <v-col cols="12">
              <v-text-field
                  label="Path"
                  v-model="opt.voicevox.path"
                  variant="outlined"
                  placeholder="VoiceVoxの実行ファイルパスを入力"
                  prepend-inner-icon="mdi-folder"
              ></v-text-field>
            </v-col>
            <!--
            <v-col cols="12">
              <v-number-input
                  label="LineLimit"
                  v-model="opt.voicevox.line_limit"
                  variant="outlined"
                  min.number="1"
                  control-variant="split"
              ></v-number-input>
            </v-col>
            -->
            <v-col cols="12">
              <v-combobox
                  label="Args"
                  v-model="opt.voicevox.args"
                  variant="outlined"
                  multiple
                  chips
              ></v-combobox>
            </v-col>
            <v-divider></v-divider>
            <v-col cols="12">
              <v-alert class="py-3" :type="vvAlert.type" :title="vvAlert.title"></v-alert>
            </v-col>
            <template v-if="running&&complete">
              <v-col cols="6">
                <v-select
                    label="Spacker"
                    v-model="selectedSpacker"
                    variant="outlined"
                    :items="spackers"
                    item-title="name"
                    item-value="uuid"
                    return-object
                ></v-select>
              </v-col>
              <v-col cols="6">
                <v-select
                    label="SpackerType"
                    v-model="selectedType"
                    variant="outlined"
                    :items="spackerTypes"
                    item-title="name"
                    item-value="id"
                    return-object
                ></v-select>
              </v-col>
            </template>
            <v-col cols="6">
              <v-btn width="100%" color="info" :disabled="running"  @click="handleVVStart" text="Start"></v-btn>
            </v-col>
            <v-col cols="6">
              <v-btn width="100%" color="error" :disabled="!running" @click="handleVVStop" text="Stop"></v-btn>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
    </v-form>
    <v-btn
        color="info"
        @click="handleSave"
        text="Save"
        width="100%"
        class="mt-4"
        :disabled="!loaded"
    ></v-btn>
</template>

<script setup lang="ts">
import {ref, onMounted} from 'vue';
import {api, options} from "../../wailsjs/go/models";
import {GetSpacker, IsVVComplete, IsVVRunning, Load, Save, StartVV, StopVV} from "../../wailsjs/go/pages/Options";
import {useMessagesStore} from "@/stores/messages";

const msg = useMessagesStore()
let loaded = ref(false)
let running = ref(false);
let complete = ref(false);
let runningTask: NodeJS.Timeout
let completeTask: NodeJS.Timeout
type alertType = 'success' | 'info' | 'warning' | 'error'

let vvAlert=ref({
  type: "error" as alertType,
  title: "VoiceVox is not running"
})

let opt = ref<options.Config>(new options.Config({
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
let spackers = ref([
  {
    name: "Default",
    id: -1,
  }
])
let selectedSpacker = ref({
  name: "Default",
  id: -1,
})

let spackerTypes = ref([
  {
    name: "Default",
    id: -1,
  }
])
let selectedType = ref({
  name: "Default",
  id: -1,
})

// Store the full speakers data for use in the watch function
let speakersData: api.Speaker[]

onMounted(() => {
  Load().then((options) => {
    opt.value = options
    loaded.value = true
  })
  runningUpdate().then(()=>{
    runningTask=setInterval(runningUpdate, 100000)
    completeUpdate().then(()=>{
      if (complete.value) {
        return
      }
      completeTask=setInterval(completeUpdate, 10000)
    })
  })
})

function runningUpdate() {
  return new Promise<void>((resolve, reject)=>{
    IsVVRunning().then((isRunning) => {
      if (isRunning) {
        vvAlert.value = {
          type: "success" as alertType,
          title: "VoiceVox is running"
        }
      }else{
        vvAlert.value = {
          type: "error" as alertType,
          title: "VoiceVox is not running"
        }
      }
      running.value = isRunning
      resolve()
    }).catch((error) => {
      reject(error)
    })
  })
}

function completeUpdate() {
  return new Promise<void>((resolve, reject)=>{
    IsVVComplete().then((c) => {
      if (complete.value===c) return
      if (c) {
        spackerUpdate()
        clearInterval(completeTask)
      }
      complete.value = c
      resolve()
    }).catch((error) => {
      reject(error)
    })
  })
}

function handleVVStart(){
  if (running.value){
    return
  }
  running.value = true
  StartVV().then(() => {
    runningUpdate().then(()=>{
      runningTask=setInterval(runningUpdate, 100000)
      completeTask=setInterval(completeUpdate, 10000)
    })
    msg.addSuccess("Started")
  })
}

function handleVVStop(){
  StopVV().then(() => {
    clearInterval(runningTask)
    clearInterval(completeTask)
    runningUpdate()
    running.value = false
    complete.value = false
    spackers.value = []
    spackerTypes.value = []
    msg.addSuccess("Stopped")
  }).catch((error) => {
    msg.addError(`Failed to stop: ${error.message}`)
  })
}


function spackerUpdate() {
  GetSpacker().then((sps) => {
    // Store the full speakers data
    speakersData = sps
    spackers.value = []
    spackerTypes.value = []
    let found = false
    for (let i = 0; i < sps.length; i++) {
      spackers.value.push({
        name: sps[i].name,
        id: i,
      })
      for (const s of sps[i].styles) {
        if (s.id === opt.value.voicevox.selected) {
          selectedSpacker.value = {
            name: sps[i].name,
            id: i,
          }
          selectedType.value = {
            name: s.name,
            id: s.id
          }
          found = true
        }
      }
    }
    if (!found) {
      if (spackers.value.length > 0) {
        selectedSpacker.value = spackers.value[0]
      }
      if (spackerTypes.value.length > 0) {
        selectedType.value = spackerTypes.value[0]
      }
    }
  })
}

watch(selectedSpacker, (newSpacker) => {
  const spTypes=speakersData[newSpacker.id].styles
  spackerTypes.value = []
  for (const s of spTypes) {
    spackerTypes.value.push({
      name: s.name,
      id: s.id
    })
  }
  if (spackerTypes.value.length > 0) {
    selectedType.value = spackerTypes.value[0]
  }
})

function handleSave() {
  // Update the selected speaker type ID in the config before saving
  opt.value.voicevox.selected = selectedType.value.id
  Save(opt.value).then(() => {
    msg.addSuccess("Saved")
  }).catch((error) => {
    msg.addError(`Failed to save: ${error.message}`)
  })
}


</script>
