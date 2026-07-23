# Fan-in / Fan-out

Fan-out / Fan-in 是一種並發模式:先把工作分散給多個 goroutine 平行處理(Fan-out),再把多個 goroutine 的結果收斂回單一 channel(Fan-in)。

## 流程說明

1. **來源(gen)**:產生一條 channel,依序送出所有待處理的數值,送完後關閉。
2. **Fan-out(一分多)**:啟動多個 `square` goroutine,讓它們同時從同一條輸入 channel 讀取資料並各自平方,達到平行處理的效果。
3. **Fan-in(多合一)**:`merge` 為每條輸入 channel 各開一個轉送 goroutine,並用 `sync.WaitGroup` 追蹤全部轉送工作是否完成,待全部完成後才關閉合併後的輸出 channel,把多條結果收斂成單一來源。
4. **彙整**:對合併後的 channel 逐一取值加總。由於加法滿足交換律,即使 fan-out 的處理順序每次執行都可能不同,最終加總結果仍然穩定。

## 執行

```sh
go test -race ./...
```
