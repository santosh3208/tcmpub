package TCMPub

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"context"
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/jvanderl/tib-eftl"
	"strconv"
	"crypto/x509"
	"crypto/tls"
	"encoding/base64"
	"fmt"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

    wsURL := t.config.GetSetting("url")
	wsAuthKey := t.config.GetSetting("authkey")	// do eval

	errChan := make(chan error, 1)

	opts := &eftl.Options{
	
		Password:  wsAuthKey,
		
	}

	conn, err := eftl.Connect(wsURL, opts, errChan)
	if err != nil {
		log.Errorf("Error connecing to TCM server: [%s]", err)
		return err
	}

	// close the connection when done
	defer conn.Disconnect()

	msg :=t.config.GetSetting("message")


	err = conn.Publish(msg)

	return true, nil
}
