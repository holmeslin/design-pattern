# design-pattern

用不同程式語言實作經典設計模式(Design Patterns)的練習集。

## 結構

每種語言各自一個資料夾,資料夾內再依模式分類:

```text
design-pattern/
├── java/
├── python/
├── go/
└── ...
```

每個模式放一個獨立子資料夾,附上簡短說明與可執行範例。例如:

```text
python/
└── singleton/
    ├── README.md      # 這個模式在做什麼、何時用
    └── singleton.py   # 範例實作
```

## 已實作

- **[go/](go/)** — GoF 23 種 + 4 種 Go 慣用並發模式,依 GoF 書本順序編號,每個模式一個獨立 module 與可跑的 `example_test.go`。詳見 [go/README.md](go/README.md)。

## 模式分類(GoF)

- **建立型(Creational)**:Singleton、Factory Method、Abstract Factory、Builder、Prototype
- **結構型(Structural)**:Adapter、Bridge、Composite、Decorator、Facade、Flyweight、Proxy
- **行為型(Behavioral)**:Chain of Responsibility、Command、Interpreter、Iterator、Mediator、Memento、Observer、State、Strategy、Template Method、Visitor

## 授權

[Apache License 2.0](LICENSE)
