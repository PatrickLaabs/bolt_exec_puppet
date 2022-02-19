# Execute commands

Golang Project to execute commands inside the golang code/binary and handle
the errors, exit codes and returning the exit code from the executed command back to 
the shell.

In this way, we can handle the exit code returned by the command.

The exit code is handled inside the go code with a switch statement.

----

### ToDos:

* [ ] Project renaming to 'puppet_bolt_exec'
* [ ] Update README for usage instructions of this binary
* [ ] Prepare Build-Pipeline on Jenkins
* [ ] Running tests on linux systems
* [ ] Running tests on Windows systems
  * [ ] Check if 'if'-Statement works as intended on Windows
  * [ ] Check for Exit Code handling on Windows systems

----
### Usage Instructions:

```/bolt_puppet_exec noop```

```./bolt_exec op\```

```./bolt_exec op```

```./bolt_exec help```

```./bolt_exec tags -add=<module> -start=--noop```

```./bolt_exec skip -add=<module> -start=--noop```