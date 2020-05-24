package Peers
type Peers struct {
	//持币数量
	Token int
	//持币时间
	Days int
	//节点地址
	Address string
}
func CreatePeers(token int,days int,address string)Peers{
	var peer Peers
	peer.Address=address
	peer.Days=days
	peer.Token=token
	return peer
}