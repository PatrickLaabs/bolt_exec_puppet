# Execute commands

Golang project to execute commands inside the golang code/binary and handle
the err, exit codes and returning the exit code from the executed command back to 
the shell.

In this way, we can handle the exit code returned by the command.

The exit code is handled inside the go code with a switch statement.

----

### ToDos:

* [ ] Maybe use a combinedOutput instead of Stdout and Stderr. 
* [ ] Running tests for checking correct exitCode handling.
* [ ] Only one binary for running puppet statement
  > running binary without --args runs puppet noop statement, 
  > set --args (--no-noop) to run operational puppet run.
* [ ] set --tags for puppet statement
* [ ] HashMap for holding different arguments calling puppet statement
  * --skip_tags; --tags (& contains); --no-noop 
