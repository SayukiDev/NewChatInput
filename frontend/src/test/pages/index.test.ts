import { describe, it, expect, beforeEach } from 'vitest'
import { mount } from '@vue/test-utils'
import { createPinia, setActivePinia } from 'pinia'
import IndexPage from '@/pages/index.vue'

describe('Index Page - New Card Design', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
  })

  it('renders correctly with two cards', () => {
    const wrapper = mount(IndexPage)
    
    expect(wrapper.findAllComponents({ name: 'v-card' })).toHaveLength(2)
    expect(wrapper.findComponent({ name: 'v-textarea' }).exists()).toBe(true)
    expect(wrapper.findAllComponents({ name: 'v-btn' })).toHaveLength(2)
  })

  it('displays correct card titles', () => {
    const wrapper = mount(IndexPage)
    
    expect(wrapper.text()).toContain('Typing')
    expect(wrapper.text()).toContain('Actions')
  })

  it('has Typing card with correct structure', () => {
    const wrapper = mount(IndexPage)
    const cards = wrapper.findAllComponents({ name: 'v-card' })
    const typingCard = cards[0]
    
    expect(typingCard.text()).toContain('Typing')
    expect(typingCard.findComponent({ name: 'v-textarea' }).exists()).toBe(true)
  })

  it('has Actions card with Send and Clear buttons', () => {
    const wrapper = mount(IndexPage)
    const cards = wrapper.findAllComponents({ name: 'v-card' })
    const actionsCard = cards[1]
    
    expect(actionsCard.text()).toContain('Actions')
    expect(actionsCard.text()).toContain('Send')
    expect(actionsCard.text()).toContain('Clear')
    expect(actionsCard.findAllComponents({ name: 'v-btn' })).toHaveLength(2)
  })

  it('displays correct textarea properties', () => {
    const wrapper = mount(IndexPage)
    const textarea = wrapper.findComponent({ name: 'v-textarea' })
    
    expect(textarea.props('label')).toBe('Enter your message')
    expect(textarea.props('placeholder')).toBe('Type your message here...')
    expect(textarea.props('variant')).toBe('outlined')
    expect(textarea.props('noResize')).toBe(true)
    expect(textarea.props('rows')).toBe(5)
  })

  it('has icons in card titles and buttons', () => {
    const wrapper = mount(IndexPage)
    const icons = wrapper.findAllComponents({ name: 'v-icon' })
    
    // Should have icons for: keyboard (Typing), gesture-tap-button (Actions), send, eraser
    expect(icons.length).toBeGreaterThanOrEqual(4)
  })

  it('buttons are disabled when input is empty', () => {
    const wrapper = mount(IndexPage)
    const buttons = wrapper.findAllComponents({ name: 'v-btn' })
    
    buttons.forEach(button => {
      expect(button.props('disabled')).toBe(true)
    })
  })

  it('buttons are enabled when input has content', async () => {
    const wrapper = mount(IndexPage)
    const textarea = wrapper.findComponent({ name: 'v-textarea' })
    
    await textarea.vm.$emit('update:modelValue', 'test message')
    await wrapper.vm.$nextTick()
    
    const buttons = wrapper.findAllComponents({ name: 'v-btn' })
    
    buttons.forEach(button => {
      expect(button.props('disabled')).toBe(false)
    })
  })

  it('has proper button functionality structure', () => {
    const wrapper = mount(IndexPage)
    const buttons = wrapper.findAllComponents({ name: 'v-btn' })
    
    const sendButton = buttons.find(btn => btn.text().includes('Send'))
    const clearButton = buttons.find(btn => btn.text().includes('Clear'))
    
    expect(sendButton?.exists()).toBe(true)
    expect(clearButton?.exists()).toBe(true)
    expect(sendButton?.props('color')).toBe('primary')
    expect(clearButton?.props('color')).toBe('secondary')
  })

  it('textarea supports multiline input', async () => {
    const wrapper = mount(IndexPage)
    const textarea = wrapper.findComponent({ name: 'v-textarea' })
    
    const multilineText = 'Line 1\nLine 2\nLine 3'
    
    await textarea.vm.$emit('update:modelValue', multilineText)
    await wrapper.vm.$nextTick()
    
    expect(textarea.props('modelValue')).toBe(multilineText)
  })

  it('has cards with consistent structure', () => {
    const wrapper = mount(IndexPage)
    
    // Check for column classes in button layout
    const cols = wrapper.findAll('[class*="col-"]')
    expect(cols.length).toBeGreaterThan(0)
    
    // Check for proper card structure - cards should be direct children
    expect(wrapper.findAllComponents({ name: 'v-card' })).toHaveLength(2)
    expect(wrapper.findComponent({ name: 'v-row' }).exists()).toBe(true)
    
    // Should not have v-container since width is controlled by layout
    expect(wrapper.findComponent({ name: 'v-container' }).exists()).toBe(false)
  })
})