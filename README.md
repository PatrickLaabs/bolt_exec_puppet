# golang binary for bolt run_command


ToDo:

- exec.Command for exeuting puppet
- exec.cmd for grep
- grep for exit codes 0 & 2 (maybe better switch for every exit code and print them to stdout),
- return exit code output from stdout

puppet apply --noop --test --debug manifest/manifest.pp | grep exit

puppet apply --noop --test --debug manifest/manifest.pp | grep -E "status code 1"
