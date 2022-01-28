package segment

import (
	"io"

	"FalconEngine/index/invert"
	"FalconEngine/tools"
)

type FalconSegmentService interface {
	AddField(mapping *tools.FalconMapping) error
	UpdateDocument(document map[string]interface{}) error
	// 持久化
	Persistence() error
	// 关闭
	io.Closer
	ToString() string

	Number() uint32
	DocumentCount() uint32
	Name() string

	SimpleSearch(field, keyword string) (invert.FalconDocList, bool, error)
}
