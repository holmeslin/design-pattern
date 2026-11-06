// accept 回呼對應型別的 visit 方法(double dispatch),這樣新增操作不用改圖形 class。
export class Circle {
  constructor(radius) {
    this.radius = radius;
  }

  accept(visitor) {
    return visitor.visitCircle(this);
  }
}

export class Rectangle {
  constructor(width, height) {
    this.width = width;
    this.height = height;
  }

  accept(visitor) {
    return visitor.visitRectangle(this);
  }
}

export const areaVisitor = {
  visitCircle: (c) => Math.PI * c.radius ** 2,
  visitRectangle: (r) => r.width * r.height,
};
