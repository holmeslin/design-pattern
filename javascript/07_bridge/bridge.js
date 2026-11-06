// Implementor:裝置只管 name / volume / setVolume(duck typing),遙控器不認識具體類別。
export class TV {
  name = 'TV';
  volume = 0;

  setVolume(v) {
    this.volume = v;
  }
}

export class Radio {
  name = 'Radio';
  volume = 0;

  setVolume(v) {
    this.volume = v;
  }
}

// Abstraction:遙控器透過「橋」(device 欄位)操作裝置,兩條階層各自獨立擴充。
export class RemoteControl {
  #device;

  constructor(device) {
    this.#device = device;
  }

  volumeUp() {
    this.#device.setVolume(this.#device.volume + 10);
  }

  status() {
    return `${this.#device.name} volume: ${this.#device.volume}`;
  }
}
