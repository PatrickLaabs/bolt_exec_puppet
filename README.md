# golang binary for bolt run_command


ToDo:

- exec.Command for exeuting puppet
- exec.cmd for grep
- grep for exit codes 0 & 2 (maybe better switch for every exit code and print them to stdout),
- return exit code output from stdout
- err handling (if puppet run is not executed)
- integrate fmt.print for start

puppet apply --noop --test --debug manifest/manifest.pp | grep exit

puppet apply --noop --test --debug manifest/manifest.pp | grep -E "status code 1"


-- 
Exit Codes
0 & 2 als 0 zurückgeben (success)
1 & 4 & 6 als 1 zurückgeben (err)