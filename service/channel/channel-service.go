package channel

import (
	"github.com/ESPOIR-DITE/tim-api/domain/channel/channel"
)

type ChannelMax struct {
	Channel channel.Channel
	Max     int64
}

//func getHighest(channelList []channel.Channel) []ChannelMax {
//	var channelLists []ChannelMax
//	for _, channel := range channelList {
//		channelLists = append(channelLists, ChannelMax{channel, role_repo2.CountSubscriptionByChannelId(channel.Id)})
//	}
//	return channelLists
//}
//
//func getTrio() {
//	var trioList []ChannelMax
//	channels := role_repo.GetChannels()
//	if len(channels) == 0 {
//		return
//	}
//	channelSubMax := getHighest(channels)
//
//	for i := 0; i < 3; i++ {
//		result := findMax(channelSubMax)
//		if result.Channel.Id != "" {
//			trioList = append(trioList, result)
//		}
//	}
//
//}
//func findMax(a []ChannelMax) ChannelMax {
//	max := a[0].Max
//	var maxValue ChannelMax
//	for _, b := range a {
//		if b.Max > max {
//			maxValue = ChannelMax{}
//			max = b.Max
//			maxValue = b
//		}
//	}
//	return maxValue
//}
