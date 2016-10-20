## Timely

Timely is the app to get the timely info for the web developers.

## Run

```
$ go build
$ ./timely
```

## Commands

You can fetch the timely info by the simple commands.

- `$ timely all` or `$ timely a`: Show all timely info
- `$ timely github` or `$ timely g`: Show Github trending repositories
- `$ timely qiita` or `$ timely q`: Show Qiita.com's popular items
- `$ timely hatena` or `$ timely h`: Show Hatena Bookmark's hot entries
- `$ timely twitter` or `$ timely t`: Fetch twitter's most shared links

When you find the article you want to see more detail, input the number
- `$ timely qiita 3` or `$ timely q 3`: Open the 3rd link in web browser

## Configuration


## Help

```bash
$ timely
NAME:
   timely - get an timely info

USAGE:
   timely [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     all, a      Show all timely info
     qiita, q    Get timely info from qiita
     github, g   Get timely info from github
     hatena, h   Get timely info from hatena bookmark
     twitter, t  Get timely shared url from twitter
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Example

```bash
$ timely g
--- Github Trending Repositories ---
1 /attic-labs/noms : The versioned, forkable, syncable database
2 /naptha/tesseract.js : Pure Javascript OCR for 62 Languages 📖🎉🖥
3 /channelcat/sanic : Python 3.5+ web server that's written to go fast
4 /yarnpkg/yarn : 📦🐈 Fast, reliable, and secure dependency management.
5 /FreeCodeCamp/FreeCodeCamp : The https://FreeCodeCamp.com open source codebase and curriculum. Learn to code and help nonprofits.
6 /open-guides/og-aws : 📙 Amazon Web Services — a practical guide
7 /jwasham/google-interview-university : A complete daily plan for studying to become a Google software engineer.
8 /Microsoft/LightGBM : LightGBM is a fast, distributed, high performance gradient boosting (GBDT, GBRT, GBM or MART) framework based on decision tree algorithms, used for ranking, classification and many other machine learning tasks.
9 /ZuzooVn/machine-learning-for-software-engineers : A complete daily plan for studying to become a machine learning engineer.
10 /arthepsy/ssh-audit : SSH server auditing (banner, key exchange, encryption, mac, compression, compatibility, security, etc)
11 /alibaba/AliSQL : AliSQL is a MySQL branch originated from Alibaba Group. Fetch document from Release Notes at bottom.
12 /haadcode/orbit : Distributed, serverless, peer-to-peer chat application on IPFS
13 /mike-schultz/materialette : Materialette - A color palette for material design
14 /mholt/caddy : Fast, cross-platform HTTP/2 web server with automatic HTTPS
15 /amark/gun : A realtime, decentralized, offline-first, graph database engine.
16 /FormidableLabs/nodejs-dashboard : Telemetry dashboard for node.js apps from the terminal!
17 /mikeal/roll-call : ☎️ Free and reliable audio calls for everyone w/ browser p2p.
18 /florent37/DiagonalLayout : With Diagonal Layout explore new styles and approaches on material design
19 /yanzhenjie/NoHttp : Android Http标准框架，底层OkHttp，与RxJava完美结合，比Retrofit更简单易用。
20 /tensorflow/tensorflow : Computation using data flow graphs for scalable machine learning
21 /chemzqm/wept : 微信小程序实时开发工具
22 /transcranial/keras-js : Run trained Keras models in the browser, with GPU support
23 /WhiteHouse/fb_messenger_bot : This Drupal module provides a tool to build a chat bot to work on Facebook's Messenger Platform.
24 /sentsin/layui : 经典模块化前端框架
25 /joshwcomeau/redux-vcr : 📼 Record and replay user sessions

$ timely g 4
Opening...
Item: 4 URL: https://github.com/yarnpkg/yarn
```

```bash
$ timely q
--- Qiita 人気の投稿 ---
1   Swift3 全予約語 (95語) の解説
2   日本語で読める GitBook まとめ
3   Serverless Framework v1.0の使い方まとめ
4   アニメーションアイコンが作れるAndroid Icon Animatorを触ってみる
5   Facebook製のJavaScriptテストツール「Jest」の逆引き使用例
6   【Hubot】毎朝決まった時間に、使う電車の運行情報をSlackに投稿するBot
7   老害と言われないためのECMA6勉強会報告・イントロダクション編
8   【メモ】個人的に見落としていたCSSの使い方
9   Swiftで始めるFirebase入門
10   本番をFTPで直接触りたい人たちと共生するためのGitHub同期ツールを作りました
11   android-apt becomes obsolete
12   新形態素解析器JUMAN++を触ってみたけど思ったより高精度でMeCabから乗り換えようかと思った話
13   複数ノードのパケットキャプチャファイルからシーケンス図を出力するツールを作成した
14   【Rails4, OmniAuth】世界一丁寧なLINE LOGIN導入講座
15   REST API設計者のための有名APIのURL例
16   AbemaTV の番組表リニューアルに伴うパフォーマンス改善
17   究極のIT系最新技術情報収集用Slackチーム公開 - モヒカンSlack -
18   Javaで業務系システムを開発するときの鉄板構成(2015年12月版)
19   TensorFlowのチュートリアルを通して、人工知能の原理について学習する
20   Homebrew使い方まとめ

$ timely q 3 # Open the 3rd link
Opening...
URL: http://...
```

```bash
$ timely h
--- はてなブックマーク 技術ブログ ホットエントリー ---
1 JavaScript 製ファミコンエミュレータを公開しました - まるまるこふこふ : 390 users
2 【レビュー】レゴWeDo 2.0で幼児からロボット&プログラミング学習！ - やめて... : 198 users
3 Deep Learning環境はKeras + Docker + Jupyter Notebook + GPUがいいカンジ - ... : 186 users
4 Brainfuck interpreter in Ruby's Regexp - 兼雑記 : 61 users
5 hagino3000's blog: Spotify社のエンジニア評価制度 : 49 users
6 Unixプロセスとシグナルの基礎をRubyで再確認した - えいのうにっき : 42 users
7 もっと瞬殺で作るMesos + Marathon + Dockerクラスタ環境 on Azure - Mitsuyuk... : 33 users
8 はてなブログの高速化（魔改造） - jQueryとJavascriptコードのフッター寄せ -... : 26 users
9 RNNLMベースの形態素解析器 JUMAN++ をhomebrewでインストールできるようにし... : 35 users
10 JavaScript 祭で発表してきました - 若き JavaScripter の悩み : 28 users
11 AndroidのCIはlintOptions.textReportを有効にしておくと捗る - Islands in th... : 25 users
12 Google Mapのオフラインエリアが海外旅行で非常に役立った件 - $shibayu36->blog; : 30 users
13 Hour of Codeとパズル型プログラミング教材の功罪 – Medium : 11 users
14 GPUでTensorFlowを動かす - もょもとの技術ノート : 10 users
15 "使える"プロトタイプ主導の開発プロセス - クックパッド開発者ブログ : 8 users
16 ECSにバッチ処理やらせるときにhako oneshotが便利 - パラボラアンテナと星の日記 : 8 users
17 Visual Studio 2015 がかなり重い、もしくはデバッグ中にOSがフリーズ の対処... : 7 users
18 Reactコンポーネントを純粋な関数で書こう - ushumpei’s blog : 6 users
19 Windows Server 2016 の Docker から Hyper-V Containers を利用する - しばや... : 5 users
20 アイスランド旅行記 - 町並み編 - $shibayu36->blog; : 5 users
21 数学 - JavaScript - パスカルの三角形と剰余 - 係数の相互関係 -剰余による分... : 4 users
22 go vet で panic - tocsatoの備忘録 : 4 users
23 Storing a Sparse Table with O(1) Worst Case Access Timeを読んだ - Echizen... : 4 users
24 Shibuya.XSS techtalk #8 開催にあたり、発表者を募集します！ - 葉っぱ日記 : 4 users
25 npmに代わるNodeパッケージマネージャーyarn - TOEIC940点の文系大学生プログ... : 3 users
```

```bash
$ timely t
--- Twitter most shared links from the engineer accounts ---
Fetch Tweets: mizuno_takaaki
Fetch Tweets: yukihiro_matz
Fetch Tweets: rob_pike
Fetch Tweets: dhh
Fetch Tweets: naoya_ito
Fetch Tweets: takoratta
Fetch Tweets: masuidrive
Fetch Tweets: yugui
Fetch Tweets: mizchi
Fetch Tweets: Jxck_
Fetch Tweets: t_wada
Fetch Tweets: miyagawa
Fetch Tweets: fshin2000
Fetch Tweets: fujiwara
Fetch Tweets: mattn_jp
Fetch Tweets: songmu
Fetch Tweets: yaotti
Fetch Tweets: kuranuki
1 https://t.co/0t3MN3HSXy : 1 times
2 https://t.co/sd4178PYXO : 1 times
3 https://t.co/uOcHdAiK2N : 1 times
4 https://t.co/w20Pj7YaeH : 1 times
5 https://t.co/ZCU4SstdRs : 1 times
6 https://t.co/kUo7nw8LvD : 1 times
7 https://t.co/83qcgob5n2 : 1 times
8 https://t.co/jWk4FU7qap : 1 times
9 https://t.co/VU7f5kCBKI : 1 times
10 https://t.co/EsUhDTpOJm : 1 times
```
