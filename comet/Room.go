package main

import (
	"im/libs/proto"
	"sync"
)

type Room struct {
	Id     int32 // room id
	rlock  sync.RWMutex
	next   *Channel // all channel of this room
	drop   bool     // whether the room is dropped or not
	Online int      // number of online users in this room
}

func NewRoom(Id int32) (r *Room) {
	r = new(Room)
	r.Id = Id
	r.drop = false
	r.next = nil
	r.Online = 0
	return
}

// put new user
func (r *Room) Put(ch *Channel) (err error) {
	if !r.drop {
		if r.next != nil {
			r.next.Prev = ch
		}
		ch.Next = r.next
		ch.Prev = nil
		r.next = ch
		r.Online++
	} else {
		err = ErrorRoomDropped
	}
	return
}

// push message
func (r *Room) Push(p *proto.Proto) {
	r.rlock.RLock()
	for ch := r.next; ch != nil; ch = ch.Next {
		ch.Push(p)
	}
	r.rlock.Unlock()
	return
}

func (r *Room) Del(ch *Channel) bool {
	r.rlock.RLock()
	if ch.Next != nil {
		// if not footer
		ch.Next.Prev = ch.Prev
	}

	if ch.Prev != nil {
		// if not header
		ch.Prev.Next = ch.Next

	} else {

		r.next = ch.Next
	}
	r.Online--
	r.drop = (r.Online == 0)
	r.rlock.RUnlock()
	return r.drop
}
