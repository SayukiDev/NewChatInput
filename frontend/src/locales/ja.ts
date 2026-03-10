export default {
  app: {
    connecting: 'バックエンドに接続中...',
    tabs: { chat: 'チャット', options: '設定', tts: '読み上げ', log: 'ログ' },
  },
  chat: {
    title: 'チャット',
    input: { placeholder: 'メッセージを入力...\n\nTips:\nEnterキーで送信\nDeleteキーでクリア', send: '送信' },
    fullInput: {
      toastSummary: 'フル入力モード',
      toastDetail: 'Escキーで終了します。',
    },
  },
  options: {
    sections: { oscPort: 'OSCポート', settings: '設定', tts: 'TTS' },
    toggles: {
      typingIndicator: 'タイピングインジケーター',
      realtime: 'リアルタイム',
      messageKeeping: 'メッセージ保持',
      voiceControl: '音声コントロール',
      enableTts: '読み上げを有効にする',
    },
    saveButton: '設定を保存',
    toast: {
      savedSummary: '保存しました',
      savedDetail: '設定を保存しました。',
      errorSummary: 'エラー',
      errorDetail: '設定の保存に失敗しました。',
      restartWarnSummary: '注意',
      restartWarnDetail: '一部の変更は再起動する必要があります。',
    },
    port: { send: '送信ポート', receive: '受信ポート' },
    speaker: {
      label: '話者',
      placeholder: '話者を選択',
      styleLabel: 'スタイル',
      stylePlaceholder: 'スタイルを選択',
    },
    audioDevice: { label: 'オーディオデバイス', placeholder: 'オーディオデバイスを選択' },
  },
  log: { title: 'ログ', refresh: '更新' },
  tts: {
    sections: { voiceSettings: '音声設定', engineSettings: '読み上げエンジン設定', cacheSettings: 'キャッシュ設定', log: 'ログ' },
    engine: { baseurl: 'ベースURL', path: 'エンジンパス', args: '起動引数（スペース区切り）', log: 'ログパス', run: 'エンジン自動起動', statusRunning: 'エンジンは実行されています', statusStopped: 'エンジンは停止しています' },
    cache: { enable: 'キャッシュを有効にする', path: 'キャッシュパス' },
    disabledNotice: '読み上げが無効です。設定タブから読み上げを有効にしてください。',
    install: {
      button: '話者をインストール',
      successSummary: 'インストール完了',
      successDetail: '話者のインストールが完了しました。',
      errorSummary: 'エラー',
      errorDetail: '話者のインストールに失敗しました。',
      restartWarnDetail: 'インストールされた話者を使うには再起動する必要があります。',
    },
  },
} as const
