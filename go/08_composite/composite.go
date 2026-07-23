// Package composite 示範組合模式(Composite Pattern)。
//
// 組合模式將物件組合成樹狀結構,以表示「部分-整體」的階層關係,
// 讓使用者可以用一致的方式處理單一物件(Leaf)與物件的組合(Composite)。
package composite

// Node 是組合模式中的 Component 角色,定義所有節點(不論是葉節點還是
// 容器節點)共同支援的操作介面。
type Node interface {
	// Size 回傳該節點所佔的大小。
	Size() int
}

// File 是組合模式中的 Leaf 角色,代表樹狀結構中不可再分割的葉節點。
type File struct {
	size int
}

// NewFile 建立一個具有指定大小的 File。
func NewFile(size int) *File {
	return &File{size: size}
}

// Size 回傳檔案自身的大小。
func (f *File) Size() int {
	return f.size
}

// Folder 是組合模式中的 Composite 角色,代表可以容納其他 Node
// (包含 File 或其他 Folder)的容器節點。
type Folder struct {
	children []Node
}

// NewFolder 建立一個空的 Folder。
func NewFolder() *Folder {
	return &Folder{}
}

// Add 將一個子節點加入到這個 Folder 底下,子節點可以是 *File 也可以是 *Folder。
func (f *Folder) Add(child Node) {
	f.children = append(f.children, child)
}

// Size 遞迴加總所有子節點的大小,對呼叫者而言,Folder 與 File 的用法一致。
func (f *Folder) Size() int {
	total := 0
	for _, child := range f.children {
		total += child.Size()
	}
	return total
}
