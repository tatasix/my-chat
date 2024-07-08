package upload

const (
	ImageFeedback    = 1
	ImageCommonTarot = 2
	File             = 3
	Scale            = 4
)

var Path map[uint32]string

func init() {
	Path = make(map[uint32]string)
	Path[ImageFeedback] = "image/feedback"
	Path[ImageCommonTarot] = "image/common/tarot"
	Path[Scale] = "image/common/scale"
	Path[File] = "file/common"
}

func GetFilePath(fileType uint32) string {
	if msg, ok := Path[fileType]; ok {
		return msg
	} else {
		return Path[ImageFeedback]
	}
}
