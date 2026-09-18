/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2016 Gin-Gonic
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.

 * This file may have been modified by CloudWeGo authors. All CloudWeGo
 * Modifications are Copyright 2022 CloudWeGo Authors.
 */

package cors

import (
	"bytes"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type cors struct {
	allowAllOrigins  bool
	allowCredentials bool
	allowOriginFunc  func(string) bool
	allowOrigins     []string
	normalHeaders    map[string]string
	preflightHeaders map[string]string
	wildcardOrigins  [][]string
}

var (
	DefaultHeaderBytes = [][]byte{
		[]byte("OPTIONS"),
		[]byte("GET"),
		[]byte("POST"),
	}
	DefaultSchemas = []string{
		"http://",
		"https://",
	}
	DefaultSchemasBytes = [][]byte{
		[]byte("http://"),
		[]byte("https://"),
	}
	ExtensionSchemas = []string{
		"chrome-extension://",
		"safari-extension://",
		"moz-extension://",
		"ms-browser-extension://",
	}
	FileSchemas = []string{
		"file://",
	}
	WebSocketSchemas = []string{
		"ws://",
		"wss://",
	}
)

func newCors(config Config) *cors {
	if err := config.Validate(); err != nil {
		panic(err.Error())
	}

	for _, origin := range config.AllowOrigins {
		if origin == "*" {
			config.AllowAllOrigins = true
		}
	}

	return &cors{
		allowOriginFunc:  config.AllowOriginFunc,
		allowAllOrigins:  config.AllowAllOrigins,
		allowCredentials: config.AllowCredentials,
		allowOrigins:     normalize(config.AllowOrigins),
		normalHeaders:    generateNormalHeaders(config),
		preflightHeaders: generatePreflightHeaders(config),
		wildcardOrigins:  config.parseWildcardRules(),
	}
}

func (cors *cors) applyCors(c *app.RequestContext) {
	origin := c.Request.Header.Get("Origin")
	if len(origin) == 0 {
		// request is not a CORS request
		return
	}
	host := c.Request.Host()

	o := str2bytes(origin)
	if bytes.HasPrefix(o, DefaultSchemasBytes[0]) {
		if compareByteSlices(o, DefaultSchemasBytes[0], host) == 0 {
			return
		}
	}
	if bytes.HasPrefix(o, DefaultSchemasBytes[1]) {
		if compareByteSlices(o, DefaultSchemasBytes[1], host) == 0 {
			return
		}
	}

	if !cors.validateOrigin(origin) {
		c.AbortWithStatus(consts.StatusForbidden)
		return
	}

	if bytes.Equal(c.Request.Method(), DefaultHeaderBytes[0]) {
		cors.handlePreflight(c)
		defer c.AbortWithStatus(consts.StatusNoContent) // Using 204 is better than 200 when the request status is OPTIONS
	} else {
		cors.handleNormal(c)
	}

	if !cors.allowAllOrigins {
		c.Header("Access-Control-Allow-Origin", origin)
	}
}

func (cors *cors) validateWildcardOrigin(origin string) bool {
	for _, w := range cors.wildcardOrigins {
		if w[0] == "*" && strings.HasSuffix(origin, w[1]) {
			return true
		}
		if w[1] == "*" && strings.HasPrefix(origin, w[0]) {
			return true
		}
		if strings.HasPrefix(origin, w[0]) && strings.HasSuffix(origin, w[1]) {
			return true
		}
	}

	return false
}

func (cors *cors) validateOrigin(origin string) bool {
	if cors.allowAllOrigins {
		return true
	}
	for _, value := range cors.allowOrigins {
		if value == origin {
			return true
		}
	}
	if len(cors.wildcardOrigins) > 0 && cors.validateWildcardOrigin(origin) {
		return true
	}
	if cors.allowOriginFunc != nil {
		return cors.allowOriginFunc(origin)
	}
	return false
}

func (cors *cors) handlePreflight(c *app.RequestContext) {
	for key, value := range cors.preflightHeaders {
		if len(value) > 0 {
			c.Response.Header.Set(key, value)
		}
	}
}

func (cors *cors) handleNormal(c *app.RequestContext) {
	for key, value := range cors.normalHeaders {
		if len(value) > 0 {
			c.Response.Header.Set(key, value)
		}
	}
}
