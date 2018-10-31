package protocol

//DataPacket is the interface that all protocol packets will extend from
type DataPacket interface {
	GetCanBeSentBeforeLogin() bool
	GetIsDecoded() bool
	GetMayHaveUnreadBytes() bool
	GetPacketID() int
	GetPacketName() string
	GetRecipentSubID() int
	GetSenderSubID() int

	Decode(payload []byte) error
	Encode(packet *DataPacket) error
}
