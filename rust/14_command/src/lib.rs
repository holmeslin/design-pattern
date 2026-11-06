#[derive(Default)]
pub struct Account {
    pub balance: i64,
}

// command 不持有 &mut Account,而是在 execute/undo 時才傳入;
// 否則 History 裡的多個 command 會同時可變借用同一個帳戶,借用檢查器不允許。
pub trait Command {
    fn execute(&self, account: &mut Account);
    fn undo(&self, account: &mut Account);
}

pub struct Deposit {
    pub amount: i64,
}

// wrapping 讓 undo 永遠是 execute 的精確反運算,而且溢位時不會 panic。
impl Command for Deposit {
    fn execute(&self, account: &mut Account) {
        account.balance = account.balance.wrapping_add(self.amount);
    }

    fn undo(&self, account: &mut Account) {
        account.balance = account.balance.wrapping_sub(self.amount);
    }
}

#[derive(Default)]
pub struct History {
    done: Vec<Box<dyn Command>>,
}

impl History {
    pub fn run(&mut self, cmd: Box<dyn Command>, account: &mut Account) {
        cmd.execute(account);
        self.done.push(cmd);
    }

    pub fn undo_last(&mut self, account: &mut Account) {
        if let Some(cmd) = self.done.pop() {
            cmd.undo(account);
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // 執行後餘額增加,undo_last 還原成 0
    #[test]
    fn run_then_undo_restores_balance() {
        let mut account = Account::default();
        let mut history = History::default();
        history.run(Box::new(Deposit { amount: 100 }), &mut account);
        assert_eq!(account.balance, 100);
        history.undo_last(&mut account);
        assert_eq!(account.balance, 0);
    }

    // 歷史已空時 undo_last 什麼都不做
    #[test]
    fn undo_on_empty_history_is_noop() {
        let mut account = Account::default();
        let mut history = History::default();
        history.run(Box::new(Deposit { amount: 100 }), &mut account);
        history.undo_last(&mut account);
        history.undo_last(&mut account);
        assert_eq!(account.balance, 0);
    }
}
