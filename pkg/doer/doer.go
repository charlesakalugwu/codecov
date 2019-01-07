package doer

//go:generate go get github.com/golang/mock/gomock
//go:generate go install github.com/golang/mock/mockgen
//go:generate mockgen -destination=../mocks/mock_$GOPACKAGE/doer.go -package=mock_$GOPACKAGE -source doer.go
//go:generate gofmt -s -l -w ../mocks/mock_$GOPACKAGE/doer.go
//go:generate goimports -local=github.com/charlesakalugwu/codecov -e -w ../mocks/mock_$GOPACKAGE/doer.go

type Doer interface {
	DoSomething(int, string) error
	SaySomething(string) error
	KickSomething(string) error
	SingSomething(string) error
}
