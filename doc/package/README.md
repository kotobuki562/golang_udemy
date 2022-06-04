# crypt 詳細

sha ver

```main.go
package main

import (
"fmt"
"io"

    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"

)

func main() {
//------------------------------
/_ SHA-1 _/
s1 := sha1.New()
io.WriteString(s1, "ABCDE")
fmt.Printf("%x\n", s1.Sum(nil))
// => "7be07aaf460d593a323d0db33da05b64bfdcb3a5"

    //------------------------------

    /* SHA-256 */
    s256 := sha256.New()
    io.WriteString(s256, "ABCDE")
    fmt.Printf("%x\n", s256.Sum(nil))
    // => "f0393febe8baaa55e32f7be2a7cc180bf34e52137d99e056c817a9c07b8f239a"

    //------------------------------

    /* SHA-512 */
    s512 := sha512.New()
    io.WriteString(s512, "ABCDE")
    fmt.Printf("%x\n", s512.Sum(nil))
    // => "9989a8fcbc29044b5883a0a36c146fe7415b1439e995b4d806ea0af7da9ca4390eb92@<dtp>{lb}a604b3ecfa3d75f9911c768fbe2aecc59eff1e48dcaeca1957bdde01dfb"

}
```

# net/http client の応用

ヘッダーを付けたり、クエリを付けたりする場合。

Get ver

```main.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//Get

func main() {
	//応用
	//ヘッダーをつけたり、クエリをつけたり
	//Parse 正しいURLか確認
	base, _ := url.Parse("https://example.com/")

	//クエリ の確認 URLの後につく
	reference, _ := url.Parse("index/lists?id=1")

	//ResolveReference
	//クエリをくっつけたURLを生成する。
	//相対パスから絶対URLに変換する。
	//baseのURLの末尾に文字列が入っていたとしても、正しいURLに直してくれる
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	//GET ver
	//リクエストを作成 nil部はPOST時のみ設定（バイトを入れる）
	//まだリクエストはしていない。structを作っただけ。
	req, _ := http.NewRequest("GET", endpoint, nil)

	//requestにheaderをつける。cash情報など
	req.Header.Add("Content-Type", `application/json"`)

	//URLのクエリを確認
	q := req.URL.Query()

	//クエリを追加
	q.Add("name", "test")

	//クエリを表示
	fmt.Println(q)

	//&など特殊文字などがある場合があるため、encodingしてからURLに追加してやる必要がある
	fmt.Println(q.Encode())

	//encodeしてからURLに戻す
	//日本語などを変換する
	req.URL.RawQuery = q.Encode()

	//実際にアクセスする
	//クライアントを作る
	var client *http.Client = &http.Client{}

	//結果 アクセスする
	resp, _ := client.Do(req)

	//読み込み
	body, _ := ioutil.ReadAll(resp.Body)

	//出力
	fmt.Println(string(body))

}
```

Post ver

```main.go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

//応用
type Account struct {
    ID       string
    PassWord string
}

//Post

func main() {
    //応用
    //ヘッダーをつけたり、クエリをつけたり
    //Parse 正しいURLか確認
    base, _ := url.Parse("https://example.com/")

    //クエリ の確認 URLの後につく
    reference, _ := url.Parse("index/lists?id=1")

    //Postの時のデータ
    AccountData := &Account{ID: "ABC-DEF", PassWord: "testtest"}
    data, _ := json.Marshal(AccountData)

    //ResolveReference
    //クエリをくっつけたURLを生成する。
    //相対パスから絶対URLに変換する。
    //baseのURLの末尾に文字列が入っていたとしても、正しいURLに直してくれる
    endpoint := base.ResolveReference(reference).String()
    fmt.Println(endpoint)

    //POST ver
    //bytes.NewBuffer([]byte("password")でリクエストの領域を作成
    //POSTの場合は、Bodyにデータを入れる。例えばパスワード。見られたらダメな情報はbodyに
    req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))

    //requestにheaderをつける。cash情報など
    req.Header.Add("Content-Type", `application/json"`)

    //URLのクエリを確認
    q := req.URL.Query()

    //クエリを追加
    q.Add("name", "test")

    //クエリを表示
    fmt.Println(q)

    //&など特殊文字などがある場合があるため、encodingしてからURLに追加してやる必要がある
    fmt.Println(q.Encode())

    //encodeしてからURLに戻す
    //日本語などを変換する
    req.URL.RawQuery = q.Encode()

    //実際にアクセスする
    //クライアントを作る
    var client *http.Client = &http.Client{}

    //結果 アクセスする
    resp, _ := client.Do(req)

    //読み込み
    body, _ := ioutil.ReadAll(resp.Body)

    //出力
    fmt.Println(string(body))


```
