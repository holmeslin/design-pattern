//! `Send` / `Sync` 讓資料競爭變成編譯錯誤:不能跨 thread 傳遞的型別,根本無法交給 `thread::spawn`。
//!
//! 合法:`Arc` 是 `Send`,可以搬進 `thread::spawn`。
//!
//! ```
//! let data = std::sync::Arc::new(5);
//! let n = std::thread::spawn(move || *data + 1).join().unwrap();
//! assert_eq!(n, 6);
//! ```
//!
//! 非法:`Rc` 的計數不是原子操作,所以不是 `Send`。
//! 與上面的範例只差 `Arc` 換成 `Rc`,證明失敗來自這一處。
//!
//! ```compile_fail
//! let data = std::rc::Rc::new(5);
//! let n = std::thread::spawn(move || *data + 1).join().unwrap();
//! assert_eq!(n, 6);
//! ```

use std::sync::{Arc, Mutex, mpsc};
use std::thread;

/// `workers` 必須大於 0;傳 0 是呼叫端的錯誤用法,所以用 `assert!` 而不是回傳 `Result`。
pub fn sum_squares_parallel(nums: &[i64], workers: usize) -> i64 {
    assert!(workers > 0, "workers must be > 0");
    // max(1):空資料時 len 為 0,chunks(0) 會 panic。
    let chunk_size = nums.len().div_ceil(workers).max(1);
    let (tx, rx) = mpsc::channel();

    // scope 保證所有 thread 在回傳前結束,所以 thread 可以直接借用 nums,不需要 Arc。
    thread::scope(|s| {
        for chunk in nums.chunks(chunk_size) {
            let tx = tx.clone();
            s.spawn(move || {
                let partial: i64 = chunk.iter().map(|x| x * x).sum();
                // send 只會在 receiver 被 drop 時失敗,而 rx 活到 scope 結束之後。
                tx.send(partial).expect("receiver alive until scope ends");
            });
        }
        // 每個 thread 拿的是 clone;原本的 tx 不 drop,channel 永遠不會關閉,下面的收集會卡住。
        drop(tx);
        rx.iter().sum()
    })
}

pub fn count_with_threads(threads: usize, increments: u64) -> u64 {
    // thread::spawn 要求 'static,不能借用堆疊資料,所以共用的狀態要放進 Arc。
    let counter = Arc::new(Mutex::new(0u64));
    let handles: Vec<_> = (0..threads)
        .map(|_| {
            let counter = Arc::clone(&counter);
            thread::spawn(move || {
                for _ in 0..increments {
                    // 中毒時取回內部資料繼續用:加 1 不會讓計數處於不一致狀態。
                    *counter.lock().unwrap_or_else(|e| e.into_inner()) += 1;
                }
            })
        })
        .collect();
    for handle in handles {
        // worker panic 代表程式有 bug,沒有可恢復的狀態,所以直接 expect。
        handle.join().expect("worker thread panicked");
    }
    *counter.lock().unwrap_or_else(|e| e.into_inner())
}

#[cfg(test)]
mod tests {
    use super::*;

    // 切成 3 塊平行計算,結果與單執行緒相同
    #[test]
    fn parallel_sum_of_squares() {
        assert_eq!(sum_squares_parallel(&[1, 2, 3, 4, 5], 3), 55);
    }

    // workers 比資料筆數多、或資料為空,都不能出錯
    #[test]
    fn more_workers_than_items_and_empty_input() {
        assert_eq!(sum_squares_parallel(&[1, 2], 8), 5);
        assert_eq!(sum_squares_parallel(&[], 3), 0);
    }

    // workers 為 0 屬於呼叫端的錯誤用法,直接 assert
    #[test]
    #[should_panic(expected = "workers must be > 0")]
    fn zero_workers_panics() {
        sum_squares_parallel(&[1], 0);
    }

    // Arc<Mutex> 讓多個 thread 安全地共用同一個計數器
    #[test]
    fn counter_with_threads() {
        assert_eq!(count_with_threads(4, 1000), 4000);
    }
}
