# go_sys

golang system programs

## getting started

下記プロビジョニングツールでVMを管理する

- vagrant
- ansible

上記コマンドで環境構築をする

```
vagrant up
```

## delve使い方

delveでシステムコールを覗いたりする
delveの基本的な操作は下記

```
dlv debug デバッグプロセスを起動する
break main.main mainパッケージのmain関数にブレークポイントを仕込む
continue ブレークポイントまで実行する
next 一行づつ実行する
step ステップイン
stepout ステップアウト
exit デバッグプロセスを終了する
restart ブレークポイントを残したままデバッグプロセスを再起動する
```

デバッグプロセス中の詳しい操作は `help` コマンドで見れる

## 構成


## memo

https://github.com/yurakawa/learn-system-programming-with-go
