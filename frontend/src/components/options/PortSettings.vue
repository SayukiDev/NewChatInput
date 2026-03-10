<script setup lang="ts">
const props = defineProps<{
  sendPort: number
  recvPort: number
}>()

const emit = defineEmits<{
  'update:sendPort': [value: number]
  'update:recvPort': [value: number]
}>()
</script>

<template>
  <div class="port-settings">
    <div class="field">
      <label>{{ $t('options.port.send') }}</label>
      <InputNumber
        :modelValue="sendPort"
        @update:modelValue="$emit('update:sendPort', $event)"
        :min="1"
        :max="65535"
        :useGrouping="false"
      />
    </div>
    <div class="field">
      <label>{{ $t('options.port.receive') }}</label>
      <InputNumber
        :modelValue="recvPort"
        @update:modelValue="$emit('update:recvPort', $event)"
        :min="1"
        :max="65535"
        :useGrouping="false"
      />
    </div>
  </div>
</template>

<style scoped>
.port-settings {
  display: flex;
  gap: 1rem;
}

.field {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
  flex: 1;
}

.field label {
  font-size: 0.85rem;
  font-weight: 500;
}
</style>
