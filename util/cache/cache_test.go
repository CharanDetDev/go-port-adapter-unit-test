package cache

import (
	"testing"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/config"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"
)

func TestCacheGet(t *testing.T) {
	tests := []struct {
		name    string
		want    bool
		wantErr string
		cache   MakeCache
	}{
		{
			name:    "GET Cache with KEY",
			want:    true,
			wantErr: "redis: nil",
			cache: MakeCache{
				Key:    "test-cache",
				Data:   "",
				Expire: "",
			},
		},
	}

	config.ConfigInitForTest()
	InitCache()
	var err error
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cache.Data, err = Get(tt.cache.Key)
			if err != nil {
				logg.Printlogger("\t\t ********** GET Redis cache ***********", "Error", err)
				tt.want = false
			}

			if tt.want {
				logg.Printlogger_JsonMarshalIndent("\t\t ********** GET Redis cache ***********", "GET Cache", tt.cache)
			}
		})
	}
}

func TestCacheSet(t *testing.T) {
	data := MakeCache{
		Key:    "test-cache",
		Data:   "set",
		Expire: "1m",
	}
	tests := []struct {
		name    string
		want    bool
		wantErr string
		cache   MakeCache
	}{
		{
			name:    "SET Cache",
			want:    true,
			wantErr: "redis: nil",
			cache: MakeCache{
				Key:    "test-cache",
				Data:   data,
				Expire: "1h",
			},
		},
	}

	config.ConfigInitForTest()
	InitCache()
	var err error
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err = Set(tt.cache)
			if err != nil {
				logg.Printlogger("\t\t ********** SET Redis cache ***********", "Error", err)
				tt.want = false
			}

			if tt.want {
				logg.Printlogger_JsonMarshalIndent("\t\t ********** SET Redis cache **********", "SET Cache", tt.cache)
			}
		})
	}
}

func TestCacheUpdate(t *testing.T) {
	data := MakeCache{
		Key:    "test-cache",
		Data:   "update",
		Expire: "1m",
	}
	tests := []struct {
		name    string
		want    bool
		wantErr string
		cache   MakeCache
	}{
		{
			name:    "UPDATE Cache",
			want:    true,
			wantErr: "redis: nil",
			cache: MakeCache{
				Key:    "test-cache",
				Data:   data,
				Expire: "1h",
			},
		},
	}

	config.ConfigInitForTest()
	InitCache()
	var err error
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err = Update(tt.cache)
			if err != nil {
				logg.Printlogger("\t\t ********** UPDATE Redis cache ***********", "Error", err)
				tt.want = false
			}

			if tt.want {
				logg.Printlogger_JsonMarshalIndent("\t\t ********** UPDATE Redis cache **********", "UPDATE Cache", tt.cache)
			}
		})
	}
}

func TestCacheDelete(t *testing.T) {
	tests := []struct {
		name    string
		want    bool
		wantErr string
		cache   MakeCache
	}{
		{
			name:    "DELETE Cache",
			want:    true,
			wantErr: "redis: nil",
			cache: MakeCache{
				Key:    "test-cache",
				Data:   "",
				Expire: "",
			},
		},
	}

	config.ConfigInitForTest()
	InitCache()
	var err error
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err = Delete(tt.cache.Key)
			if err != nil {
				logg.Printlogger("\t\t ********** DELETE Redis cache ***********", "Error", err)
				tt.want = false
			}

			if tt.want {
				logg.Printlogger_JsonMarshalIndent("\t\t ********** DELETE Redis cache **********", "DELETE Cache", tt.cache)
			}
		})
	}
}
