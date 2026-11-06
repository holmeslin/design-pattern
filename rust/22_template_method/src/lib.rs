// 三個步驟沒有預設實作,漏寫就編譯不過;run 是模板,固定步驟順序。
pub trait DataJob {
    fn read(&self) -> Vec<i32>;
    fn process(&self, data: Vec<i32>) -> Vec<i32>;
    fn write(&self, data: Vec<i32>) -> String;

    fn run(&self) -> String {
        self.write(self.process(self.read()))
    }
}

pub struct EvenFilterJob;

impl DataJob for EvenFilterJob {
    fn read(&self) -> Vec<i32> {
        vec![1, 2, 3, 4]
    }

    fn process(&self, data: Vec<i32>) -> Vec<i32> {
        data.into_iter().filter(|n| n % 2 == 0).collect()
    }

    fn write(&self, data: Vec<i32>) -> String {
        let items: Vec<String> = data.iter().map(i32::to_string).collect();
        format!("[{}]", items.join(" "))
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // run 固定「讀 -> 處理 -> 寫」的順序,子類只填三個步驟
    #[test]
    fn run_follows_read_process_write() {
        assert_eq!(EvenFilterJob.run(), "[2 4]");
    }
}
