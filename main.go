package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	// カレントディレクトリをデフォルトとして設定し、
	// コマンドライン引数でディレクトリが指定されている場合はそのディレクトリを使用する
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	// 指定されたディレクトリ内のファイル・ディレクトリ一覧を取得する
	files, err := os.ReadDir(dir)
	// エラーハンドリング: ファイル一覧を取得できなかった場合、エラーメッセージを表示して終了する
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// 取得したファイル一覧をループで処理する
	for _, file := range files {
		fileName := file.Name()
		// ファイル名の末尾が ".txt" である場合は赤文字で表示、そうでない場合は通常の色で表示
		if strings.HasSuffix(fileName, ".txt") {
			color.Red(fileName)
		} else {
			fmt.Println(fileName)
		}
	}
}
