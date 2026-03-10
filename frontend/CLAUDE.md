# ChatInput-Frontend

VRChat ChatBox送信アプリ（OSC + TTS読み上げ）のWails v2フロントエンド。

## Tech Stack
- Vue 3 + TypeScript + Vite 7
- PrimeVue 4 (Aura Dark theme) + unplugin-vue-components (auto-import)
- Pinia 3 (state management)
- Vue Router 4 (hash mode)
- Vitest + @vue/test-utils + happy-dom

## Commands
- `npm run dev` — 開発サーバー起動
- `npm run build` — 型チェック + ビルド
- `npm run test` — Vitestでテスト実行
- `npm run preview` — ビルド成果物プレビュー

## Architecture
- `src/stores/` — Piniaストア（app, chat, options, tts）
- `src/views/` — ルートビュー（ChatView, OptionsView）
- `src/components/` — UIコンポーネント
- `wailsjs/` — 自動生成のWailsバインディング（編集不可）
- `wailsjs/go/pages/` — バックエンドAPI（App, Input, Options）
- `wailsjs/go/models.ts` — 型定義（options.Config, pages.AudioDeviceRsp等）

## Restrictions
- バックエンド（`../`以下のGoコード）は読み取り・編集ともに禁止
- 作業範囲は `frontend/` ディレクトリ内のみ

## Conventions
- PrimeVueコンポーネントはauto-importされるため明示的importは不要
- Wails APIの呼び出しはストア内でのみ行い、コンポーネントから直接呼ばない
- Vue Router hashモード使用（Wailsのファイルベース読み込みに対応）
- テストでは`wailsjs/`のバインディングをvi.mockでモック
