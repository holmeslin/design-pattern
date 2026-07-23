# Pipeline(管線)

Pipeline 是一種將資料處理拆成多個階段、以 channel 串接起來的並發模式,每個階段各自在獨立的 goroutine 中運作。

## 說明

本範例以三個階段組成一條管線:

1. `gen`:來源階段,將傳入的整數逐一送進 channel,送完後關閉 channel。
2. `square`:中間階段,從上一階段的 channel 讀入整數並計算平方,送到自己的輸出 channel;讀取完畢(上游 channel 關閉)後,關閉自己的輸出 channel。
3. `sum`:匯集階段,持續讀取前一階段的 channel 直到其關閉,加總所有數值後回傳結果。

資料流向為:`gen -> square -> sum`。每個階段只負責關閉「自己產生」的輸出 channel,關閉狀態會沿著管線依序往下游傳遞,下游只需以 `range` 讀取 channel,不需額外的同步機制,即可知道資料何時讀取完畢。

## 執行

```sh
go test -race ./...
```
