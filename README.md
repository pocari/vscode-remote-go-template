## goの標準ディレクトリ構造みたいなのがあるらしくて、それに合わせたリポジトリ

https://github.com/pocari/vscode-remote-go-standard-dir-layout-template


# goのプロジェクトでのvscodeのremote設定サンプル

- remote側でのbreakpoint、デバッグ実行ができる
- remote側でlanguage server(gopls)を実行して補完ができる
- サンプルとしてapi serverを実装しており、realizeでのhot reloadができる

など

# 始め方
前提、Remote Containers関連のpluginがすでにインストールされていること

## 適当な場所にcloneする

```
git clone git@github.com:pocari/vscode-remote-go-template.git
```

## vscode側

vscodeを開いて、 F1 から
`Remote Containers: Open Folder in Container`
を実行して、先程cloneしてきたディレクトリを指定する

すると、初回実行時は少し待っていると (`docker build` しているため) remote側のディレクトリを開いた状態で vscode が再起動される。
この状態で、ターミナルを追加すると、ホスト側でなく、コンテナ上で新規にシェルのセッションが起動され、コンテナ内の作業ができる。

## その他

- デバッグ実行
    - 普通にホスト側で開発するときのように ブレイクポイントを設定して、F5 で起動するとAPI serverが立ち上がる

- hot reload
    - ターミナルから `make start` すると、 realizeで管理されたhot reloadモードのapi serverが立ち上がる

