package client

import (
	"github.com/grayrail000/MiraiGo/client/internal/network"
	"github.com/grayrail000/MiraiGo/client/pb/msg"
)

var privateMsgDecoders = map[int32]func(*QQClient, *msg.Message, *network.Packet){
	9: privateMessageDecoder, 10: privateMessageDecoder, 31: privateMessageDecoder,
	79: privateMessageDecoder, 97: privateMessageDecoder, 120: privateMessageDecoder,
	132: privateMessageDecoder, 133: privateMessageDecoder, 166: privateMessageDecoder,
	167: privateMessageDecoder, 140: tempSessionDecoder, 141: tempSessionDecoder,
	208: privatePttDecoder,
}

var nonSvcNotifyTroopSystemMsgDecoders = map[int32]func(*QQClient, *msg.Message, *network.Packet){
	36: troopSystemMessageDecoder, 85: troopSystemMessageDecoder,
}

var troopSystemMsgDecoders = map[int32]func(*QQClient, *msg.Message, *network.Packet){
	35: troopSystemMessageDecoder, 37: troopSystemMessageDecoder,
	45: troopSystemMessageDecoder, 46: troopSystemMessageDecoder, 84: troopSystemMessageDecoder,
	86: troopSystemMessageDecoder, 87: troopSystemMessageDecoder,
} // IsSvcNotify

var sysMsgDecoders = map[int32]func(*QQClient, *msg.Message, *network.Packet){
	187: systemMessageDecoder, 188: systemMessageDecoder, 189: systemMessageDecoder,
	190: systemMessageDecoder, 191: systemMessageDecoder,
} // IsSvcNotify

var otherDecoders = map[int32]func(*QQClient, *msg.Message, *network.Packet){
	33: troopAddMemberBroadcastDecoder, 529: msgType0x211Decoder,
}
