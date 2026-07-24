// BNO 留英日數計算機 — Windows 本地版。
//
// 單一執行檔：把計算機頁面用 embed 打包進來，在本機 127.0.0.1 提供，
// 並自動開啟瀏覽器。純標準函式庫，無 CGO、無外部相依。
// 資料存在瀏覽器本機（localStorage），配合頁面的「匯出／匯入」備份。
package main

import (
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

// 與 GitHub Pages 共用同一個檔：index.html 既是網站首頁，也內嵌進本機版。
//go:embed index.html
var content embed.FS

// 固定連接埠 → 每次啟動來源相同，localStorage 資料才會保留。
const addr = "127.0.0.1:8787"

func main() {
	page, err := content.ReadFile("index.html")
	if err != nil {
		log.Fatal("讀取內嵌頁面失敗：", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		_, _ = w.Write(page)
	})

	// 先開始監聽，確認埠可用，再開瀏覽器。
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("無法在 %s 啟動（可能已有另一個視窗在執行）：%v", addr, err)
	}

	url := "http://" + addr + "/"
	fmt.Println("BNO 留英日數計算機 已啟動：", url)
	fmt.Println("瀏覽器會自動開啟。關閉此視窗即結束程式。")

	go func() {
		time.Sleep(300 * time.Millisecond)
		openBrowser(url)
	}()

	if err := http.Serve(ln, mux); err != nil {
		log.Fatal(err)
	}
}

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	if err := cmd.Start(); err != nil {
		log.Println("自動開啟瀏覽器失敗，請手動開啟：", url)
	}
}
