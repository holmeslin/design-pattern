// 子系統不匯出:對外只有 HomeTheater 這個門面。
const lights = { dim: () => 'lights dim' };
const projector = { on: () => 'projector on' };
const amplifier = { on: () => 'amplifier on' };

export class HomeTheater {
  watchMovie() {
    return [lights.dim(), projector.on(), amplifier.on()];
  }
}
