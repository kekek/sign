package sign

import (
	"fmt"
	"testing"
)

func TestSignature_SignURL(t *testing.T) {
	type fields struct {
		Secret       string
		SignParamKey string
		SignKeyList  []string
		SignFunc     func(string, string) string
	}
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "case001", fields: fields{}, args: args{data: "http://sr-report-test.zljgp.com/pdf/1681463647835_13724_LoadRunner_Winsocket协议知识总结V1[1].1(修正版).pdf?id=23657228271226880"}},
		{name: "case002", fields: fields{}, args: args{data: "http://sr-report-test.zljgp.com/pdf/1681463647835_13724_LoadRunner_Winsocket协议知识总结V1[1].1(修正版).pdf?id=23657228271226880&age=18&name=ll"}},
		{name: "case003", fields: fields{}, args: args{data: "https://example.com/test?id=1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := defaultSignature
			got, err := s.SignURL(tt.args.data)
			fmt.Println("SignURL", got, err)

			val, err := s.VerifyURL(got)
			fmt.Println("VerifyURL : ", val, err)

			expired, err := s.Expired(got, 1)
			fmt.Println("Expired : ", expired, err)
		})
	}
}

func BenchmarkSignURL(b *testing.B) {

	s := defaultSignature

	data := "http://sr-report-test.zljgp.com/pdf/1681463647835_13724_LoadRunner_Winsocket协议知识总结V1[1].1(修正版).pdf?id=23657228271226880&age=18&name=ll"

	for n := 0; n < b.N; n++ {
		s.SignURL(data) // run fib(30) b.N times
	}
}
