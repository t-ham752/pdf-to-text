## pdf-to-text
setup
```shell
$ make build && make up
```

extract text from pdf
```shell
$ go run . sample.pdf
output: 山田 太朗
プログラマー
050-1234-5678 | taro.yamada@example.com

| 000-0000 東京都渋谷区

要約
私はこれまでお仕事を頑張ってきました。これからも頑張ります。

資格
基本情報技術者試験

2024 年
```
