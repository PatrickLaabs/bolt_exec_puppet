# golang binary for bolt run_command


ToDo:

[x] exec.Command for exeuting puppet
[ ] exec.cmd for grep
[x] return exit code output from stdout
[x] integrate fmt.print for start
[ ] stderrpipe muss ausgegeben werden
[ ] save exit code into var e, and return it. Use switch statement for advanced handling

puppet apply --noop --test --debug manifest/manifest.pp | grep exit

puppet apply --noop --test --debug manifest/manifest.pp | grep -E "status code 1"


-- 
Exit Codes
0 & 2 als 0 zurückgeben (success)
1 & 4 & 6 als 1 zurückgeben (err)