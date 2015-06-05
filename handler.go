// Package of functions for the technotes page
package technotes

import (
	"fmt"
	"net/http"
	//"math/rand"
	"time"
)

// rand.Seed(time.Now().Unix())

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Handler - code")
	fmt.Fprintln(w, "<html>")
	fmt.Fprintln(w, "<head>")
	fmt.Fprintln(w, "<link rel='stylesheet' type='text/css' href='technotes.css' />")
	fmt.Fprintln(w, "<link href='http://fonts.googleapis.com/css?family=Special+Elite' rel='stylesheet' type='text/css'>")
	fmt.Fprintln(w, "<link href='http://fonts.googleapis.com/css?family=EB+Garamond' rel='stylesheet' type='text/css'>")
	fmt.Fprintln(w, "<link href='http://fonts.googleapis.com/css?family=Ubuntu+Mono:400,700,700italic' rel='stylesheet' type='text/css'>")
	fmt.Fprintln(w, time.Now().Unix())

	fmt.Fprintln(w, "</head>")
	fmt.Fprintln(w, "<body>")

	// hoizontal rule
	fmt.Fprintln(w, "<hr>")

	// h2
	fmt.Fprintln(w, "<h2>Connect or run remote commands with a bash alias</h2>")

	// paragraph
	fmt.Fprintln(w, "<p>")
	fmt.Fprintln(w, "You can issue remote commands on a server or connect to a server with a simple alias.  This will serve as a shortcut to connect to the remote host or just to execute a command.  This is particularly elegant when you pair it with pre shared keys, which can be setup with the <code>ssh-copy-id</code> command that comes with Ubuntu.")
	fmt.Fprintln(w, "</p>")

	// pre
	fmt.Fprintln(w, "<pre>")
	fmt.Fprintln(w, "<code>## create the alias</code>")
	fmt.Fprintln(w, "alias booger='ssh jdoe@booger.com'")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "<code>## will connect you to remote host.</code>")
	fmt.Fprintln(w, "booger")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "<code>## will list the contents of the remote home dir.</code>")
	fmt.Fprintln(w, "booger ls")
	fmt.Fprintln(w, "</pre>")

	// hoizontal rule
	fmt.Fprintln(w, "<hr>")

	fmt.Fprintln(w, "<h2>Setup vim and set it as the default editor on Ubuntu.</h2>")

	// code
	fmt.Fprintln(w, "<pre>")
	fmt.Fprintln(w, "sudo apt-get install vim-nox")
	fmt.Fprintln(w, "sudo apt-get install vim-nox")
	fmt.Fprintln(w, "sudo update-alternatives -config editor")
	fmt.Fprintln(w, "</pre>")

	// hoizontal rule
	fmt.Fprintln(w, "<hr>")

	// table
	fmt.Fprintln(w, "<table>")
	fmt.Fprintln(w, "This is an example of some sample text in a table")
	fmt.Fprintln(w, "</table>")

	fmt.Fprintln(w, "</body>")
	fmt.Fprintln(w, "</html>")
}
