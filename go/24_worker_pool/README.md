# Worker Pool(工作池)

以固定數量的 goroutine(worker)共同消費同一個 job channel 來處理一批工作,藉此限制並發數量、避免為每筆工作都建立一個 goroutine。

## 角色

- **Job channel**:所有工作的來源,由一個 goroutine 負責派送並在派送完畢後關閉。
- **Worker**:固定數量的 goroutine,各自從同一個 job channel 取值運算(平方),再把結果送進 result channel。因為 channel 保證每筆工作只會被其中一個 worker 取走,所以可用少量 goroutine 消化任意數量的工作。
- **Result channel**:彙整所有 worker 的運算結果,交由主 goroutine 讀取加總。
- **sync.WaitGroup**:追蹤所有 worker 是否已收工,待全部完成後才 close result channel,讓讀取端能正常結束迴圈。

由於加總(整數相加)具有交換律,結果與 worker 的執行順序或完成順序無關,因此輸出穩定、可重現。

## 執行

```sh
go test -race ./...
```
