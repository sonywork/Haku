/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co,Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the 'Software'), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2017/08/15     Tang Xiaoji
 */

package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"Haku/route"
	"Haku/general"
	"Haku/orm/cockroach"
	"Haku/middleware/jwt"
	"Haku/mongo"
	"Haku/task"
)

var (
	server  *echo.Echo
)

func startEchoServer() {
	server = echo.New()

	server.Use(middleware.Recover())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: configuration.corsHosts,
		AllowMethods: []string{echo.GET, echo.POST},
		AllowCredentials: true,
	}))

	server.Use(jwt.CustomJWT(configuration.tokenKey))

	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${uri} ${method} ${status} ${remote_ip} ${latency_human} ${bytes_in} ${bytes_out}\n",
	}))

	server.HTTPErrorHandler = general.EchoRestfulErrHandler
	server.Validator = general.NewEchoValidator()

	if configuration.isDebug {
		log.SetLevel(log.DEBUG)
	} else {
		log.SetLevel(log.INFO)
	}

	route.InitRoute(server, configuration.tokenKey)

	server.Start(configuration.address)
}

func init() {
	readConfiguration()
	cockroach.InitCockroachPool()
	mongo.InitHaku(configuration.mongoURL)

	task.StartJob()
}
