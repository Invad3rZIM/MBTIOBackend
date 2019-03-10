package main

type MessageCache struct {
	Inbox map[int]map[int]*[]*Message
}

func NewMessageCache() *MessageCache {
	return &MessageCache{
		Inbox: make(map[int]map[int]*[]*Message),
	}
}

//adds mailbox to inbox
func (mc *MessageCache) AddMailbox(rid int) {
	mc.Inbox[rid] = make(map[int]*[]*Message)
}

//adds cubby to mailbox in inbox
func (mc *MessageCache) AddCubby(rid int, sid int) {
	messageCubby := []*Message{}
	mc.Inbox[rid][sid] = &messageCubby
}

//adds message to cubby in inbox
func (mc *MessageCache) PostMessage(message *Message) {

	//if recipient has never received mail before
	if _, ok := mc.Inbox[message.Rid]; !ok {
		mc.AddMailbox(message.Rid)
	}

	//if sender has never sent message to this particular recipient before
	if _, ok := mc.Inbox[message.Rid][message.Sid]; !ok {
		mc.AddCubby(message.Rid, message.Sid)
	}

	mbox := *mc.Inbox[message.Rid][message.Sid]
	mbox = append(mbox, message)
	mc.Inbox[message.Rid][message.Sid] = &mbox
}

func (mc *MessageCache) GetMessages(rid int) []*Message {
	m := []*Message{}

	recipientMailbox, ok := mc.Inbox[rid]

	//if there's no mailbox found, return empty array after creating a mailbox
	if !ok {
		mc.AddMailbox(rid)
		return m
	}

	for sender, messages := range recipientMailbox {
		m = append(m, (*messages)...)    //add mail to return list
		delete(recipientMailbox, sender) //delete mailbox after done for cleanup
	}

	return m
}
