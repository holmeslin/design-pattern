// 實作端:各種裝置共同的低階介面。
pub trait Device {
    fn name(&self) -> &str;
    fn volume(&self) -> u8;
    fn set_volume(&mut self, volume: u8);
}

#[derive(Default)]
pub struct Tv {
    volume: u8,
}

#[derive(Default)]
pub struct Radio {
    volume: u8,
}

impl Device for Tv {
    fn name(&self) -> &str {
        "TV"
    }
    fn volume(&self) -> u8 {
        self.volume
    }
    fn set_volume(&mut self, volume: u8) {
        self.volume = volume;
    }
}

impl Device for Radio {
    fn name(&self) -> &str {
        "Radio"
    }
    fn volume(&self) -> u8 {
        self.volume
    }
    fn set_volume(&mut self, volume: u8) {
        self.volume = volume;
    }
}

// 抽象端:遙控器以泛型持有裝置(編譯期綁定),兩條繼承軸各自擴充。
pub struct RemoteControl<D: Device> {
    device: D,
}

impl<D: Device> RemoteControl<D> {
    pub fn new(device: D) -> Self {
        Self { device }
    }

    pub fn volume_up(&mut self) {
        // saturating_add:音量到頂就停在 u8::MAX,不溢位。
        self.device
            .set_volume(self.device.volume().saturating_add(10));
    }

    pub fn status(&self) -> String {
        format!("{} volume: {}", self.device.name(), self.device.volume())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // 遙控器(抽象)與裝置(實作)各自獨立,任意組合
    #[test]
    fn remote_works_with_any_device() {
        let mut tv = RemoteControl::new(Tv::default());
        tv.volume_up();
        assert_eq!(tv.status(), "TV volume: 10");

        let mut radio = RemoteControl::new(Radio::default());
        radio.volume_up();
        assert_eq!(radio.status(), "Radio volume: 10");
    }
}
