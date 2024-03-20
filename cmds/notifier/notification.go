package main

import (
	"fmt"

	pbnotify "github.com/safing/portbase/notifications"
)

// Notification represents a notification that is to be delivered to the user.
type Notification struct {
	pbnotify.Notification

	// systemID holds the ID returned by the dbus interface on Linux or by WinToast library on Windows.
	systemID NotificationID
}

// IsSupported returns whether the action is supported on this system.
func IsSupportedAction(a pbnotify.Action) bool {
	switch a.Type {
	case pbnotify.ActionTypeNone:
		return true
	default:
		return false
	}
}

// SelectAction sends an action back to the portmaster.
func (n *Notification) SelectAction(action string) {
	new := &pbnotify.Notification{
		EventID:          n.EventID,
		SelectedActionID: action,
	}

	// FIXME: check response
	apiClient.Update(fmt.Sprintf("%s%s", dbNotifBasePath, new.EventID), new, nil)
}
