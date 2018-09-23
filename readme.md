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

delveの起動後の詳しい操作は `help` コマンドで見れる
