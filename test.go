package spec

import "github.com/open-fsm/spec/proto"

var (
	v0o1,v1o1,v1o2,v1o3,v1o4,v1o5,v1o6,v1o7,v2o0,v2o2,v2o4,v2o5,v2o6,v2o7,v2o8,v3o2,v2o1,v2o3,v3o3,
	v3o4,v4o3,v4o4,v4o5,v4o6,v4o1,v4o2 proto.ViewStamp
)

func InitViewStampCase() {
	v0o1 = proto.ViewStamp{ViewNum: 0, OpNum: 1}
	v1o1 = proto.ViewStamp{ViewNum: 1, OpNum: 1}
	v1o2 = proto.ViewStamp{ViewNum: 1, OpNum: 2}
	v1o3 = proto.ViewStamp{ViewNum: 1, OpNum: 3}
	v1o4 = proto.ViewStamp{ViewNum: 1, OpNum: 4}
	v1o5 = proto.ViewStamp{ViewNum: 1, OpNum: 5}
	v1o6 = proto.ViewStamp{ViewNum: 1, OpNum: 6}
	v1o7 = proto.ViewStamp{ViewNum: 1, OpNum: 7}
	v2o0 = proto.ViewStamp{ViewNum: 2, OpNum: 0}
	v2o2 = proto.ViewStamp{ViewNum: 2, OpNum: 2}
	v2o4 = proto.ViewStamp{ViewNum: 2, OpNum: 4}
	v2o5 = proto.ViewStamp{ViewNum: 2, OpNum: 5}
	v2o6 = proto.ViewStamp{ViewNum: 2, OpNum: 6}
	v2o7 = proto.ViewStamp{ViewNum: 2, OpNum: 7}
	v2o8 = proto.ViewStamp{ViewNum: 2, OpNum: 8}
	v3o2 = proto.ViewStamp{ViewNum: 3, OpNum: 2}
	v3o4 = proto.ViewStamp{ViewNum: 3, OpNum: 4}
	v2o1 = proto.ViewStamp{ViewNum: 2, OpNum: 1}
	v2o3 = proto.ViewStamp{ViewNum: 2, OpNum: 3}
	v3o3 = proto.ViewStamp{ViewNum: 3, OpNum: 3}
	v4o3 = proto.ViewStamp{ViewNum: 4, OpNum: 3}
	v4o4 = proto.ViewStamp{ViewNum: 4, OpNum: 4}
	v4o5 = proto.ViewStamp{ViewNum: 4, OpNum: 5}
	v4o6 = proto.ViewStamp{ViewNum: 4, OpNum: 6}
	v4o1 = proto.ViewStamp{ViewNum: 4, OpNum: 1}
	v4o2 = proto.ViewStamp{ViewNum: 4, OpNum: 2}
}