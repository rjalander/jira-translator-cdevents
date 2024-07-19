1. The Jira event JSON is too long to create and maintain as go struct 
2. The below event structures needs to be maintained for any changes in future

````json
jira:issue_created ==> dev.cdevents.ticket.created
{
	"timestamp": 1716991458232,
	"webhookEvent": "jira:issue_created",
	"issue_event_type_name": "issue_created",
	"user": {},
	"issue": {}
}


jira:issue_updated & jira:worklog_updated ==> dev.cdevents.ticket.updated
{
	"timestamp": 1716992387526,
	"webhookEvent": "jira:issue_updated",
	"issue_event_type_name": "issue_updated",
	"user": {},
	"issue": {},
	"changelog": {}	
}
{
	"timestamp": 1716992564052,
	"webhookEvent": "jira:issue_updated",
	"issue_event_type_name": "issue_commented",
	"user": {},
	"issue": {},
	"comment": {}	
}
{
	"timestamp": 1717163183565,
	"webhookEvent": "jira:issue_updated",
	"issue_event_type_name": "issue_assigned",
	"user": {},
	"issue": {},
	"changelog": {},
	"comment": {}	
}
{
	"timestamp": 1717164468987,
	"webhookEvent": "jira:worklog_updated",
	"issue_event_type_name": "issue_work_logged",
	"user": {},
	"issue": {},
	"changelog": {}
}
{
	"timestamp": 1717164704071,
	"webhookEvent": "jira:worklog_updated",
	"issue_event_type_name": "issue_worklog_updated",
	"user": {},
	"issue": {},
	"changelog": {}
}
{	
	"timestamp": 1717167743759,
	"webhookEvent": "jira:worklog_updated",
	"issue_event_type_name": "issue_worklog_deleted",
	"user": {},
	"issue": {},
	"changelog": {}
}

jira:issue_updated ==> dev.cdevents.ticket.closed
{
	"timestamp": 1717165102467,
	"webhookEvent": "jira:issue_updated",
	"issue_event_type_name": "issue_generic",
	"user": {},
	"issue": {},
	"changelog": {}	
}
````