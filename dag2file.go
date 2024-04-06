package merkledag

import (
	"io/ioutil"
	"os"
)

// Hash2File 从 KVStore 中读取哈希对应数据，然后根据 path 返回文件内容
func Hash2File(store KVStore, hash []byte, path string, hp HashPool) ([]byte, error) {
	// 根据提供的哈希值从 KVStore 中获取数据
	data, err := store.Get(hash)
	if err != nil {
		return nil, err
	}

	// 将获取的数据写入指定的文件中
	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return nil, err
	}

	// 读取文件内容，作为返回值
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// 删除临时文件
	err = os.Remove(path)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}
