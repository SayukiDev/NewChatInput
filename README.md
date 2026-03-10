# CHATInput

VRCHATのChatBox用入力補助プログラムです、TTS（読み上げ）機能付属しております。

## 使い方

- 基本利用

  1. バイナリファイル[ダウンロードする](https://github.com/SayukiDev/NewChatInput/releases/latest/download/ChatInputWithTTS.zip)、
     あるいは自分でソースコードからビルドしてバイナリーファイルを取得する。
  2. `ChatInput.exe`を実行する。
  3. チャットタブでなにか入力して送信ボタン押すかキーボードのEnter押せば正常に送信されるはずです。
- TTS（読み上げ）の設定

  1. [VB-Cable](https://vb-audio.com/Cable/)をインストールする。
  2. 読み上げたぶで話者を選択して設定する
  3. オーディオデバイスを`VB-Audio Virtul Cable`に設定する
  4. VRChat内のマイク設定も`VB-Audio Virtul Cable`に設定する


## モデルの追加
読み上げエンジンはAivisSpeech Engineをつかっているので、
現時点ではAIVM / AIVMXフォーマットのStyle-Bert-VITS2モデルのみサポートされています。  
モデルを追加インストールしたい場合は読み上げタブのインストールボタンを押し追加したいモデルファイルを選択すればおけです

#### モデルの取得、変換など
- Style-Bert-VITS2モデルはhuggingfaceやBoothでたくさん配布あるいは販売されてるので、それらのサイトから取得できます。
- Style-Bert-VITS2デフォルートのフォーマットは使えないので、変換する必要があります、AivisSpeech Engineから提供されてる[AIVM Generator](https://aivm-generator.aivis-project.com/)で簡単にStyle-Bert-VITS2モデルから変換できます。

## 問題

## TODO

- RealTime Sendの実装

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
- [AivisSpeech](https://github.com/Aivis-Project/AivisSpeech)
