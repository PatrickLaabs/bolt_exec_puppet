$ErrorActionPreference = 'Stop'; # stop on all errors
$toolsDir   = "$(Split-Path -parent $MyInvocation.MyCommand.Definition)"
$installDir = "C:\Program Files\bolt_exec_puppet"
$zipfile = "$toolsDir\bolt_exec_puppet.zip"

# bolt_exec_puppet Installation Options
$packageArgs = @{
    FileFullPath = $zipfile
    Destination  = $installDir
}

# Wenn Installation schon vorhanden, dann vorher alte Installation löschen (Upgrade)
If (Test-Path -Path $installDir) {    
    
    & $toolsDir\chocolateyuninstall.ps1

    # Entpacke diffutils im Zielverzeichnis
    Get-ChocolateyUnzip @packageArgs
}
Else {
    Get-ChocolateyUnzip @packageArgs
}

# Adding executable into Path
if ($env:Path -like '*bolt_exec_puppet*')
{
    # Write-Host 'Path already added'
}
elseif ( -not ($env:Path -like '*bolt_exec_puppet*'))
{
    [Environment]::SetEnvironmentVariable("PATH", $Env:PATH + ";C:\Program Files\bolt_exec_puppet\bolt_exec_puppet.exe", [EnvironmentVariableTarget]::Machine)
}