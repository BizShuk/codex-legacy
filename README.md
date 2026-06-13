# codex-legacy

> 個人程式碼與知識彙整 monorepo — 從 [BizShuk](https://github.com/BizShuk) 旗下 11 個 GitHub repo 合併而成

## 結構總覽

| 分類 | 用途 | 收錄 repo |
|------|------|-----------|
| `concepts/` | CS 概念索引 (design pattern, system design...) | code_concept |
| `notes/` | 個人技術筆記 (Backend, Frontend, DevOps...) | NoteForTech, techstack |
| `algos/` | 演算法/資料結構刷題 | algo, code_algo |
| `frameworks/` | 框架樣板、設計模式、整合型 code pattern (含 rundeck) | eventbus, PizzaFactoryAndBuilder, defaultExpress, GenericRundeckTomcatDeployment |
| `playground/` | 練習專案、多語言沙盒 | fullstack-practice, code_sandbox |

## 目錄樹

```
codex-legacy/
├── README.md                    (本檔)
├── LICENSE
├── .gitignore
│
├── concepts/
│   └── code-concept/             ← code_concept        (CS 概念庫)
│
├── notes/
│   ├── tech-notes/               ← NoteForTech         (技術主題筆記)
│   └── techstack-notes/          ← techstack           (技術棧彙整)
│
├── algos/
│   ├── go-algo/                  ← algo                (Go 演算法/DS)
│   └── multi-lang-algo/          ← code_algo           (多語 LeetCode)
│
├── frameworks/
│   ├── eventbus-go/              ← eventbus            (Go eventbus 函式庫)
│   ├── pizza-patterns/           ← PizzaFactoryAndBuilder (Factory/Builder 範例)
│   ├── express-boilerplate/      ← defaultExpress      (Express 起手式)
│   └── rundeck-tomcat/           ← GenericRundeckTomcatDeployment (Rundeck 部署)
│
└── playground/
    ├── fullstack-node/           ← fullstack-practice  (Node fullstack 練習)
    └── multi-lang-sandbox/       ← code_sandbox        (多語言沙盒)
```

## 各分類快速索引

### concepts/

CS 概念筆記 — `design_pattern` / `system_design` / `http` / `cryptographic` / `database` / `ml` / `security` / `math` / `interview` / `server`

### notes/

個人工作筆記 — `Backend` / `Frontend` / `DevOps` / `CHC` / `ElasticSearch` / `Git` / `Leadership` / `SoftwareEngineer` / `protocol` / `interview`

### algos/

演算法/資料結構練習 — 以 Go 為主，輔以 Python/JS

### frameworks/

- `eventbus-go` — Go eventbus 函式庫 (含 SNS/DLQ 範例)
- `pizza-patterns` — Factory/Builder 設計模式
- `express-boilerplate` — Express.js + webpack 起手式
- `rundeck-tomcat` — Rundeck 部署 Tomcat 設定 (jobs.xml)

### playground/

- `fullstack-node` — Node.js fullstack 練習 (含 public/src)
- `multi-lang-sandbox` — 多語言沙盒 (JS/MD/C++/Go/Python)

## 合併資訊

- **合併日期**: 2026-06-13
- **合併策略**: 純檔案搬移 (Strategy A) — 各 repo 的獨立 git 歷史已移除
- **來源**: https://github.com/BizShuk
