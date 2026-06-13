# codex-legacy — 技術脈絡 (Technical Context)

## 專案結構 (Project Structure)

```tree
codex-legacy/
├── README.md                    (對外說明 — 業務領域 / 使用方式)
├── CLAUDE.md                    (本檔 — 技術脈絡)
├── README.todo                  (個人規劃草稿, 26 KB 待整合)
├── LICENSE
├── .gitignore                   (Node/Go/Python/Java/C++/IDE/secret 規則)
│
├── concepts/                    CS 概念索引 (knowledge-only)
│   ├── basic_cs.md
│   ├── testing.md
│   ├── cryptographic/
│   ├── database/
│   ├── design_pattern/          Abstract_Factory / Builder / Factory_method / Singleton / Prototype / Object_pool / OO / software_methodology
│   ├── http/
│   ├── interview/               booking.com / fb / google / line / job.md
│   ├── math/
│   ├── ml/
│   ├── other/
│   ├── security/
│   ├── server/                  apache / gogs / hadoop / linux / mysql / nginx / php5 / R_pi
│   └── system_design/           cache / recommend_system / TCPIP_note
│
├── notes/                       個人工作筆記 + 兩個 Go 模組根
│   ├── go.mod                   (module github.com/bizshuk/techstack, Go 1.19)
│   ├── go.sum
│   ├── main.go                  (練習 entry: IntHeap 範例)
│   ├── .gitignore.*
│   ├── default.code-workspace
│   ├── LICENSE.note-fortech
│   ├── README.note-fortech.md
│   ├── TODO.md / TODO_*.md / TODO_auto.md.obsolete
│   ├── programming_language.todo / system_design.todo / tech_stack.todo
│   ├── todo.sh                  (todo 維護腳本)
│   ├── domain/                  CHC / Leadership / Management / Regulation / SoftwareEngineer / sport / talk
│   ├── platform/                ElasticSearch / Git / golang / java / kubernetes / linux / Node / VScode / VScode-note-fortech
│   └── topic/                   Backend / cs_concept / Data / DevOps / FrontEnd / interview / monitoring / project / protocol / security / storage / system_design
│
├── algo/                        演算法 / 資料結構練習
│   ├── go/                      (獨立 Go module, 詳見 go.mod)
│   │   ├── go.mod
│   │   ├── main.go              (entry — 切換呼叫不同子 package)
│   │   ├── README.md
│   │   ├── stack.md
│   │   ├── algo/  basic-algo/  basic-ds/  interview/  playground/  questions/  topic/  util/
│   │   ├── algo.todo  terminology.todo
│   │   └── TimeComplexityAnalysis.png
│   └── multi-lang/              跨語言 LeetCode (algo/  leecode/)
│
├── frameworks/                  框架樣板 / 設計模式展示
│   ├── eventbus-go/             (Go module: Dispatcher / Event / EventService / SNSOption / SNSService / ExampleService + test/)
│   ├── express-boilerplate/     (Node: app.js / bin / common.js / lib / log / public / routes / setting.js / src / views / webpack.config.js + package.json)
│   ├── pizza-patterns/          (Node: index.js / lib + package.json — Factory / Builder 展示)
│   └── rundeck-tomcat/          (jobs.xml + LICENSE — Rundeck 部署設定)
│
└── playground/                  練習與沙盒
    ├── fullstack-node/          (public/ + src/ + package.json + package-lock.json)
    ├── python/                  (basic.md / decorator.py / fib.py / function.py / itertools.py / log_analysis/ / oo.py / queue.py / reserved_seat.py / test_*.py ...)
    ├── go/                      (bit / go_test / goroutine / language_spec.md / list / math / slice / string / syntax.md / test / tools.md / tree + DEBUG.md / README.md)
    ├── slURL/                   (Node: app.js / bin / public / routes / views + package.json + webpack.config.js + LICENSE — Express + React 15 URL 短網址服務, 2016 歷史快照, lib/ 與 setting.js 未 commit)
    ├── c_cpp/  cache/  cryptographic/  css/  design_pattern/  framework/
    ├── FreeCodeCamp.md  makefile.md  Object-Oriented.md  README.md  README2.md  testing.md  html_header_sample.md  session.md
    ├── fullstack-node/  http/  internet/  java/  js/  language/  lua/
    └── markdown/  math/  ml/  other/  perl/  php/  pro_lang/  PWA/  security/  server/  system_design/  test/  xml/  yaml/
```

> 目錄名單引自實際 `ls -la`; 完整子目錄以 `find` 結果為準, 此處只列高訊號節點。

## 技術棧 (Tech Stack)

- Language: 多語言 — Go (1.19)、JavaScript (Node.js)、Python、Java、C/C++、Lua、PHP、Perl、HTML/CSS/MD
- Framework:
    - Go: 自製 `eventbus-go` 函式庫 (含 SNS / DLQ 範例)
    - Node: `express-boilerplate` (Express + webpack 起手)
    - Node: `playground/slURL` (Express + React 15 + Jade URL 短網址服務, 2016 歷史快照)
    - JS pattern: `pizza-patterns` (Factory + Builder)
- Build tool: 各子模組獨立 — `go build` / `npm` / `make` (部分 `playground/c_cpp`); root **無** 統一 build script
- Key dependencies:
    - `algo/go/go.mod` — Go 1.19, 演算法練習自製
    - `notes/go.mod` — `github.com/sirupsen/logrus v1.9.0`, `github.com/stretchr/testify v1.8.0`
    - `frameworks/eventbus-go/go.mod` — eventbus 套件
    - `frameworks/express-boilerplate/package.json` — Express + webpack
    - `playground/fullstack-node/package.json` — Node fullstack 練習
    - `playground/slURL/package.json` — `express ^4.14.0`, `react ^15.4.1`, `webpack ^1.13.3`, `mysql ^2.12.0`, `crypto-random-string ^1.0.0`, `jade ^1.11.0` (2016 歷史快照)
- 測試框架: Go 用 `testing` + `testify`; JS 用框架內建; Python 部分檔案用 `test_*.py` 命名
- Lint: 未偵測到 root `.golangci.yml` / `.eslintrc` / `setup.cfg`; 已有 `markdownlint` skill 可在生成 markdown 時使用

## 關鍵決策 (Key Decisions)

- 合併策略 A (純檔案搬移): 來自 12 個獨立 repo, 已於 2026-06-13 合併, **獨立 git 歷史已移除**; 沒有 submodule / 沒有 monorepo tool (Nx / Turborepo)
- 多語言 monorepo 但各語言子模組獨立: Go 模組各自有 `go.mod` (`algo/go/`, `notes/`, `frameworks/eventbus-go/`); JS 各自有 `package.json` (`frameworks/express-boilerplate/`, `frameworks/pizza-patterns/`, `playground/fullstack-node/`, `playground/slURL/`)
- `playground/slURL/` 為 2016 歷史快照: 套件 (React 15 / Webpack 1 / Babel 6) 保留原貌, `.gitignore` 刻意排除 `lib/` 與 `setting.js` (DB 連線 / short-domain 設定) — 不可直接 `npm start`,需自行補上設定檔
- 知識 vs 程式分離: 純 markdown 概念放 `concepts/`, 個人補充筆記放 `notes/`, 可執行程式放 `algo/ / frameworks/ / playground/`
- 沒有統一 build / test: 各子模組獨立 `go test ./...` 或 `npm test`, root 不提供 aggregate
- `.gitignore` 已聚合多語言: Node (`node_modules/`, `package-lock.json.bak`) + Go (`vendor/`, `*.exe`) + Python (`__pycache__/`, `venv/`) + Java (`target/`, `build/`) + C/C++ (`build/`, `out/`) + 通用 OS/IDE 規則

## 模組對應 (Module Mapping)

| 業務領域 (Domain)                  | 套件/模組 (Package/Module)                                                                                               | 進入點 (Entry Point)                                                                                                                                        |
| ---------------------------------- | ------------------------------------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 知識概念庫 (Knowledge Concepts)    | `concepts/<topic>/`                                                                                                      | `concepts/<topic>/README.md` (e.g. `concepts/design_pattern/Factory_method.md`)                                                                             |
| 個人技術筆記 (Personal Tech Notes) | `notes/{domain,platform,topic}/` + `notes/main.go`                                                                       | `notes/topic/Backend/`, `notes/platform/ElasticSearch/`, `notes/main.go`                                                                                    |
| 演算法與資料結構 (Algorithms & DS) | `algo/go/{algo,basic-algo,basic-ds,interview,playground,questions,topic,util}/`                                          | `algo/go/main.go`, `algo/multi-lang/algo/`                                                                                                                  |
| 框架樣板 (Framework Templates)     | `frameworks/eventbus-go/`, `frameworks/express-boilerplate/`, `frameworks/pizza-patterns/`, `frameworks/rundeck-tomcat/` | `frameworks/eventbus-go/Dispatcher.go`, `frameworks/express-boilerplate/app.js`, `frameworks/pizza-patterns/index.js`, `frameworks/rundeck-tomcat/jobs.xml` |
| 練習與沙盒 (Playground)            | `playground/{fullstack-node,python,go,js,c_cpp,slURL,...}/`                                                              | `playground/fullstack-node/package.json`, `playground/python/decorator.py`, `playground/go/goroutine/`, `playground/slURL/app.js`                          |

## 開發指南 (Development Guide)

### 前置需求 (Prerequisites)

- Go 1.19+ (給 `algo/go/`, `notes/`, `frameworks/eventbus-go/`)
- Node.js + npm (給 `frameworks/express-boilerplate/`, `frameworks/pizza-patterns/`, `playground/fullstack-node/`, `playground/slURL/`)
- Python 3 (給 `playground/python/`)
- MySQL (給 `playground/slURL/`, 運行時需自行建立 `url` schema)
- 其他語言 (Java / C/C++ / Lua / PHP / Perl) 視需要

### 安裝 (Installation)

> 本 repo 不需要 root 安裝; 各子模組獨立安裝依賴

```bash
# Go 子模組
cd algo/go && go mod download
cd notes && go mod download
cd frameworks/eventbus-go && go mod download

# JS / Node 子模組
cd frameworks/express-boilerplate && npm install
cd frameworks/pizza-patterns && npm install
cd playground/fullstack-node && npm install
cd playground/slURL && npm install        # 歷史快照, 需自行補 lib/database.js + setting.js 才能啟動
```

### 建置 (Build)

```bash
# Go 範例 / 套件
cd algo/go && go build ./...        # 注意 main.go 是練習切換, 多數子 package 為函式庫
cd frameworks/eventbus-go && go build ./...

# Node 樣板 (部分含 webpack)
cd frameworks/express-boilerplate && npm run build    # 若 package.json 有此 script
cd playground/slURL && npm run build                 # webpack -p (2016 版, 若升級可能壞)
```

> Root 無 `Makefile`; 沒有統一的 build 指令

### 測試 (Test)

```bash
# Go
cd algo/go && go test ./...
cd notes && go test ./...
cd frameworks/eventbus-go && go test ./...

# JS
cd frameworks/express-boilerplate && npm test
cd frameworks/pizza-patterns && npm test
# 註: playground/slURL 原始 package.json 沒有 test script (echo "Error: no test specified"),無單元測試

# Python
cd playground/python && python3 -m unittest discover -s . -p 'test_*.py'
```

### 部署 (Deploy)

未偵測到部署設定 (No deployment config detected)

- 唯一可能與部署相關: `frameworks/rundeck-tomcat/jobs.xml` 是 Rundeck job 設定, 用於把 Tomcat 部署工作交給 Rundeck 排程; 但 `frameworks/rundeck-tomcat/` 內無 `rundeck` CLI / hook 整合, 屬於設定片段
- 無 CI (`./.github/` 未發現), 無 Dockerfile, 無 helm chart

## 慣例 (Conventions)

- 命名:
    - Go package: 小寫單字 (`algo`, `basicds`, `util`)
    - Go test: `<file>_test.go` 與 source 並列 (`frameworks/eventbus-go/Dispatcher_test.go`)
    - JS 樣板: `app.js` / `common.js` / `setting.js` (camelCase 檔名, `routes/` 子目錄分層)
    - Markdown: 全小寫 + 底線 (`design_pattern/Abstract_Factory.md`, `testing.md`)
- 錯誤處理: Go 慣用 `if err != nil { ... }` + 回傳; 詳細 pattern 未集中規範, 散見各子模組
- Logging: Go 用 `github.com/sirupsen/logrus` (`notes/go.mod`); JS 端有 `frameworks/express-boilerplate/log/` 子目錄
- Testing: Go 用 `testing` + `testify/assert`; Python 用 `test_*.py` + `unittest` 風格; JS 由各樣板自帶
- 結構: 「概念 / 筆記 / 程式 / 樣板 / 沙盒」五類分離, 不互相 import; 子模組之間以 `cd <dir>` 切換
- 個人 TODO: `notes/TODO*.md` + `notes/*.todo` + `notes/todo.sh` 維護; 與正式文件 (`README.md` / `CLAUDE.md`) 分離
