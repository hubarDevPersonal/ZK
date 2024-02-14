package players

type Player struct {
	Symbol   string
	Nickname string
}

func (p *Player) SetNickname(nickname string) {
	p.Nickname = nickname
}
