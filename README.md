# tk-tinygo (tinygoを利用してraspberry pi picoを動かす)
### 準備（Windows)
1. [Scoop](https://scoop.sh/)をダウンロードしてインストール
2. `scoop install go`: goのインストール
3. `scoop install tinygo`: tinygoのインストール
4. `scoop install avr-gcc`: ardiuno依存環境をインストール（その１）
5. `scoop install avrdude`: ardiuno依存環境をインストール（その２）

6. [Visual Studio Code](https://code.visualstudio.com/)をダウンロードしてインストール
7. [extensions.json](.vscode/extensions.json)へ記載した拡張機能を追加
8. [settings.json](.vscode/settings.json)内の`go.toolsEnvVars`設定値の取得
    1. `tinygo info pico`: picoの情報取得
    2. `build tags`及び`cached GOROOT`の値をコピーして貼り付け（`build tags`の値はカンマ区切りへ変更）
9. Visual Studio Codeのコマンドパレットにて`TinyGo target pico`を実行
10. ここまで実施すると、`machine`パッケージを参照することができるようになり、ビルドエラーが消える