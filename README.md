## gen wire

`cd cmd && wire gen`

## build

`go run cmd/main.go wire_gen.go`

`cd cmd && go build -o go-di`


### echo について
* e.Groupでrouteのグループ化ができる、それによって、たとえば/adminがつくURIにたいして認証を設置するみたいなこともできるっぽい。
* JSON Blobってのがある。
* Basic認証の仕組み(https://www.atmarkit.co.jp/ait/articles/1608/10/news021.html)
* デベロッパーツールでみたところ、Authenticationヘッダみたいなのがあったので、それで毎回認証通してるっぽい。
* e.File()とかで、静的ファイルとのルーティング管理もできる。e.File("/", "public/index.html")みたいな。
