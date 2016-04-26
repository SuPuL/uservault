// Copyright © 2016 Wolf Bauer <mailsupul@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/supul/uservault/vendor/github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start uservault server",
	Long: `Start a uservault instance with the given config parameters. If no config is present the server
	       will start in a release envirnoment and can be found under localhost:80.

	       You can test a successful start with:

	       http://localhost/health
	       or
	       http://uservault.local/health

	       if your host file is set properly.`,
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Serve called")

		env, _ = cmd.Flags().GetString("env")

		log.Println("Set environment:", env)
		gin.SetMode(env)

		router := gin.Default()

		v1 := router.Group("/v1.0")

		v1.GET("/health", func(c *gin.Context) {
		        c.JSON(http.StatusOK, gin.H{
		                "message": "ok",
		        })
		})
		//
		//// User Handling
		//v1.GET("/users", func (c *gin.Context) {
		//        c.JSON(http.StatusForbidden, nil)
		//})
		//
		//v1.POST("/users", func (c *gin.Context) {
		//        c.JSON(http.StatusForbidden, nil)
		//})
		//
		//v1User := v1.Group("/users")
		//
		//v1User.GET("/:id", func(c *gin.Context) {
		//        c.JSON(http.StatusNotFound, gin.H{})
		//})

		var err error;

		//if (useTls) {
		//	err = router.RunTLS(addr, certFile, keyFile);
		//} else {
		//	err = router.Run(addr);
		//}

		log.Printf("About to listen on 10443. Go to https://127.0.0.1:10443/")
		log.Fatal(err)
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("env", "release", "Set the envirnoment for the command (e.g. release or dev).")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
