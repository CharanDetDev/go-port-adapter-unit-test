package main

import "testing"

func TestHelloWorld(t *testing.T) { //! Function TestHelloWorld จะถูกสร้างขึ้นมาอัตโนมัติ

	//* Arrage
	type args struct {
		hello string //! argument ที่ต้องการส่งไปทดสอบ
	}
	tests := []struct {
		name string //! ชื่อของ test case
		args args   //! explicit field ของ argument
		want string //! expected
	}{
		// TODO: Add test cases.
		{
			name: "Case_Hello_World",
			args: args{
				hello: "Hello",
			},
			want: "World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//* Action
			got := HelloWorld(tt.args.hello)

			//* Assert
			if got != tt.want {
				t.Errorf("HelloWorld() = %v, want %v", got, tt.want)
			}
		})
	}
}
