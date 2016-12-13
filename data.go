package main

const gitattribs = `*.zip filter=lfs diff=lfs merge=lfs -text
*.xlsm filter=lfs diff=lfs merge=lfs -text
*.kml filter=lfs diff=lfs merge=lfs -text
`

var (
	pathToGitIgnoreContents map[string]string
)

func init() {
	pathToGitIgnoreContents = make(map[string]string)

	pathToGitIgnoreContents[".gitignore"] = `
################################################################################
# This .gitignore file was automatically created by Microsoft(R) Visual Studio.
################################################################################

#Swyfft-specific stuff
*.orig
/swyf-tests/reports/*
/swyf-tests/*.log
/TestResults/
*.RData
*.RHistory
desktop.ini
.DS_Store

# Jenkins Stuff
**/[Oo]ut/*
JenkinsResults.trx

# User-specific files
*.user
*.suo
tmp*.tmp
*.userosscache
*.sln.docstates

# Build results
[Dd]ebug/
[Dd]ebugPublic/
[Rr]elease/
[Rr]eleases/
x64/
x86/
build/
bld/
[Bb]in/
[Oo]bj/
Swyfft.Web/js/**/*.min.js
Swyfft.Web/css/**/*.min.css
Swyfft.Web/css/Dist/

# Visual Studio 2015 cache/options directory
.vs/

# ReSharper
_ReSharper*/
*.[Rr]e[Ss]harper
*.DotSettings.user

# Ignore NuGet Packages
*.nupkg
# Ignore the packages folder
**/packages/*
# except build/, which is used as an MSBuild target.
!**/packages/build/
# except bootstrap 
!**/packages/bootstrap.3.3.4
# Uncomment if necessary however generally it will be regenerated when needed
#!**/packages/repositories.config
~*.xlsm
.metadata/*
.recommenders/*

# Ignore the npm package directory
**/node_modules/*

# Visual Studio profiler
*.psess
*.vsp
*.vspx

# Backup & report files from converting an old project file
# to a newer Visual Studio version. Backup files are not needed,
# because we have git ;-)
_UpgradeReport_Files/
Backup*/
UpgradeLog*.XML
UpgradeLog*.htm

# SQL Server files
*.mdf
*.ldf

# Business Intelligence projects
*.rdl.data
*.bim.layout
*.bim_*.settings

# Microsoft Fakes
FakesAssemblies/

# Node.js Tools for Visual Studio
.ntvs_analysis.dat

# Visual Studio 6 build log
*.plg

# Visual Studio 6 workspace options file
*.opt
lastMigrationSeedDate
Swyfft.Web/css/BuyerFwdStyles.css
Swyfft.Web/css/BuyerFwdStyles.css
Swyfft.Web/App/Dist
xunit-results-Swyfft.Common.Tests.xml
xunit-results-Swyfft.Seeding.Tests.xml
xunit-results-Swyfft.Services.AcceptanceTests.Critical.xml
xunit-results-Swyfft.Services.Tests.xml
xunit-results-Swyfft.Web.AcceptanceTests.Critical.xml
xunit-results-Swyfft.Web.Tests.xml
`

	// TODO nested contents
}
