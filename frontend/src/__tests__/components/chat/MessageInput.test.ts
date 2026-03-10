import { describe, it, expect, vi, beforeEach } from 'vitest'
import { mount } from '@vue/test-utils'
import { createPinia, setActivePinia } from 'pinia'
import MessageInput from '../../../components/chat/MessageInput.vue'

// Mock PrimeVue components used by MessageInput
vi.mock('primevue/textarea', () => ({
  default: {
    name: 'Textarea',
    template: '<textarea :value="modelValue" @input="$emit(\'update:modelValue\', $event.target.value)" @keydown="$emit(\'keydown\', $event)" :disabled="disabled"></textarea>',
    props: ['modelValue', 'disabled', 'autoResize', 'rows', 'placeholder'],
    emits: ['update:modelValue', 'keydown'],
  },
}))

vi.mock('primevue/button', () => ({
  default: {
    name: 'Button',
    template: '<button :disabled="disabled" @click="$emit(\'click\')"><slot /></button>',
    props: ['icon', 'disabled', 'loading', 'severity', 'size', 'text', 'rounded', 'label'],
    emits: ['click'],
  },
}))

describe('MessageInput', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
  })

  function createWrapper(props = {}) {
    return mount(MessageInput, {
      props: {
        modelValue: '',
        isSending: false,
        isFullInput: false,
        ...props,
      },
      global: {
        stubs: {
          Textarea: {
            template: '<textarea :value="modelValue" @input="$emit(\'update:modelValue\', $event.target.value)" @keydown="$emit(\'keydown\', $event)" :disabled="disabled"></textarea>',
            props: ['modelValue', 'disabled', 'autoResize', 'rows', 'placeholder'],
            emits: ['update:modelValue', 'keydown'],
          },
          Button: {
            template: '<button :disabled="disabled" @click="$emit(\'click\')"><slot /></button>',
            props: ['icon', 'disabled', 'loading', 'severity', 'size', 'text', 'rounded', 'label'],
            emits: ['click'],
          },
        },
        directives: {
          tooltip: {},
        },
      },
    })
  }

  it('displays character counter', () => {
    const wrapper = createWrapper({ modelValue: 'Hello' })
    expect(wrapper.text()).toContain('5/144')
  })

  it('displays 0/144 when empty', () => {
    const wrapper = createWrapper()
    expect(wrapper.text()).toContain('0/144')
  })

  it('applies warning class when over 120 chars', () => {
    const wrapper = createWrapper({ modelValue: 'a'.repeat(121) })
    const counter = wrapper.find('.char-counter')
    expect(counter.classes()).toContain('counter-warn')
  })

  it('applies error class when over 144 chars', () => {
    const wrapper = createWrapper({ modelValue: 'a'.repeat(145) })
    const counter = wrapper.find('.char-counter')
    expect(counter.classes()).toContain('counter-error')
  })

  it('emits send on Enter key', async () => {
    const wrapper = createWrapper({ modelValue: 'Hello' })
    const textarea = wrapper.find('textarea')
    await textarea.trigger('keydown', { key: 'Enter', shiftKey: false })
    expect(wrapper.emitted('send')).toBeTruthy()
  })

  it('does not emit send on Shift+Enter', async () => {
    const wrapper = createWrapper({ modelValue: 'Hello' })
    const textarea = wrapper.find('textarea')
    await textarea.trigger('keydown', { key: 'Enter', shiftKey: true })
    expect(wrapper.emitted('send')).toBeFalsy()
  })

  it('does not emit send when message is empty', async () => {
    const wrapper = createWrapper({ modelValue: '' })
    const textarea = wrapper.find('textarea')
    await textarea.trigger('keydown', { key: 'Enter', shiftKey: false })
    expect(wrapper.emitted('send')).toBeFalsy()
  })

  it('emits update:modelValue on input', async () => {
    const wrapper = createWrapper()
    const textarea = wrapper.find('textarea')
    await textarea.setValue('test')
    expect(wrapper.emitted('update:modelValue')).toBeTruthy()
  })
})
