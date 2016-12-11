package oss

// AwsOSS aws s3 oss
type AwsOSS struct {
}

// Put put
func (oss *AwsOSS) Put(key, localfile string) error {
	return nil
}

// Del del
func (oss *AwsOSS) Del(key string) error {
	return nil
}
