import { describe, it, expect, vi, beforeEach } from 'vitest'
import { setActivePinia, createPinia } from 'pinia'
import { useOptionsStore } from '../../stores/options'

const mockConfig = {
  send_port: 9000,
  recv_port: 9001,
  enable_typing_msg: true,
  realtime: false,
  msg_keeping: false,
  tts: false,
  voice_control: false,
  tts_engine: {
    run: false,
    log: '',
    path: '',
    args: [],
    now_spacker: '',
    now_style: 0,
    device: 0,
  },
}

vi.mock('../../../wailsjs/go/pages/Options', () => ({
  Load: vi.fn(() => Promise.resolve({ ...mockConfig })),
  Save: vi.fn(() => Promise.resolve()),
}))

import { Load, Save } from '../../../wailsjs/go/pages/Options'

describe('useOptionsStore', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
    vi.clearAllMocks()
  })

  it('has null config initially', () => {
    const store = useOptionsStore()
    expect(store.config).toBeNull()
  })

  it('loads config from backend', async () => {
    const store = useOptionsStore()
    await store.loadConfig()
    expect(Load).toHaveBeenCalled()
    expect(store.config).not.toBeNull()
    expect(store.config!.send_port).toBe(9000)
  })

  it('saves config to backend', async () => {
    const store = useOptionsStore()
    await store.loadConfig()
    store.config!.send_port = 8888
    await store.saveConfig()
    expect(Save).toHaveBeenCalledWith(
      expect.objectContaining({ send_port: 8888 }),
    )
  })

  it('does not save when config is null', async () => {
    const store = useOptionsStore()
    await store.saveConfig()
    expect(Save).not.toHaveBeenCalled()
  })
})
