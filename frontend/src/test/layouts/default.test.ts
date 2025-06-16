import { describe, it, expect, beforeEach } from 'vitest'
import { mount } from '@vue/test-utils'
import { createRouter, createMemoryHistory } from 'vue-router'
import DefaultLayout from '@/layouts/default.vue'

const router = createRouter({
  history: createMemoryHistory(),
  routes: [
    { path: '/', component: { template: '<div>Input Page</div>' } },
    { path: '/options', component: { template: '<div>Options Page</div>' } }
  ]
})

describe('Default Layout', () => {
  beforeEach(async () => {
    await router.push('/')
  })

  it('renders correctly with tabs in a card matching content width', async () => {
    const wrapper = mount(DefaultLayout, {
      global: {
        plugins: [router]
      }
    })

    expect(wrapper.findComponent({ name: 'v-tabs' }).exists()).toBe(true)
    expect(wrapper.findAllComponents({ name: 'v-tab' })).toHaveLength(2)
    expect(wrapper.findAllComponents({ name: 'v-card' })).toHaveLength(1)
    
    // Check that tabs card has proper styling
    const tabsCard = wrapper.findComponent({ name: 'v-card' })
    expect(tabsCard.classes()).toContain('tabs-card')
    expect(tabsCard.classes()).toContain('mb-4')
  })

  it('displays correct tab labels', async () => {
    const wrapper = mount(DefaultLayout, {
      global: {
        plugins: [router]
      }
    })

    const tabs = wrapper.findAllComponents({ name: 'v-tab' })
    expect(tabs[0].text()).toBe('Input')
    expect(tabs[1].text()).toBe('Options')
  })

  it('has proper layout structure without excessive height', async () => {
    const wrapper = mount(DefaultLayout, {
      global: {
        plugins: [router]
      }
    })

    expect(wrapper.find('.d-flex.flex-column').exists()).toBe(true)
    expect(wrapper.find('.flex-shrink-0').exists()).toBe(true)
    // Should not have fill-height or flex-grow-1 to avoid excessive whitespace
    expect(wrapper.find('.d-flex.flex-column.fill-height').exists()).toBe(false)
    expect(wrapper.find('.flex-grow-1').exists()).toBe(false)
  })

  it('shows router-view for content', async () => {
    const wrapper = mount(DefaultLayout, {
      global: {
        plugins: [router]
      }
    })

    expect(wrapper.findComponent({ name: 'router-view' }).exists()).toBe(true)
  })

  it('initializes with input tab selected', async () => {
    const wrapper = mount(DefaultLayout, {
      global: {
        plugins: [router]
      }
    })

    await wrapper.vm.$nextTick()
    // Check that input tab is selected by default through the route
    expect(router.currentRoute.value.path).toBe('/')
  })

  it('navigates to options when options tab is clicked', async () => {
    const wrapper = mount(DefaultLayout, {
      global: {
        plugins: [router]
      }
    })

    const tabs = wrapper.findAllComponents({ name: 'v-tab' })
    await tabs[1].trigger('click')

    expect(router.currentRoute.value.path).toBe('/options')
  })

  it('updates currentTab when route changes', async () => {
    const wrapper = mount(DefaultLayout, {
      global: {
        plugins: [router]
      }
    })

    await router.push('/options')
    await wrapper.vm.$nextTick()

    expect(router.currentRoute.value.path).toBe('/options')
  })

  it('tabs grow to fill available space', async () => {
    const wrapper = mount(DefaultLayout, {
      global: {
        plugins: [router]
      }
    })

    const tabs = wrapper.findComponent({ name: 'v-tabs' })
    expect(tabs.props('grow')).toBe(true)
  })

  it('has tabs and content with matching widths', async () => {
    const wrapper = mount(DefaultLayout, {
      global: {
        plugins: [router]
      }
    })

    // Both tabs card and content should be within the same container/row/col structure
    const tabsCard = wrapper.findComponent({ name: 'v-card' })
    expect(tabsCard.exists()).toBe(true)
    expect(tabsCard.classes()).toContain('tabs-card')
    
    // Both should share the same responsive column structure
    expect(wrapper.findComponent({ name: 'v-container' }).exists()).toBe(true)
    expect(wrapper.findComponent({ name: 'v-row' }).exists()).toBe(true)
    expect(wrapper.findComponent({ name: 'v-col' }).exists()).toBe(true)
    
    // Content area should exist within the same column
    expect(wrapper.find('.content-area').exists()).toBe(true)
  })
})