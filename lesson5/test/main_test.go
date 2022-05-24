// mainパッケージのテストを行う
package main

// main_test.goのように、テスト用のファイルを作成することで、
// testingをインポートする
import "testing"

// テストファイルのあるディレクトリで`go test`でテストを実行する
// `go test -v`で詳細な情報を表示する
var Debug bool = false

func TestIsOne(t *testing.T) {
	i := 1
	if Debug {
		t.Skip("スキップさせる")
	}

	v := IsOne(i)

	if !v {
		t.Errorf("%v A= %v", i, 1)
	}
}

/*
成功すると
PASS
ok      golang_udemy/lesson5/test       0.223s
*/

/*
失敗すると
FAIL
exit status 1
FAIL    golang_udemy/lesson5/test       0.211s
*/