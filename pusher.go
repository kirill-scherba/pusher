// Copyright 2022 Kirill Scherba <kirill@scherba.ru>.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Pusher send metrics to Prometheus Push Server using http post request
package pusher

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Pusher is package methods receiver
type Pusher struct {
	*http.Client
	url      string
	job      string
	instance string
}

// New creates new Pusher client
func New(prometheusPushURL, job, instance string) (pusher *Pusher) {
	pusher = &Pusher{&http.Client{}, prometheusPushURL, job, instance}
	return
}

// Push metrics to prometheus push server
func (p *Pusher) Push(metrics ...string) (metricsRaw string, err error) {

	// Create reader from metrics
	metricsRaw = strings.TrimSpace(strings.Join(metrics, "\n"))
	if len(metricsRaw) == 0 {
		err = errors.New("input metrics list is empty")
		return
	}
	metricsRaw += "\n"
	reader := strings.NewReader(metricsRaw)

	// Create and send http post request with body
	req, err := http.NewRequest(
		"POST",
		p.url+"/metrics/job/"+p.job+"/instance/"+p.instance,
		reader,
	)
	if err != nil {
		return
	}
	_, err = p.Do(req)

	return
}

// Metric make metric string from name and value
func (p *Pusher) Metric(name string, value float64) string {
	return Metric(name, value)
}

// Metric make metric string from name and value
func Metric(name string, value float64) string {
	return fmt.Sprintf("%s %f", name, value)
}
