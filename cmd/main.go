/*
Copyright 2024 Hurricane1988 Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hurricane1988/harbor-webhook/middleware"
	"github.com/hurricane1988/harbor-webhook/pkg/utils"
	"github.com/hurricane1988/harbor-webhook/pkg/version"
	"github.com/hurricane1988/harbor-webhook/routes"
	ctrl "sigs.k8s.io/controller-runtime"
	"strings"
)

var (
	setupLog = ctrl.Log.WithName("setup")
)

func main() {
	// 打印logo
	fmt.Println(utils.Term())
	// 打印版本信息
	fmt.Println(version.String())
	// 初始化gin对象
	router := gin.New()
	// 启用跨域中间件
	router.Use(middleware.Cors)
	routes.Router(router)
	// gin程序启动
	err := router.Run(strings.Join([]string{"0.0.0.0", "8080"}, ":"))
	if err != nil {
		setupLog.Error(err, "startup harbor-webhook failed.")
		return
	}
}
