package oss

// OSS oss interface
type OSS interface {
	Put(key, localfile string) error
	Del(key string) error
}
