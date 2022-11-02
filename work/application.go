package work

import (
	"github.com/ArtisanCloud/PowerLibs/v2/logger"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/yrzs/wechat-work/kernel"
	"github.com/yrzs/wechat-work/kernel/providers"
	"github.com/yrzs/wechat-work/work/accountService"
	"github.com/yrzs/wechat-work/work/accountService/customer"
	message3 "github.com/yrzs/wechat-work/work/accountService/message"
	"github.com/yrzs/wechat-work/work/accountService/serviceState"
	"github.com/yrzs/wechat-work/work/accountService/servicer"
	tag3 "github.com/yrzs/wechat-work/work/accountService/tag"
	"github.com/yrzs/wechat-work/work/agent"
	"github.com/yrzs/wechat-work/work/agent/workbench"
	"github.com/yrzs/wechat-work/work/auth"
	"github.com/yrzs/wechat-work/work/base"
	"github.com/yrzs/wechat-work/work/corpgroup"
	"github.com/yrzs/wechat-work/work/department"
	"github.com/yrzs/wechat-work/work/externalContact"
	"github.com/yrzs/wechat-work/work/externalContact/contactWay"
	"github.com/yrzs/wechat-work/work/externalContact/customerStrategy"
	"github.com/yrzs/wechat-work/work/externalContact/groupChat"
	"github.com/yrzs/wechat-work/work/externalContact/groupWelcomeTemplate"
	"github.com/yrzs/wechat-work/work/externalContact/messageTemplate"
	"github.com/yrzs/wechat-work/work/externalContact/moment"
	"github.com/yrzs/wechat-work/work/externalContact/momentStrategy"
	"github.com/yrzs/wechat-work/work/externalContact/school"
	"github.com/yrzs/wechat-work/work/externalContact/statistics"
	tag2 "github.com/yrzs/wechat-work/work/externalContact/tag"
	"github.com/yrzs/wechat-work/work/externalContact/transfer"
	"github.com/yrzs/wechat-work/work/groupRobot"
	"github.com/yrzs/wechat-work/work/invoice"
	"github.com/yrzs/wechat-work/work/media"
	"github.com/yrzs/wechat-work/work/menu"
	"github.com/yrzs/wechat-work/work/message"
	"github.com/yrzs/wechat-work/work/message/appChat"
	externalContact2 "github.com/yrzs/wechat-work/work/message/externalContact"
	linkedCorp2 "github.com/yrzs/wechat-work/work/message/linkedCorp"
	msgaudit "github.com/yrzs/wechat-work/work/msgAudit"
	"github.com/yrzs/wechat-work/work/oa"
	"github.com/yrzs/wechat-work/work/oa/calendar"
	"github.com/yrzs/wechat-work/work/oa/dial"
	"github.com/yrzs/wechat-work/work/oa/journal"
	"github.com/yrzs/wechat-work/work/oa/living"
	"github.com/yrzs/wechat-work/work/oa/meeting"
	"github.com/yrzs/wechat-work/work/oa/meetingroom"
	"github.com/yrzs/wechat-work/work/oa/pstncc"
	"github.com/yrzs/wechat-work/work/oa/schedule"
	"github.com/yrzs/wechat-work/work/oa/webdrive"
	"github.com/yrzs/wechat-work/work/oauth"
	"github.com/yrzs/wechat-work/work/server"
	"github.com/yrzs/wechat-work/work/user"
	"github.com/yrzs/wechat-work/work/user/batchJobs"
	"github.com/yrzs/wechat-work/work/user/exportJobs"
	"github.com/yrzs/wechat-work/work/user/linkedCorp"
	"github.com/yrzs/wechat-work/work/user/tag"
)

type Work struct {
	*kernel.ServiceContainer

	Base        *base.Client
	AccessToken *auth.AccessToken
	OAuth       *oauth.Manager

	Config     *kernel.Config
	Department *department.Client

	Agent          *agent.Client
	AgentWorkbench *workbench.Client

	Message                *message.Client
	Messager               *message.Messager
	MessageAppChat         *appChat.Client
	MessageExternalContact *externalContact2.Client
	MessageLinkedCorp      *linkedCorp2.Client

	Encryptor *kernel.Encryptor
	Server    *server.Guard

	User           *user.Client
	UserBatchJobs  *batchJobs.Client
	UserExportJobs *exportJobs.Client
	UserLinkedCorp *linkedCorp.Client
	UserTag        *tag.Client

	ExternalContact                     *externalContact.Client
	ExternalContactContactWay           *contactWay.Client
	ExternalContactCustomerStrategy     *customerStrategy.Client
	ExternalContactStatistics           *statistics.Client
	ExternalContactGroupWelcomeTemplate *groupWelcomeTemplate.Client
	ExternalContactSchool               *school.Client
	ExternalContactMoment               *moment.Client
	ExternalContactMomentStrategy       *momentStrategy.Client
	ExternalContactMessageTemplate      *messageTemplate.Client
	ExternalContactGroupChat            *groupChat.Client
	ExternalContactTag                  *tag2.Client
	ExternalContactTransfer             *transfer.Client

	AccountService         *accountService.Client
	AccountServiceCustomer *customer.Client
	AccountServiceMessage  *message3.Client
	AccountServiceServicer *servicer.Client
	AccountServiceState    *serviceState.Client
	AccountServiceTag      *tag3.Client

	Media *media.Client
	Menu  *menu.Client

	OA            *oa.Client
	OACalendar    *calendar.Client
	OADial        *dial.Client
	OAJournal     *journal.Client
	OALiving      *living.Client
	OAMeeting     *meeting.Client
	OAMeetingRoom *meetingroom.Client
	OAPSTNCC      *pstncc.Client
	OASchedule    *schedule.Client
	OAWebDrive    *webdrive.Client

	MsgAudit *msgaudit.Client

	CorpGroup *corpgroup.Client

	Invoice *invoice.Client

	GroupRobot          *groupRobot.Client
	GroupRobotMessenger *groupRobot.Messager

	Logger *logger.Logger
}

type UserConfig struct {
	CorpID      string
	AgentID     int
	Secret      string
	Token       string
	AESKey      string
	CallbackURL string

	ResponseType string
	Log          Log
	OAuth        OAuth
	Cache        kernel.CacheInterface

	HttpDebug bool
	Debug     bool
	NotifyURL string
	Sandbox   bool
}

type Log struct {
	Level string
	File  string
	ENV   string
}

type OAuth struct {
	Callback string
	Scopes   []string
}

func NewWork(config *UserConfig) (*Work, error) {
	var err error

	userConfig, err := MapUserConfig(config)
	if err != nil {
		return nil, err
	}

	// init an app container
	container := &kernel.ServiceContainer{
		UserConfig: userConfig,
		DefaultConfig: &object.HashMap{
			"http": &object.HashMap{
				"base_uri": "https://qyapi.weixin.qq.com/",
			},
		},
	}
	container.GetConfig()

	// init app
	app := &Work{
		ServiceContainer: container,
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	//-------------- register auth --------------
	app.AccessToken, err = auth.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register Base --------------
	app.Base, err = base.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register oauth --------------
	app.OAuth, err = oauth.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register agent --------------
	app.Agent,
		app.AgentWorkbench, err = agent.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Department --------------
	app.Department, err = department.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Message --------------
	app.Message, app.Messager,
		app.MessageAppChat,
		app.MessageExternalContact,
		app.MessageLinkedCorp, err = message.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Encryptor --------------
	app.Encryptor, app.Server, err = server.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register user --------------
	app.User,
		app.UserBatchJobs,
		app.UserExportJobs,
		app.UserLinkedCorp,
		app.UserTag, err = user.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register external contact --------------
	app.ExternalContact,
		app.ExternalContactContactWay,
		app.ExternalContactCustomerStrategy,
		app.ExternalContactGroupChat,
		app.ExternalContactGroupWelcomeTemplate,
		app.ExternalContactMessageTemplate,
		app.ExternalContactMoment,
		app.ExternalContactMomentStrategy,
		app.ExternalContactSchool,
		app.ExternalContactStatistics,
		app.ExternalContactTag,
		app.ExternalContactTransfer, err = externalContact.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register account service --------------
	app.AccountService,
		app.AccountServiceCustomer,
		app.AccountServiceMessage,
		app.AccountServiceServicer,
		app.AccountServiceState,
		app.AccountServiceTag, err = accountService.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- media --------------
	app.Media, err = media.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- menu --------------
	app.Menu, err = menu.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- oa --------------
	app.OA,
		app.OACalendar,
		app.OADial,
		app.OAJournal,
		app.OALiving,
		app.OAMeeting,
		app.OAMeetingRoom,
		app.OAPSTNCC,
		app.OASchedule,
		app.OAWebDrive, err = oa.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- msg audit --------------
	app.MsgAudit, err = msgaudit.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- corp group --------------
	app.CorpGroup, err = corpgroup.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- invoice --------------
	app.Invoice, err = invoice.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	app.GroupRobot, app.GroupRobotMessenger, err = groupRobot.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	app.Logger, err = logger.NewLogger("", &object.HashMap{
		"env":        app.Config.GetString("env", "develop"),
		"outputPath": app.Config.GetString("log.file", "./wechat.log"),
		"errorPath":  app.Config.GetString("log.file", "./wechat.log"),
	})
	if err != nil {
		return nil, err
	}

	return app, err
}

func (app *Work) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}

func (app *Work) GetAccessToken() *kernel.AccessToken {
	return app.AccessToken.AccessToken
}

func (app *Work) GetConfig() *kernel.Config {
	return app.Config
}

func (app *Work) GetComponent(name string) interface{} {

	switch name {
	case "Base":
		return app.Base
	case "AccessToken":
		return app.AccessToken
	case "OAuth":
		return app.OAuth
	case "Config":
		return app.Config
	case "Department":
		return app.Department

	case "Message":
		return app.Message
	case "Messager":
		return app.Messager
	case "MessageAppChat":
		return app.MessageAppChat
	case "MessageExternalContact":
		return app.MessageExternalContact
	case "MessageLinkedCorp":
		return app.MessageLinkedCorp

	case "Encryptor":
		return app.Encryptor
	case "Server":
		return app.Server

	case "UserClient":
		return app.User
	case "UserBatchJobsClient":
		return app.UserBatchJobs
	case "UserExportJobs":
		return app.UserExportJobs
	case "UserLinkedCorpClient":
		return app.UserLinkedCorp
	case "UserTagClient":
		return app.UserTag

	case "ExternalContact":
		return app.ExternalContact
	case "ExternalContactContactWay":
		return app.ExternalContactContactWay
	case "ExternalContactStatistics":
		return app.ExternalContactStatistics
	case "ExternalContactGroupWelcomeTemplate":
		return app.ExternalContactGroupWelcomeTemplate
	case "ExternalContactSchool":
		return app.ExternalContactSchool
	case "ExternalContactMoment":
		return app.ExternalContactMoment
	case "ExternalContactMomentStrategy":
		return app.ExternalContactMomentStrategy
	case "ExternalContactMessageTemplate":
		return app.ExternalContactMessageTemplate

	case "AccountService":
		return app.AccountService
	case "AccountServiceCustomer":
		return app.AccountServiceCustomer
	case "AccountServiceMessage":
		return app.AccountServiceMessage
	case "AccountServiceServicer":
		return app.AccountServiceServicer
	case "AccountServiceState":
		return app.AccountServiceState
	case "AccountServiceTag":
		return app.AccountServiceTag

	case "Menu":
		return app.Menu
	case "OA":
		return app.OA
	case "OACalendar":
		return app.OACalendar
	case "OADial":
		return app.OADial
	case "OAJournal":
		return app.OAJournal
	case "OALiving":
		return app.OALiving
	case "OAMeeting":
		return app.OAMeeting
	case "OAMeetingRoom":
		return app.OAMeetingRoom
	case "OAPSTNCC":
		return app.OAPSTNCC
	case "OASchedule":
		return app.OASchedule
	case "OAWebDrive":
		return app.OAWebDrive

	case "MsgAudit":
		return app.MsgAudit
	case "CorpGroup":
		return app.CorpGroup
	case "Invoice":
		return app.Invoice

	case "GroupRobot":
		return app.GroupRobot
	case "GroupRobotMessenger":
		return app.GroupRobotMessenger

	case "Logger":
		return app.Logger

	default:
		return nil
	}

}

func MapUserConfig(userConfig *UserConfig) (*object.HashMap, error) {

	config := &object.HashMap{
		"corp_id":      userConfig.CorpID,
		"agent_id":     userConfig.AgentID,
		"secret":       userConfig.Secret,
		"token":        userConfig.Token,
		"aes_key":      userConfig.AESKey,
		"callback_url": userConfig.CallbackURL,

		"response_type": userConfig.ResponseType,
		"log": &object.HashMap{
			"level": userConfig.Log.Level,
			"file":  userConfig.Log.File,
			"env":   userConfig.Log.ENV,
		},
		"oauth.callbacks": userConfig.OAuth.Callback,
		"oauth.scopes":    userConfig.OAuth.Scopes,
		"cache":           userConfig.Cache,
		"http_debug":      userConfig.HttpDebug,
		"debug":           userConfig.Debug,
	}

	return config, nil

}

