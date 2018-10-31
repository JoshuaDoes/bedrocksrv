package protocol

import (
	"github.com/JoshuaDoes/bedrocksrv/entity"
	//"github.com/JoshuaDoes/bedrocksrv/network/bedrock/handler"
)

//PacketLogin exports the methods required to handle the Login packet
type PacketLogin struct {
	IsDecoded         bool
	Username          string
	Protocol          int
	ClientUUID        string
	ClientID          int
	XUID              string
	IdentityPublicKey string
	ServerAddress     string
	Locale            string
	Skin              *entity.Skin
}

//GetCanBeSentBeforeLogin returns whether this packet can be sent before a successful login
func (*PacketLogin) GetCanBeSentBeforeLogin() bool {
	return true
}

//GetIsDecoded returns whether this packet is decoded yet or not
func (packet *PacketLogin) GetIsDecoded() bool {
	return packet.IsDecoded
}

//GetMayHaveUnreadBytes returns whether this packet may legally have unread bytes left in the buffer
func (packet *PacketLogin) GetMayHaveUnreadBytes() bool {
	if packet.Protocol == 0 {
		return false
	}
	if packet.Protocol != CurrentProtocol {
		return false
	}
	return true
}

//GetPacketID returns the packet ID
func (*PacketLogin) GetPacketID() int {
	return OpLogin
}

//GetPacketName returns the packet name
func (*PacketLogin) GetPacketName() string {
	return "login"
}

//Decode decodes the packet
func (*PacketLogin) Decode(payload []byte) error {

}
