$ErrorActionPreference = 'SilentlyContinue'
$installDir = "C:\Program Files\bolt_exec_puppet"

# Deinstallation bolt_exec_puppet
If (Test-Path $installDir)
    {
    "Entferne bolt_exec_puppet..."
    Remove-Item "$installDir" -Recurse -Force
    "Fertig!"
    }
Else
    {
    "Keine bolt_exec_puppet Installation gefunden. Deinstallation wird beendet."
    }

# Remove bolt_exec_puppet from PATH
$Remove = 'C:\Program Files\bolt_exec_puppet'
$env:Path = ($env:Path.Split(';') | Where-Object -FilterScript {$_ -ne $Remove}) -join ';'

# [Environment]::SetEnvironmentVariable("PATH", $null, [EnvironmentVariableTarget]::Machine)