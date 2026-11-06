// 快照的欄位私有:使用這個 crate 的程式碼只能把它交還給 Editor,看不到也改不了內容。
pub struct Memento {
    content: String,
}

#[derive(Default)]
pub struct Editor {
    content: String,
}

impl Editor {
    // 取代整份內容;`type` 是關鍵字,所以方法叫 type_text。
    pub fn type_text(&mut self, text: &str) {
        self.content = text.to_string();
    }

    pub fn content(&self) -> &str {
        &self.content
    }

    pub fn save(&self) -> Memento {
        Memento {
            content: self.content.clone(),
        }
    }

    // 取得所有權:快照用掉就沒了,同一份不能還原兩次。
    pub fn restore(&mut self, memento: Memento) {
        self.content = memento.content;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // 存檔後繼續編輯,還原就回到存檔當時的內容
    #[test]
    fn restore_returns_to_saved_content() {
        let mut editor = Editor::default();
        editor.type_text("v1");
        let saved = editor.save();
        editor.type_text("v2");
        assert_eq!(editor.content(), "v2");
        editor.restore(saved);
        assert_eq!(editor.content(), "v1");
    }
}
