package oa

import (
	"github.com/yrzs/wechat-work/kernel"
	"github.com/yrzs/wechat-work/work/oa/calendar"
	"github.com/yrzs/wechat-work/work/oa/dial"
	"github.com/yrzs/wechat-work/work/oa/journal"
	"github.com/yrzs/wechat-work/work/oa/living"
	"github.com/yrzs/wechat-work/work/oa/meeting"
	"github.com/yrzs/wechat-work/work/oa/meetingroom"
	"github.com/yrzs/wechat-work/work/oa/pstncc"
	"github.com/yrzs/wechat-work/work/oa/schedule"
	"github.com/yrzs/wechat-work/work/oa/webdrive"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client,
	*calendar.Client,
	*dial.Client,
	*journal.Client,
	*living.Client,
	*meeting.Client,
	*meetingroom.Client,
	*pstncc.Client,
	*schedule.Client,
	*webdrive.Client,
	error,

) {
	//config := app.GetConfig()

	Client, err := NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	Calendar, err := calendar.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	Dial, err := dial.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	Journal, err := journal.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	Living, err := living.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	Meeting, err := meeting.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	MeetingRoom, err := meetingroom.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	PSTNCC, err := pstncc.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	Schedule, err := schedule.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	WebDrive, err := webdrive.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	return Client,
		Calendar,
		Dial,
		Journal,
		Living,
		Meeting,
		MeetingRoom,
		PSTNCC,
		Schedule,
		WebDrive,
		nil

}
