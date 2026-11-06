// 小寫工廠函式:避免和 Node 全域的 File 類別撞名。
// Leaf 與 Composite 都有 size(),呼叫端不必分辨是檔案還是資料夾。

export const file = (size) => ({ size: () => size });

export function folder() {
  const children = [];
  return {
    add: (node) => {
      children.push(node);
    },
    size: () => children.reduce((sum, node) => sum + node.size(), 0),
  };
}
