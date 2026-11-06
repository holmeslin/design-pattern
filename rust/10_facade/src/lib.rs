// 子系統不對外公開:呼叫端只能透過 HomeTheater 這個門面操作。
#[derive(Default)]
struct Lights;
#[derive(Default)]
struct Projector;
#[derive(Default)]
struct Amplifier;

impl Lights {
    fn dim(&self) -> &'static str {
        "lights dim"
    }
}

impl Projector {
    fn on(&self) -> &'static str {
        "projector on"
    }
}

impl Amplifier {
    fn on(&self) -> &'static str {
        "amplifier on"
    }
}

#[derive(Default)]
pub struct HomeTheater {
    lights: Lights,
    projector: Projector,
    amplifier: Amplifier,
}

impl HomeTheater {
    pub fn watch_movie(&self) -> Vec<&'static str> {
        vec![self.lights.dim(), self.projector.on(), self.amplifier.on()]
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // 一個呼叫就依序驅動三個子系統,呼叫端不必認識它們
    #[test]
    fn watch_movie_runs_subsystems_in_order() {
        assert_eq!(
            HomeTheater::default().watch_movie(),
            vec!["lights dim", "projector on", "amplifier on"]
        );
    }
}
