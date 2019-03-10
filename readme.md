## delve使い方

基本的な操作は下記

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

下記プロビジョニングツールでVMを管理する

- vagrant
- ansible


```
vagrant box add centos/7
vagrant init centos/7
```

上記コマンドで環境構築をする


https://qiita.com/pugiemonn/items/bcd95a35c3ec7624cd61
https://qiita.com/0ta2/items/69c2378600bf8adaac5c
https://qiita.com/ringo0321/items/ddf1c8c40014eb88b0b6
http://stupiddog.jp/note/archives/1148
https://github.com/luckyriver0/vagrant/tree/master/golang/vagrant
https://xn--o9j8h1c9hb5756dt0ua226amc1a.com/?p=512

```
vagrant up
```

