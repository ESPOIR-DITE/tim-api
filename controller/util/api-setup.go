package util

import (
	role_repo2 "tim-api/storage/chanel/channel-video-repository"
	"tim-api/storage/security"
	role_repo "tim-api/storage/user/role-repo"
	user_account_repo "tim-api/storage/user/user-account-repo"
	user_repo3 "tim-api/storage/user/user-bank"
	user_repo2 "tim-api/storage/user/user-details"
	user_repo "tim-api/storage/user/user-repo"
	user_sub_repo "tim-api/storage/user/user-sub-repo"
	user_video_repo "tim-api/storage/user/user-video-repo"
	video_repo2 "tim-api/storage/video/category"
	video_category "tim-api/storage/video/video-category"
	video_category2 "tim-api/storage/video/video-comment"
	video_data_repo "tim-api/storage/video/video-data"
	video_reaction_repo "tim-api/storage/video/video-reaction-repo"
	video_related "tim-api/storage/video/video-related"
	video_repo "tim-api/storage/video/video-repo"

	channel_repo "tim-api/storage/chanel/channel-repository"
	channel_type_repo "tim-api/storage/chanel/channel-type-repository"
	channel_subscription_repo "tim-api/storage/chanel/channelSubscription"
)

func TableSetUp() []TableSetUpReport {
	var result []TableSetUpReport
	result = append(result, TableSetUpReport{"USER TABLE", user_repo.CreateUserTable()})
	result = append(result, TableSetUpReport{"USER- ACCOUNT TABLE", user_account_repo.CreateUserAccountTable()})
	result = append(result, TableSetUpReport{"USER- SUBSCRIPTION TABLE", user_sub_repo.CreateUserSubscriptionTable()})
	result = append(result, TableSetUpReport{"USER- VIDEO TABLE", user_video_repo.CreateUserVideoTable()})
	result = append(result, TableSetUpReport{"ROLE TABLE", role_repo.CreateRoleTable()})
	result = append(result, TableSetUpReport{"USER-DETAIL TABLE", user_repo2.CreateUserDetailsTable()})
	result = append(result, TableSetUpReport{"USER-BANK TABLE", user_repo3.CreateUserBankTable()})

	result = append(result, TableSetUpReport{"VIDEO TABLE", video_repo.CreateVideoTable()})
	result = append(result, TableSetUpReport{"VIDEO DATA TABLE", video_data_repo.CreateVideoDataTable()})
	result = append(result, TableSetUpReport{"VIDEO- CATEGORY TABLE", video_category.CreateVideoCategoryTable()})
	result = append(result, TableSetUpReport{"VIDEO- COMMENT TABLE", video_category2.CreateVideoCommentTable()})
	result = append(result, TableSetUpReport{"CATEGORY TABLE", video_repo2.CreateCategoryTable()})
	result = append(result, TableSetUpReport{"VIDEO REACTION TABLE", video_reaction_repo.CreateVideoReactionTable()})
	result = append(result, TableSetUpReport{"VIDEO RELATED TABLE", video_related.CreateVideoRelatedTable()})

	result = append(result, TableSetUpReport{"CHANNEL TABLE", channel_repo.CreateChannelTable()})
	result = append(result, TableSetUpReport{"CHANNEL-TYPE TABLE", channel_type_repo.CreateChannelTypeTable()})
	result = append(result, TableSetUpReport{"CHANNEL-SUBSCRIPTION TABLE", channel_subscription_repo.CreateChannelSubscriptionTable()})
	result = append(result, TableSetUpReport{"CHANNEL-VIDEO TABLE", role_repo2.CreateChannelVideoTable()})

	result = append(result, TableSetUpReport{"SYSTEM SECURITY TABLE", security.CreateSystemSecurityTable()})

	return result
}
