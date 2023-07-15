package tx

type TxOssStore struct {
	A string
	B map[string]string
}

func NewTxOssStore() *TxOssStore {
	return &TxOssStore{}
}

func (t *TxOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	return nil
}
