package merEmun

//任务状态枚举
type MerStatus int

//状态 0 未开始 1 进行中  2 已结束\
const (
	Verified MerStatus = 0 // 审核通过

	Review = -1 //审核中
)

type MerUserStatus int

const (
	MerUserStatusPass MerUserStatus = 0 // 审核通过

	MerUserStatusReview = -1 //审核中
)

type MerApplyStatus int

const (
	MerApplyStatusReview MerApplyStatus = 0 // 审核中

	MerApplyStatusVerified = 1 //审核通过

	MerApplyStatusFail = -1 //审核失败
)

func (s MerApplyStatus) Text() string {
	switch s {
	case MerApplyStatusReview:
		return "审核处理中"
	case MerApplyStatusVerified:
		return "审核通过"
	case MerApplyStatusFail:
		return "审核失败"

	default:
		return ""
	}
}
