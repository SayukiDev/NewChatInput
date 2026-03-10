export default {
  app: {
    connecting: 'Connecting to backend...',
    tabs: { chat: 'Chat', options: 'Options', tts: 'TTS', log: 'Log' },
  },
  chat: {
    title: 'Chat',
    input: { placeholder: 'Type a message...', send: 'Send' },
    fullInput: {
      toastSummary: 'Full Input Mode',
      toastDetail: 'Press Esc to exit.',
    },
  },
  options: {
    sections: { oscPort: 'OSC Port', settings: 'Settings', tts: 'TTS' },
    toggles: {
      typingIndicator: 'Typing Indicator',
      realtime: 'Realtime',
      messageKeeping: 'Message Keeping',
      voiceControl: 'Voice Control',
      enableTts: 'Enable TTS',
    },
    saveButton: 'Save Settings',
    toast: {
      savedSummary: 'Saved',
      savedDetail: 'Settings saved successfully.',
      errorSummary: 'Error',
      errorDetail: 'Failed to save settings.',
    },
    port: { send: 'Send Port', receive: 'Receive Port' },
    speaker: {
      label: 'Speaker',
      placeholder: 'Select a speaker',
      styleLabel: 'Style',
      stylePlaceholder: 'Select a style',
    },
    audioDevice: { label: 'Audio Device', placeholder: 'Select an audio device' },
  },
  log: { title: 'Log', refresh: 'Refresh' },
  tts: {
    sections: { voiceSettings: 'Voice Settings', engineSettings: 'TTS Engine Settings', cacheSettings: 'Cache Settings', log: 'Log' },
    engine: { baseurl: 'Base URL', path: 'Engine Path', args: 'Launch Arguments (space-separated)', log: 'Log Path', run: 'Auto-start Engine', statusRunning: 'Engine is running', statusStopped: 'Engine is stopped' },
    cache: { enable: 'Enable Cache', path: 'Cache Path' },
    disabledNotice: 'TTS is disabled. Enable it from the Options tab.',
    install: {
      button: 'Install Speaker',
      successSummary: 'Install Complete',
      successDetail: 'Speaker installation completed.',
      errorSummary: 'Error',
      errorDetail: 'Failed to install speaker.',
    },
  },
} as const
