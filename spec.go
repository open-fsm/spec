package spec

import "github.com/open-fsm/spec/proto"

var (
	V0o1,V1o1,V1o2,V1o3,V1o4,V1o5,V1o6,V1o7,V2o0,V2o2,V2o4,V2o5,V2o6,V2o7,V2o8,V3o2,V2o1,V2o3,V3o3,
	V3o4,V4o3,V4o4,V4o5,V4o6,V4o1,V4o2 proto.ViewStamp
)

func InitViewStampCase() {
	V0o1 = proto.ViewStamp{ViewNum: 0, OpNum: 1}
	V1o1 = proto.ViewStamp{ViewNum: 1, OpNum: 1}
	V1o2 = proto.ViewStamp{ViewNum: 1, OpNum: 2}
	V1o3 = proto.ViewStamp{ViewNum: 1, OpNum: 3}
	V1o4 = proto.ViewStamp{ViewNum: 1, OpNum: 4}
	V1o5 = proto.ViewStamp{ViewNum: 1, OpNum: 5}
	V1o6 = proto.ViewStamp{ViewNum: 1, OpNum: 6}
	V1o7 = proto.ViewStamp{ViewNum: 1, OpNum: 7}
	V2o0 = proto.ViewStamp{ViewNum: 2, OpNum: 0}
	V2o2 = proto.ViewStamp{ViewNum: 2, OpNum: 2}
	V2o4 = proto.ViewStamp{ViewNum: 2, OpNum: 4}
	V2o5 = proto.ViewStamp{ViewNum: 2, OpNum: 5}
	V2o6 = proto.ViewStamp{ViewNum: 2, OpNum: 6}
	V2o7 = proto.ViewStamp{ViewNum: 2, OpNum: 7}
	V2o8 = proto.ViewStamp{ViewNum: 2, OpNum: 8}
	V3o2 = proto.ViewStamp{ViewNum: 3, OpNum: 2}
	V3o4 = proto.ViewStamp{ViewNum: 3, OpNum: 4}
	V2o1 = proto.ViewStamp{ViewNum: 2, OpNum: 1}
	V2o3 = proto.ViewStamp{ViewNum: 2, OpNum: 3}
	V3o3 = proto.ViewStamp{ViewNum: 3, OpNum: 3}
	V4o3 = proto.ViewStamp{ViewNum: 4, OpNum: 3}
	V4o4 = proto.ViewStamp{ViewNum: 4, OpNum: 4}
	V4o5 = proto.ViewStamp{ViewNum: 4, OpNum: 5}
	V4o6 = proto.ViewStamp{ViewNum: 4, OpNum: 6}
	V4o1 = proto.ViewStamp{ViewNum: 4, OpNum: 1}
	V4o2 = proto.ViewStamp{ViewNum: 4, OpNum: 2}
}