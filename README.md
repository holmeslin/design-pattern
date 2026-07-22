# design-pattern

用不同程式語言實作經典設計模式(Design Patterns)的練習集。

## 結構

每種語言各自一個資料夾,資料夾內再依模式分類:

```
design-pattern/
├── java/
├── python/
├── go/
└── ...
```

每個模式建議放一個獨立子資料夾,附上簡短說明與可執行範例。例如:

```
python/
└── singleton/
    ├── README.md      # 這個模式在做什麼、何時用
    └── singleton.py   # 範例實作
```

## 模式分類(GoF)

- **建立型(Creational)**:Singleton、Factory Method、Abstract Factory、Builder、Prototype
- **結構型(Structural)**:Adapter、Bridge、Composite、Decorator、Facade、Flyweight、Proxy
- **行為型(Behavioral)**:Chain of Responsibility、Command、Iterator、Mediator、Memento、Observer、State、Strategy、Template Method、Visitor

## 授權

[Apache License 2.0](LICENSE)
