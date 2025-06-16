import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import OptionsPage from '@/pages/options.vue'

describe('Options Page', () => {
  it('renders correctly with simplified structure', () => {
    const wrapper = mount(OptionsPage)
    
    // Should have card but no container structure
    expect(wrapper.findComponent({ name: 'v-card' }).exists()).toBe(true)
    expect(wrapper.findComponent({ name: 'v-container' }).exists()).toBe(false)
    expect(wrapper.findComponent({ name: 'v-row' }).exists()).toBe(false)
    expect(wrapper.findComponent({ name: 'v-col' }).exists()).toBe(false)
  })

  it('displays correct title and icon', () => {
    const wrapper = mount(OptionsPage)
    
    expect(wrapper.text()).toContain('Options')
    expect(wrapper.findComponent({ name: 'v-icon' }).exists()).toBe(true)
  })

  it('shows placeholder message for future implementation', () => {
    const wrapper = mount(OptionsPage)
    
    expect(wrapper.text()).toContain('Options configuration will be available in future updates.')
  })

  it('uses default layout', () => {
    const wrapper = mount(OptionsPage)
    
    // Check if the component is using the default layout
    expect(wrapper.vm.$route?.meta?.layout).toBe('default')
  })

  it('has proper card structure matching other pages', () => {
    const wrapper = mount(OptionsPage)
    
    // Should have exactly one card with title and content
    expect(wrapper.findAllComponents({ name: 'v-card' })).toHaveLength(1)
    expect(wrapper.findComponent({ name: 'v-card-title' }).exists()).toBe(true)
    expect(wrapper.findComponent({ name: 'v-card-text' }).exists()).toBe(true)
  })
})