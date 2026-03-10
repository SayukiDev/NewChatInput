import { describe, it, expect, vi, beforeEach } from 'vitest'
import { setActivePinia, createPinia } from 'pinia'
import { useLogStore } from '../../stores/log'

vi.mock('../../../wailsjs/go/pages/TTS', () => ({
  ReadLog: vi.fn(() => Promise.resolve('line1\nline2\nline3')),
}))

import { ReadLog } from '../../../wailsjs/go/pages/TTS'

describe('useLogStore', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
    vi.clearAllMocks()
  })

  it('has correct initial state', () => {
    const store = useLogStore()
    expect(store.content).toBe('')
    expect(store.isLoading).toBe(false)
  })

  it('fetches log and stores content', async () => {
    const store = useLogStore()
    await store.fetchLog()
    expect(ReadLog).toHaveBeenCalledOnce()
    expect(store.content).toBe('line1\nline2\nline3')
    expect(store.isLoading).toBe(false)
  })

  it('sets isLoading true while fetching', async () => {
    let resolve!: (v: string) => void
    ;(ReadLog as ReturnType<typeof vi.fn>).mockReturnValueOnce(
      new Promise<string>((r) => { resolve = r }),
    )
    const store = useLogStore()
    const promise = store.fetchLog()
    expect(store.isLoading).toBe(true)
    resolve('done')
    await promise
    expect(store.isLoading).toBe(false)
    expect(store.content).toBe('done')
  })
})
