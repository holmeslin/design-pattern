// Package mediator 示範中介者模式(Mediator Pattern)。
//
// 中介者模式將物件之間原本錯綜複雜的直接溝通,改由一個中介者(Mediator)
// 集中處理,讓各個同事物件(Colleague)彼此不需要知道對方的存在,
// 只需要跟中介者互動,藉此降低耦合度。
//
// 本範例以聊天室為情境:
//   - Mediator 角色:ChatRoom,負責將訊息轉發給聊天室內其他成員。
//   - Colleague 角色:User,只認識 ChatRoom,不直接持有其他 User 的參照。
package mediator

import "fmt"

// Mediator 定義中介者對外提供的行為:將訊息從某個成員廣播出去。
type Mediator interface {
	// Broadcast 將 msg 由 from 發送,轉發給其他所有成員。
	Broadcast(from string, msg string)
}

// ChatRoom 是具體中介者(Concrete Mediator),
// 持有所有加入聊天室的 User,並負責訊息轉發邏輯。
type ChatRoom struct {
	users []*User // 依加入順序保存,確保轉發順序穩定
}

// NewChatRoom 建立一個空的聊天室。
func NewChatRoom() *ChatRoom {
	return &ChatRoom{}
}

// register 將 user 加入聊天室,依加入順序附加在 users 尾端。
func (r *ChatRoom) register(u *User) {
	r.users = append(r.users, u)
}

// Broadcast 依加入順序,將訊息轉發給除了 from 之外的每一位成員。
func (r *ChatRoom) Broadcast(from string, msg string) {
	for _, u := range r.users {
		if u.name == from {
			continue
		}
		u.receive(from, msg)
	}
}

// User 是具體同事(Concrete Colleague),
// 只透過 Mediator 介面與聊天室溝通,不直接依賴其他 User。
type User struct {
	name string
	room Mediator
}

// NewUser 建立一個 User 並將其登記到指定的 ChatRoom 中。
func NewUser(name string, room *ChatRoom) *User {
	u := &User{name: name, room: room}
	room.register(u)
	return u
}

// Send 透過中介者將訊息廣播給聊天室內的其他成員。
func (u *User) Send(msg string) {
	u.room.Broadcast(u.name, msg)
}

// receive 由中介者呼叫,代表這個 User 收到了某人送來的訊息。
func (u *User) receive(from string, msg string) {
	fmt.Printf("%s received from %s: %s\n", u.name, from, msg)
}
