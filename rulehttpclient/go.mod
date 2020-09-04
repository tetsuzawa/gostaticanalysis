module github.com/tetsuzawa/gostaticanalysis/rulehttpclient

go 1.15

require (
	github.com/gostaticanalysis/analysisutil v0.1.0
	github.com/tetsuzawa/gostaticanalysis/analysisutil/milestonequeue v0.0.0-20200904010851-73db240b16fb
	github.com/tetsuzawa/gostaticanalysis/analysisutil/usedtype v0.0.0-20200904031439-0c30ebcf7fb7
	golang.org/x/tools v0.0.0-20200903185744-af4cc2cd812e
)

replace github.com/tetsuzawa/gostaticanalysis/analysisutil/milestonequeue => /Users/tetsu/repo/mercari/playground/github.com/tetsuzawa/gostaticanalysis/analysisutil/milestonequeue
replace github.com/tetsuzawa/gostaticanalysis/analysisutil/usedtype => /Users/tetsu/repo/mercari/playground/github.com/tetsuzawa/gostaticanalysis/analysisutil/usedtype
