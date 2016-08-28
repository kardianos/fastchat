package fcd

import (
	"context"
	"time"
)

type (
	UserID         int64
	ConversationID int64
)

// User in DB.
type User struct {
	ID       UserID
	Name     string
	Email    string
	Phone    string
	PassHash string
	IsAdmin  bool

	Frustration Frustration
}

// Resource represents a logical client program.
// It is important to maintain a server side list of all
// logical resources to know which clients to notify when
// a user has an incoming request.
type Resource struct {
	Name                string
	User                *User
	ClientUUID          [16]byte
	SendNotificationsTo bool
	Status              Status
}

// UserFrustration in DB.
// Records user frustration over time, only update when no user record is
// found or when user frustration changes.
type UserFrustration struct {
	At          time.Time
	User        *User
	Frustration Frustration
}

type Status int

const (
	StatusOffline Status = iota
	StatusAvailable
	StatusBusy
	StatusAway
)

type ConversationState int

const (
	CStateGone ConversationState = iota
	CStateNew
	CStateVisible
	CStateHidden
)

type Frustration byte // 0 No frustration, 100 extreemly frustrated.

type Presence struct {
	User      *User
	Status    Status
	StatusMsg string
}

// Group in DB.
type Group struct {
	ID   int64
	Name string
	In   []*User
	See  []*User
}

// Conversation in DB.
type Conversation struct {
	ID    ConversationID
	Name  string
	Start time.Time
	List  []*User
}

// Message in DB.
type Message struct {
	ID           int64
	From         *User
	Conversation *Conversation
	At           time.Time
	Text         string
	Attachment   *Attachment
	Pin          bool
	ReplyTo      *Message
}

// Attachment in DB.
type Attachment struct {
	Hash     [64]byte
	Filename string
	Data     []byte
}

// UserConversation in DB.
type UserConversation struct {
	Conversation *Conversation
	User         *User
	State        ConversationState

	Frustration Frustration
	RaisedHand  bool
}

// From DB.
var (
	Users                []*Presence
	OpenUserConversation []*UserConversation
)

type Server struct{}

func (s *Server) SendMessage(ctx context.Context, from UserID, on ConversationID, at time.Time, text string, hash *[64]byte) error {
	return nil
}

func (s *Server) CreateConversation(ctx context.Context, from UserID, cname string) (ConversationID, error) {
	return 0, nil
}

func (s *Server) AddToConversation(ctx context.Context, cid ConversationID, user UserID) error {
	return nil
}
func (s *Server) RemoveFromConversation(ctx context.Context, cid ConversationID, user UserID) error {
	return nil
}

// UpdateConversation should take N arguments in last arguments to:
// add or remove users from conversations, set frustration, raise or lower hand, change conversation state, update name.
func (s *Server) UpdateConversation(ctx context.Context, cid ConversationID) error {
	return nil
}
