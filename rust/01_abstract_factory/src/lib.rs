// trait 明確宣告「一族產品」的形狀;呼叫端只認 trait,不知道具體是哪一族。
pub trait Button {
    fn render(&self) -> String;
}

pub trait Checkbox {
    fn render(&self) -> String;
}

pub trait GuiFactory {
    fn create_button(&self) -> Box<dyn Button>;
    fn create_checkbox(&self) -> Box<dyn Checkbox>;
}

struct DarkButton;
struct DarkCheckbox;
struct LightButton;
struct LightCheckbox;

impl Button for DarkButton {
    fn render(&self) -> String {
        "[dark button]".to_string()
    }
}

impl Checkbox for DarkCheckbox {
    fn render(&self) -> String {
        "[dark checkbox]".to_string()
    }
}

impl Button for LightButton {
    fn render(&self) -> String {
        "[light button]".to_string()
    }
}

impl Checkbox for LightCheckbox {
    fn render(&self) -> String {
        "[light checkbox]".to_string()
    }
}

pub struct DarkFactory;
pub struct LightFactory;

impl GuiFactory for DarkFactory {
    fn create_button(&self) -> Box<dyn Button> {
        Box::new(DarkButton)
    }
    fn create_checkbox(&self) -> Box<dyn Checkbox> {
        Box::new(DarkCheckbox)
    }
}

impl GuiFactory for LightFactory {
    fn create_button(&self) -> Box<dyn Button> {
        Box::new(LightButton)
    }
    fn create_checkbox(&self) -> Box<dyn Checkbox> {
        Box::new(LightCheckbox)
    }
}

// 用 &dyn 做執行期選擇:主題通常是讀設定決定的,編譯期不知道。
pub fn render_ui(factory: &dyn GuiFactory) -> String {
    format!(
        "{} {}",
        factory.create_button().render(),
        factory.create_checkbox().render()
    )
}

#[cfg(test)]
mod tests {
    use super::*;

    // 換工廠就換整族產品,render_ui 不用改
    #[test]
    fn factory_produces_matching_family() {
        assert_eq!(render_ui(&DarkFactory), "[dark button] [dark checkbox]");
        assert_eq!(render_ui(&LightFactory), "[light button] [light checkbox]");
    }
}
