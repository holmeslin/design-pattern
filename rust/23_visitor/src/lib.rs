pub struct Circle {
    pub radius: f64,
}

pub struct Rectangle {
    pub width: f64,
    pub height: f64,
}

// 每種形狀一個 visit_*:新增操作只要多一個 Visitor 實作,不用改形狀。
pub trait Visitor {
    fn visit_circle(&self, c: &Circle) -> f64;
    fn visit_rectangle(&self, r: &Rectangle) -> f64;
}

pub trait Shape {
    fn accept(&self, v: &dyn Visitor) -> f64;
}

// accept 把自己的具體型別交給 visitor(雙重分派)。
impl Shape for Circle {
    fn accept(&self, v: &dyn Visitor) -> f64 {
        v.visit_circle(self)
    }
}

impl Shape for Rectangle {
    fn accept(&self, v: &dyn Visitor) -> f64 {
        v.visit_rectangle(self)
    }
}

pub struct AreaVisitor;

impl Visitor for AreaVisitor {
    fn visit_circle(&self, c: &Circle) -> f64 {
        std::f64::consts::PI * c.radius * c.radius
    }

    fn visit_rectangle(&self, r: &Rectangle) -> f64 {
        r.width * r.height
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // 同一個 visitor 對不同形狀各自分派到對應的 visit_*
    #[test]
    fn area_visitor_computes_area_per_shape() {
        assert_eq!(
            Rectangle {
                width: 2.0,
                height: 3.0
            }
            .accept(&AreaVisitor),
            6.0
        );
        assert_eq!(
            Circle { radius: 1.0 }.accept(&AreaVisitor),
            std::f64::consts::PI
        );
    }
}
