# CHATInput
VRCHATのChatBox用入力補助プログラムです、TTS（読み上げ）機能付属しております。

## 使い方
- 基本利用
    1. BoothかReleaseからバイナリーファイルをダウンロードする、あるいは自分でソースコードからビルドしてバイナリーファイルを取得する。
    2. `ChatInput.exe`を実行する。
    3. オプションタブのオプションをチェックする。
    4. Inputタブでなにか入力してSend押すかキーボードのEnter押せば正常に送信されるはずです。
- TTS（読み上げ）
    - Depsの準備
        1. [VB-Cable](https://vb-audio.com/Cable/)をインストールする。
    - 使用
        1. ChatInputのオプションタブのTTS項目をOnにして、適当になにか入力して送信してWindowsの音量ミキサーで`ChatInput.exe`の出力デバイスを`CABLE Input`に設定する。
        2. VRCHATでマイクデバイスを`CABLE Output`に設定する。
        3. ほかは基本利用とおなじです
## 問題

## TODO
- RealTime Sendの実装
- Typing Messageの実装
- 出力デバイス自動と手動的に選択する機能の追加

## 注意事項
- これは私が仕事の合間に作ったおもちゃ程度のプロジェクトなので、色々大雑把です。
- バグの報告以外はIssuesを使わないでください。
- 問題の修正はこちらが暇なときにだけさせていただきます。
- 私が使いやすいように仕込んでるので、ほかの人にとっては使いづらい可能性があります。

## ライセンス
本プロジェクトは [GNU一般公衆ライセンス3.0](https://www.gnu.org/licenses/gpl-3.0.ja.html) を基づき発行しております、ユーザーには使用、二次開発、二次配布などの自由があります。ただし二次開発する場合など本プロジェクトを利用していることの声明と二次開発後のソースコードを同じ [GNU一般公衆ライセンス3.0](https://www.gnu.org/licenses/gpl-3.0.ja.html) でオプンソースすることが必要です。具体的なライセンスの内容は [GNU一般公衆ライセンス](https://www.gnu.org/licenses/gpl-3.0.ja.html) をご覧ください。

```
Copyright (C) 2025  Sayuki

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
```

# Thanks
- [wails](https://wails.io/)
- [VoiceVox Engine](https://github.com/VOICEVOX/voicevox_engine)
