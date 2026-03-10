import { describe, it, expect, vi, beforeEach } from 'vitest'
import { setActivePinia, createPinia } from 'pinia'
import { useChatStore, MAX_LENGTH } from '../../stores/chat'

vi.mock('../../../wailsjs/go/pages/Input', () => ({
  SendMessage: vi.fn(() => Promise.resolve()),
  SetFullInputMode: vi.fn(() => Promise.resolve()),
}))

vi.mock('../../../wailsjs/go/pages/Options', () => ({
  SetTyping: vi.fn(() => Promise.resolve()),
}))

import { SendMessage } from '../../../wailsjs/go/pages/Input'

describe('useChatStore', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
    vi.clearAllMocks()
  })

  it('has correct initial state', () => {
    const store = useChatStore()
    expect(store.message).toBe('')
    expect(store.isSending).toBe(false)
  })

  it('sends a message and clears input', async () => {
    const store = useChatStore()
    store.message = 'Hello VRChat!'
    await store.sendMessage()
    expect(SendMessage).toHaveBeenCalledWith('Hello VRChat!')
    expect(store.message).toBe('')
    expect(store.isSending).toBe(false)
  })

  it('does not send empty messages', async () => {
    const store = useChatStore()
    store.message = '   '
    await store.sendMessage()
    expect(SendMessage).not.toHaveBeenCalled()
  })

  it('does not send while already sending', async () => {
    const store = useChatStore()
    store.message = 'test'
    store.isSending = true
    await store.sendMessage()
    expect(SendMessage).not.toHaveBeenCalled()
  })

  it('exports MAX_LENGTH as 144', () => {
    expect(MAX_LENGTH).toBe(144)
  })
})
