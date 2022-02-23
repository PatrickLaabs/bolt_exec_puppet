# Execute commands

Golang Project to execute commands inside the golang code/binary and handle
the errors, exit codes and returning the exit code from the executed command back to 
the shell.

In this way, we can handle the exit code returned by the command.

The exit code is handled inside the go code with a switch statement.

----

### ToDos:

* [ ] Project renaming to 'bolt_exec_puppet'
* [x] Update README for usage instructions of this binary
* [ ] Prepare Build-Pipeline on Jenkins
* [ ] Running tests on linux systems
* [ ] Running tests on Windows systems
  * [ ] Check if 'if'-Statement works as intended on Windows
  * [ ] Check for Exit Code handling on Windows systems
* [ ] Print detailed Exit code before handling func exitHandle switch statement.
  (Using a Flag, printing the exit code without handling)
* [x] **FIX:** running _only_ the binary does not result in an error. Print help menu instead.
* [x] **FIX:** running `./binary <flag>` returns the command into the exec.Command interface. Return an error instead.
* [x] **FIX:** improve `./binary --help` menu, to show a more specific way how to use this binary.
* [x] **FIX:** running `./binary agent --test --tags=<module>`  without `--noop` results in an error, because we're returning an empty string.
----
### Usage Instructions:

General usage:

`bolt_exec_puppet agent [--test] [--noop] [--tags TAGS] [--skip_tags SKIP_TAGS]`

Some examples:

  `./bolt_exec_puppet agent --test`

  `./bolt_exec_puppet agent --test --noop`

  `./bolt_exec_puppet agent --test --noop --tags=<module>`

  `./bolt_exec_puppet agent --test --noop --tags <module>`

  `./bolt_exec_puppet agent --test --noop --skip_tags=<module>`

  `./bolt_exec_puppet agent --test --noop --skip_tags <module>`

A combination of both --tags and --skip_tags is also possible:

  `./bolt_exec_puppet agent --test --noop --tags=<module> --skip_tags=<module>`

Every possibility can be run without --noop