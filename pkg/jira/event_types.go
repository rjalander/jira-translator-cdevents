/*
Copyright (C) 2024 Nordix Foundation.
For a full list of individual contributors, please see the commit history.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
SPDX-License-Identifier: Apache-2.0
*/

package jira

// Define structs for nested JSON objects

type AvatarUrls struct {
	X48 string `json:"48x48"`
	X24 string `json:"24x24"`
	X16 string `json:"16x16"`
	X32 string `json:"32x32"`
}

type User struct {
	Self         string     `json:"self"`
	Name         string     `json:"name"`
	Key          string     `json:"key"`
	EmailAddress string     `json:"emailAddress"`
	AvatarUrls   AvatarUrls `json:"avatarUrls"`
	DisplayName  string     `json:"displayName"`
	Active       bool       `json:"active"`
	TimeZone     string     `json:"timeZone"`
}

type Project struct {
	Self           string     `json:"self"`
	ID             string     `json:"id"`
	Key            string     `json:"key"`
	Name           string     `json:"name"`
	ProjectTypeKey string     `json:"projectTypeKey"`
	AvatarUrls     AvatarUrls `json:"avatarUrls"`
}

type IssueType struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	IconUrl     string `json:"iconUrl"`
	Name        string `json:"name"`
	Subtask     bool   `json:"subtask"`
}

type Priority struct {
	Self    string `json:"self"`
	IconUrl string `json:"iconUrl"`
	Name    string `json:"name"`
	ID      string `json:"id"`
}

type Watches struct {
	Self       string `json:"self"`
	WatchCount int    `json:"watchCount"`
	IsWatching bool   `json:"isWatching"`
}

type StatusCategory struct {
	Self      string `json:"self"`
	ID        int    `json:"id"`
	Key       string `json:"key"`
	ColorName string `json:"colorName"`
	Name      string `json:"name"`
}

type Status struct {
	Self           string         `json:"self"`
	Description    string         `json:"description"`
	IconUrl        string         `json:"iconUrl"`
	Name           string         `json:"name"`
	ID             string         `json:"id"`
	StatusCategory StatusCategory `json:"statusCategory"`
}

type Votes struct {
	Self     string `json:"self"`
	Votes    int    `json:"votes"`
	HasVoted bool   `json:"hasVoted"`
}

type Progress struct {
	Progress int `json:"progress"`
	Total    int `json:"total"`
}

type AggregateProgress struct {
	Progress int `json:"progress"`
	Total    int `json:"total"`
}

type IssueFields struct {
	Issuetype                     IssueType         `json:"issuetype"`
	Timespent                     *int              `json:"timespent"`
	Project                       Project           `json:"project"`
	FixVersions                   []interface{}     `json:"fixVersions"`
	Customfield_10110             string            `json:"customfield_10110"`
	Customfield_10111             any               `json:"customfield_10111"`
	Aggregatetimespent            *interface{}      `json:"aggregatetimespent"`
	Resolution                    *interface{}      `json:"resolution"`
	Customfield_10107             *interface{}      `json:"customfield_10107"`
	Customfield_10108             *interface{}      `json:"customfield_10108"`
	Customfield_10109             *interface{}      `json:"customfield_10109"`
	Resolutiondate                *interface{}      `json:"resolutiondate"`
	Workratio                     int               `json:"workratio"`
	LastViewed                    *interface{}      `json:"lastViewed"`
	Watches                       Watches           `json:"watches"`
	Created                       string            `json:"created"`
	Priority                      Priority          `json:"priority"`
	Customfield_10100             *interface{}      `json:"customfield_10100"`
	Customfield_10101             *interface{}      `json:"customfield_10101"`
	Customfield_10102             *interface{}      `json:"customfield_10102"`
	Labels                        []interface{}     `json:"labels"`
	Customfield_10103             *interface{}      `json:"customfield_10103"`
	Timeestimate                  *interface{}      `json:"timeestimate"`
	Aggregatetimeoriginalestimate *interface{}      `json:"aggregatetimeoriginalestimate"`
	Versions                      []interface{}     `json:"versions"`
	Issuelinks                    []interface{}     `json:"issuelinks"`
	Assignee                      User              `json:"assignee"`
	Updated                       string            `json:"updated"`
	Status                        Status            `json:"status"`
	Components                    []interface{}     `json:"components"`
	Timeoriginalestimate          *interface{}      `json:"timeoriginalestimate"`
	Description                   string            `json:"description"`
	Timetracking                  struct{}          `json:"timetracking"`
	Archiveddate                  *interface{}      `json:"archiveddate"`
	Attachment                    []interface{}     `json:"attachment"`
	Aggregatetimeestimate         *interface{}      `json:"aggregatetimeestimate"`
	Summary                       string            `json:"summary"`
	Creator                       User              `json:"creator"`
	Subtasks                      []interface{}     `json:"subtasks"`
	Reporter                      User              `json:"reporter"`
	Customfield_10000             string            `json:"customfield_10000"`
	Aggregateprogress             AggregateProgress `json:"aggregateprogress"`
	Environment                   *interface{}      `json:"environment"`
	Duedate                       *interface{}      `json:"duedate"`
	Progress                      Progress          `json:"progress"`
	Comment                       struct {
		Comments   []interface{} `json:"comments"`
		MaxResults int           `json:"maxResults"`
		Total      int           `json:"total"`
		StartAt    int           `json:"startAt"`
	} `json:"comment"`
	Votes   Votes `json:"votes"`
	Worklog struct {
		StartAt    int           `json:"startAt"`
		MaxResults int           `json:"maxResults"`
		Total      int           `json:"total"`
		Worklogs   []interface{} `json:"worklogs"`
	} `json:"worklog"`
	Archivedby *interface{} `json:"archivedby"`
}

type Issue struct {
	ID     string      `json:"id"`
	Self   string      `json:"self"`
	Key    string      `json:"key"`
	Fields IssueFields `json:"fields"`
}

type IssueCreatedEvent struct {
	Timestamp     float64 `json:"timestamp"`
	WebhookEvent  string  `json:"webhookEvent"`
	IssueTypeName string  `json:"issue_event_type_name"`
	User          User    `json:"user"`
	Issue         Issue   `json:"issue"`
}
