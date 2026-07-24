# BNO-VISA-Day-Tracker

**BNO 簽證留英日數計算機** —— 追蹤你每次離開英國的日數，檢查永久居留（ILR）「在任何連續 12 個月內離境不超過 180 天」的規則，並估算 5 年資格與最早申請日期。

**BNO visa UK-residence day tracker** — track your days of absence from the UK, check the settlement (ILR) rule of *no more than 180 days outside the UK in any rolling 12-month period*, and estimate your 5-year eligibility and earliest application dates.

離線可用 · 資料只留在你自己的瀏覽器 · 無需帳戶
Works offline · your data stays in your own browser · no account needed

---

## 目錄 · Contents

- [如何使用 · How to Use](#如何使用--how-to-use)
- [功能總覽 · Features](#功能總覽--features)
- [使用說明 · Usage Guide](#使用說明--usage-guide)
- [計算規則詳解 · Calculation Rules](#計算規則詳解--calculation-rules)
- [資料與私隱 · Data & Privacy](#資料與私隱--data--privacy)
- [編譯與建置 · Building (Windows / macOS / Linux)](#編譯與建置--building-windows--macos--linux)
- [技術說明 · Technical Notes](#技術說明--technical-notes)
- [免責聲明 · Disclaimer](#免責聲明--disclaimer)

---

## 如何使用 · How to Use

有兩種方式，任選一種。/ Two ways to use it — pick either one.

### 1. 直接執行網頁 · Run it in your browser

整個工具就是**一個自足的網頁**（單一 HTML，內含全部 CSS/JavaScript），不需安裝、不需後端。/ The whole tool is a **single self-contained web page** (one HTML file with all CSS/JS inline) — nothing to install, no backend.

- **線上版**：開啟已部署的網址（GitHub Pages 等）：
  `https://woonnngggg.github.io/BNO-VISA-Day-Tracker/`
  **Online:** open the deployed URL, e.g.
  `https://woonnngggg.github.io/BNO-VISA-Day-Tracker/`
- **或**下載單一檔 `index.html`，用任何瀏覽器直接開啟（雙擊即可）。
  **Or** download the single `index.html` and just open it in any browser (double-click).

> 想自行部署到 GitHub Pages（免費、附 HTTPS），步驟見 **[INSTALL.md](INSTALL.md)**。
> To host your own free copy on GitHub Pages, see **[INSTALL.md](INSTALL.md)**.

### 2. 複製或下載（桌面程式）· Download or copy (desktop app)

桌面版是一個把**同一個網頁**打包進去的**單一執行檔**（以 Go 製作），雙擊即在本機開啟，完全離線。/ The desktop version is a **single executable** (built with Go) that bundles **the same page** and opens it locally, fully offline.

- **下載**：取得對應平台的執行檔並直接執行 ——
  Windows（`BNO-VISA-Day-Tracker.exe`）、macOS、Linux 皆可。
  **Download** the binary for your platform (Windows / macOS / Linux) and run it.
- **複製原始碼自行編譯**：`git clone` 之後執行 `go build`（見
  [編譯與建置](#編譯與建置--building-windows--macos--linux)）。
  **Copy the source and build:** `git clone`, then `go build` (see
  [Building](#編譯與建置--building-windows--macos--linux)).

---

## 功能總覽 · Features

| 區塊 · Section | 功能 · What it does |
|---|---|
| **設定 · Settings** | 簽證獲批日期、首次入境日期、永居 5 年起算日（獲批日／入境日）、離境計算方式（標準／保守）<br>Grant date, first-entry date, 5-year start basis (grant / entry), counting method (standard / conservative) |
| **摘要儀表板 · Dashboard** | 總離境日數、目前 12 個月已用、最壞 12 個月窗口、永居資格日期<br>Total absence, current 12-month usage, worst 12-month window, eligibility date |
| **離境限額狀態 · Limit status** | 最壞窗口 vs 180 日進度條、狀態標籤、最壞窗口起訖日<br>Worst window vs the 180-day limit, status pill, worst-window dates |
| **永居申請時程 · ILR timeline** | 住滿 5 年（資格日）、最早可提交申請日（資格日前 28 天內）<br>5-year completion date and earliest application date (within 28 days before) |
| **行程紀錄 · Trip log** | 新增／編輯／刪除行程；每程顯示離境日數與「任何 12 個月最多離英」<br>Add/edit/delete trips; each shows its absence days and worst 12-month exposure |
| **進階分析 · Analysis** | 門檻預警清單、單日／範圍逐日查詢<br>Threshold alert list, single-day / date-range lookup |
| **資料 · Data** | 匯出／匯入 JSON、清除全部<br>Export / import JSON, clear all |
| **介面 · Interface** | 繁體中文、深／淺色主題切換、離線可用<br>Traditional Chinese UI, light/dark theme, works offline |

---

## 使用說明 · Usage Guide

### 1. 設定 · Settings

- **BNO 簽證獲批日期 / Grant date**：簽證批出（生效）日。The date the visa was granted (took effect).
- **首次入境英國日期 / First-entry date**：你實際首次入境的日期。The date you first physically entered the UK.
- **永居 5 年起算日 / 5-year start basis**：選擇 5 年由**簽證獲批日**或**首次入境日**起算；
  此選擇同時影響「獲批至入境的空檔是否計入離境」（見[計算規則](#5-年起算日與獲批入境空檔--start-basis--the-grant-to-entry-gap)）。
  Choose whether the 5 years count from the **grant date** or the **first-entry date**;
  this also controls whether the grant-to-entry gap counts as absence.
- **離境日數計算方式 / Counting method**：**標準** 或 **保守**（見[計算規則](#離境日數計算方式標準保守--counting-method-standard--conservative)）。
  **Standard** or **conservative**.

### 2. 摘要儀表板 · Dashboard

四張卡片一眼掌握現況：總離境日數、目前 12 個月已用 / 180、最壞 12 個月窗口、永居資格日期（含倒數）。
Four cards at a glance: total absence, current 12-month usage / 180, worst 12-month window, and eligibility date (with countdown).

### 3. 離境限額狀態 · Limit status

進度條顯示最壞 12 個月窗口對 180 日上限的比例，並標示狀態：🟢 符合（< 150 日）、🟡 接近上限（150–180 日）、🔴 超出上限（> 180 日）。下方顯示離境最多的那段 12 個月期間起訖日。
A meter shows the worst 12-month window against the 180-day limit, with a status: 🟢 OK (< 150), 🟡 near limit (150–180), 🔴 exceeded (> 180). Below it, the start/end of that worst 12-month period.

### 4. 永居申請時程 · ILR timeline

- **住滿 5 年（資格日）/ 5-year completion**：起算日 + 5 年，含倒數。Start date + 5 years, with a countdown.
- **最早可提交申請 / Earliest application**：資格日**前 28 天**。28 days before the completion date.

### 5. 行程紀錄 · Trip log

在「加入離英行程」填入離境日期、返英日期、可選備註。表格每列顯示：離境、返英、備註、**離境日數**、**任何 12 個月最多離英**、編輯／刪除。
In "add trip", enter departure/return dates and an optional note. Each row shows: departure, return, note, **absence days**, **worst 12-month exposure**, edit/delete.

- **離境日數 / Absence days**：該程本身的離英天數。The trip's own days outside the UK.
- **任何 12 個月最多離英 / Worst 12-month exposure**：在**包含該程**的任何連續 12 個月期間中，離英最多的日數（向前、向後各看 12 個月取最大值）—— 即該程在 180 日規則下的最壞曝險。
  The maximum absence in any rolling 12-month window that **includes this trip** (looking both 12 months forward and backward) — the trip's worst exposure to the 180-day rule.

### 6. 進階分析 · Analysis

- **門檻預警 / Threshold alert**：輸入日數門檻（預設 180，可調低作安全預警），列出所有「過去 12 個月離英達到／超過該門檻」的日期區間及各段最高值。
  Enter a day threshold (default 180; lower it as an early warning) to list all date ranges where the past-12-month absence meets/exceeds it, with each peak.
- **逐日 / 範圍查詢 / Per-day lookup**：只填開始日 → 顯示截至該日的過去 12 個月離英；填開始與結束 → 標出範圍內離英最多的一天並逐日列出。
  One date → the past-12-month absence as of that day; a range → the worst day within it plus a per-day listing.

### 7. 資料 · Data

- **匯出 JSON / Export JSON**：把全部設定與行程存成備份檔。Save all settings and trips to a backup file.
- **匯入 JSON / Import JSON**：從備份檔還原（會覆蓋現有資料）。Restore from a backup (overwrites current data).
- **清除全部 / Clear all**：刪除所有資料（會先確認）。Delete everything (with confirmation).

---

## 計算規則詳解 · Calculation Rules

> 以下為本工具採用的計算邏輯；實際審批以英國內政部指引及個案為準（見[免責聲明](#免責聲明--disclaimer)）。
> The logic below is what this tool uses; actual decisions follow Home Office guidance and your own case (see [Disclaimer](#免責聲明--disclaimer)).

### 180 日 / 12 個月規則 · The 180-day / 12-month rule

核心：**在任何連續 12 個月期間內，離英不得超過 180 日。** 程式對居留期內每一天 `E`，取以 `E` 結尾、往前推整整 12 個曆月的窗口 `[E − 12 個月 + 1 日, E]`，加總窗口內離英日數，取所有窗口的**最大值**即「最壞窗口」。視窗以**日曆月**計算，不是固定 365 日，因此**跨閏日時自然為 366 日**，不會因閏年算錯。

Core rule: **no more than 180 days outside the UK in any rolling 12-month period.** For each day `E`, the tool takes the window `[E − 12 months + 1 day, E]`, sums the absence days inside it, and takes the **maximum** across all windows — the "worst window". The window is **calendar-based**, not a fixed 365 days, so it is naturally **366 days across a leap day** and never miscounts due to leap years.

### 離境日數計算方式（標準／保守）· Counting method (standard / conservative)

一次行程（離境日 `D`、返英日 `R`）的離英天數 / Absence days for a trip (departure `D`, return `R`):

| 方式 · Method | 說明 · Meaning | 天數 · Days | 計入的日子 · Days counted |
|---|---|---|---|
| **標準 · Standard** | 旅行當日視為身處英國 · travel days count as in the UK | `R − D − 1` | `D+1` … `R−1` |
| **保守 · Conservative** | 旅行當日也算離英 · travel days count as absent | `R − D + 1` | `D` … `R` |

官方對「旅行當日」的計法存在不同詮釋，**保守**為較嚴謹的估算。
Interpretations of "travel days" differ; **conservative** is the stricter estimate.

### 5 年起算日與獲批入境空檔 · Start basis & the grant-to-entry gap

由「永居 5 年起算日」設定決定，兩種模式一致地影響資格日與空檔計算：
Set by the "5-year start basis" option; both modes affect the eligibility date and the gap consistently:

| 起算日 · Basis | 資格日期 · Eligibility | 獲批 → 入境的空檔 · Grant→entry gap |
|---|---|---|
| **簽證獲批日 · Grant date** | 獲批日 + 5 年 · grant + 5y | **計入離境**（入境前屬離境；範圍 `[獲批日, 入境前一日]`）· **counts as absence** (`[grant, day before entry]`) |
| **首次入境日 · Entry date** | 入境日 + 5 年 · entry + 5y | **不計**（居留由入境日才開始）· **not counted** (residence starts at entry) |

設定區的動態提示會顯示目前空檔天數與其處理方式；若選「獲批日」且空檔**超過 180 日**，會以琥珀色警示連續居留可能中斷，可考慮改由入境日起算。
A live hint shows the current gap and how it is treated; if "grant date" is chosen and the gap **exceeds 180 days**, it warns (amber) that continuous residence may be broken and suggests switching to the entry date.

### 永居資格與申請日期 · Eligibility & application dates

- **資格日（住滿 5 年）/ Eligibility (5 years)** ＝ 起算日 + 5 年（以日曆計算，閏日如 2 月 29 日正規化為 3 月 1 日）。Start date + 5 years (calendar-based; a leap day such as Feb 29 normalises to Mar 1).
- **最早可提交申請 / Earliest application** ＝ 資格日 − 28 日。Eligibility date − 28 days.

### 每程「任何 12 個月最多離英」· Per-trip worst 12-month exposure

對每一程，程式掃描**所有包含該程離英日**的連續 12 個月窗口（窗口尾端由該程首日掃到末日 + 12 個月），取離英日數的**最大值**。因此若兩程落在同一個 12 個月窗口內，該欄會顯示兩程**合併後**的總數，反映該程真正的最壞曝險。

For each trip, the tool scans **every rolling 12-month window that contains any of the trip's absence days** (window end swept from the trip's first day to its last day + 12 months) and takes the **maximum**. So if two trips fall in the same 12-month window, this column shows their **combined** total — the trip's true worst exposure.

---

## 資料與私隱 · Data & Privacy

資料只存在**你自己的瀏覽器**（localStorage）。線上版的主機（如 GitHub）只負責送出網頁，**永遠看不到你的行程資料**。不同瀏覽器／裝置不會自動同步，跨裝置請用「匯出／匯入」。

Your data is stored **only in your own browser** (localStorage). The host (e.g. GitHub)merely serves the page and **never sees your travel data**. Different browsers/devices don't sync automatically — use export/import to move data between them.

---

## 編譯與建置 · Building (Windows / macOS / Linux)

需求 · Requires: **Go 1.26+**（<https://go.dev/dl/>）。

因為本專案是**純 Go 標準函式庫、無 CGO**，在任何一台裝有 Go 的電腦上，都能編出各平台的執行檔，不需要目標平台的 C 編譯器或任何額外工具。
Because the project is **pure Go standard library with no CGO**, one machine with Go can build binaries for every platform — no target-platform C compiler or extra tooling needed.

### 1. 為目前這台電腦編譯 · Build for the current machine

```sh
cd BNO-VISA-Day-Tracker

# Windows
go build -o BNO-VISA-Day-Tracker.exe .

# macOS / Linux
go build -o BNO-VISA-Day-Tracker .
```

產生單一執行檔（約 8–9 MB），可複製到同型電腦直接執行，免安裝。開發時可用 `go run .`。
Produces a single ~8–9 MB executable that runs on the same platform with no install. Use `go run .` during development.

### 2. 交叉編譯 · Cross-compilation

用 `GOOS`（作業系統）與 `GOARCH`（架構）指定目標 / Set `GOOS` and `GOARCH` for the target:

| 目標 · Target | GOOS | GOARCH | 建議檔名 · Suggested name |
|---|---|---|---|
| Windows 64-bit | `windows` | `amd64` | `BNO-VISA-Day-Tracker.exe` |
| macOS Apple Silicon (M1–M4) | `darwin` | `arm64` | `BNO-VISA-Day-Tracker-macos-arm64` |
| macOS Intel | `darwin` | `amd64` | `BNO-VISA-Day-Tracker-macos-intel` |
| Linux 64-bit | `linux` | `amd64` | `BNO-VISA-Day-Tracker-linux-amd64` |
| Linux ARM64 | `linux` | `arm64` | `BNO-VISA-Day-Tracker-linux-arm64` |

**Windows PowerShell：**

```powershell
$env:GOOS="darwin"; $env:GOARCH="arm64"; go build -o BNO-VISA-Day-Tracker-macos-arm64 .
$env:GOOS="linux";  $env:GOARCH="amd64"; go build -o BNO-VISA-Day-Tracker-linux-amd64 .
Remove-Item Env:GOOS, Env:GOARCH        # 用完清掉 · reset afterwards
```

**macOS / Linux / Git Bash：**

```bash
GOOS=darwin GOARCH=arm64 go build -o BNO-VISA-Day-Tracker-macos-arm64 .
GOOS=linux  GOARCH=amd64 go build -o BNO-VISA-Day-Tracker-linux-amd64 .
```

一次編出全部平台（bash，輸出到 `dist/`）· Build all platforms at once (bash → `dist/`):

```bash
for t in windows/amd64/.exe darwin/arm64/-macos-arm64 darwin/amd64/-macos-intel linux/amd64/-linux-amd64 linux/arm64/-linux-arm64; do
  IFS=/ read -r os arch suffix <<< "$t"
  GOOS=$os GOARCH=$arch go build -o "dist/BNO-VISA-Day-Tracker$suffix" .
done
```

### 3. 在各平台執行 · Running on each platform

- **Windows**：雙擊執行檔；關閉黑色命令視窗即結束。Double-click; close the console window to quit.
- **macOS**：
  ```bash
  chmod +x BNO-VISA-Day-Tracker-macos-arm64
  ./BNO-VISA-Day-Tracker-macos-arm64        # 終端機 Ctrl+C 結束 · Ctrl+C to quit
  ```
  未簽章的執行檔首次會被 **Gatekeeper** 擋。解法任一：右鍵 →「開啟」；或
  `xattr -d com.apple.quarantine 檔名`；或系統設定 → 隱私權與安全性 →「仍要打開」。
  An unsigned binary is blocked by **Gatekeeper** first time — right-click → Open,
  or `xattr -d com.apple.quarantine <file>`, or allow it in System Settings → Privacy & Security.
- **Linux**：
  ```bash
  chmod +x BNO-VISA-Day-Tracker-linux-amd64
  ./BNO-VISA-Day-Tracker-linux-amd64
  ```
  自動開瀏覽器需要 `xdg-open`（`xdg-utils`）；若無，手動開 <http://127.0.0.1:8787/>。
  Auto-open needs `xdg-open` (`xdg-utils`); otherwise open <http://127.0.0.1:8787/> manually.

### 4. 其他選項 · Other options

- **隱藏 Windows 命令視窗 · Hide the Windows console**：
  `go build -ldflags "-H=windowsgui" -o BNO-VISA-Day-Tracker.exe .`
  （這樣就沒有視窗可關閉，需從「工作管理員」結束 · then quit via Task Manager.）
- **縮小檔案 · Smaller binary**：加 `-ldflags "-s -w"`（去除除錯符號，約省 25% · strips debug symbols, ~25% smaller）。

---

## 技術說明 · Technical Notes

- **語言／執行環境 · Language / runtime**：Go 1.26，僅用標準函式庫（`embed`、`net/http`、`os/exec`）。**無 CGO、無第三方相依。** Pure Go standard library, no CGO, no third-party dependencies.
- **前端 · Front end**：單一 HTML 檔，內嵌 CSS 與原生 JavaScript，無框架、無外部資源，離線可用。One HTML file with inline CSS and vanilla JS — no framework, no external resources, works offline.
- **內嵌打包 · Embedding**：`//go:embed index.html` 把頁面編進 `.exe`，因此只需散佈一個檔；同一個 `index.html` 也是 GitHub Pages 的網站首頁。The page is embedded into the binary; the same `index.html` is also the GitHub Pages homepage.
- **連接埠 · Port**：固定 `127.0.0.1:8787`。固定埠讓瀏覽器 `localStorage` 的來源（origin）維持不變，資料才會跨啟動保留。A fixed port keeps the browser origin constant so localStorage persists across launches.
- **日期運算 · Dates**：一律以 UTC 整日計算，避免時區／日光節約偏移；12 個月窗口以日曆月推算，正確處理閏年。All dates are whole UTC days; 12-month windows use calendar months and handle leap years correctly.

---

## 免責聲明 · Disclaimer

**本工具僅供參考，並非法律或移民意見。** 英國內政部指引會不時更新，個案可能有離境豁免（如 COVID、嚴重疾病、受僱等），「旅行當日」計法與 5 年起算方式亦存在不同詮釋。實際資格請以 **GOV.UK** 官方指引（「BN(O) visa」「Indefinite leave to remain」）及你的個案為準，並於需要時諮詢合資格的移民顧問。

**This tool is for reference only and is not legal or immigration advice.** Home Office guidance changes over time, individual cases may qualify for absence exceptions (e.g. COVID, serious illness, employment), and there are differing interpretations of travel-day counting and the 5-year start date. Rely on official **GOV.UK** guidance ("BN(O) visa"; "Indefinite leave to remain") and your own circumstances, and consult a qualified immigration adviser where needed.
