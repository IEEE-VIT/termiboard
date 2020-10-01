package main

import "testing"

func Test_getReadableSize(t *testing.T) {
	type args struct {
		sizeInBytes uint64
	}
	tests := []struct {
		name                   string
		args                   args
		wantReadableSizeString string
	}{
		{name: "B", args: args{123}, wantReadableSizeString: "123.00 B"},
		{name: "KB", args: args{806092}, wantReadableSizeString: "787.20 KB"},
		{name: "MB", args: args{1519616}, wantReadableSizeString: "1.45 MB"},
		{name: "GB", args: args{4610440364}, wantReadableSizeString: "4.29 GB"},
		{name: "TB", args: args{1234567890123}, wantReadableSizeString: "1.12 TB"},
		{name: "PB", args: args{9876543210987654}, wantReadableSizeString: "8.77 PB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotReadableSizeString := getReadableSize(tt.args.sizeInBytes); gotReadableSizeString != tt.wantReadableSizeString {
				t.Errorf("getReadableSize() = %v, want %v", gotReadableSizeString, tt.wantReadableSizeString)
			}
		})
	}
}
