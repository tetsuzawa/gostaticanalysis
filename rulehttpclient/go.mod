module github.com/tetsuzawa/gostaticanalysis/rulehttpclient

go 1.15

require (
	github.com/gostaticanalysis/analysisutil v0.1.0
	github.com/tetsuzawa/gostaticanalysis/analysisutil/milestonequeue v0.0.0-20200903162001-9ef594e91a1e
	github.com/tetsuzawa/gostaticanalysis/analysisutil/usedtype v0.0.0-20200903161404-f88bb3a9763f
	golang.org/x/tools v0.0.0-20200903153655-76a6aac657c7
)

replace github.com/tetsuzawa/gostaticanalysis/analysisutil/milestonequeue => /Users/tetsu/repo/mercari/playground/github.com/tetsuzawa/gostaticanalysis/analysisutil/milestonequeue
replace github.com/tetsuzawa/gostaticanalysis/analysisutil/usettype => /Users/tetsu/repo/mercari/playground/github.com/tetsuzawa/gostaticanalysis/analysisutil/usedtype
