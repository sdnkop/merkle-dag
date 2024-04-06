package merkledag

import (
	"hash"
)

func Add(store KVStore, node Node, h hash.Hash) ([]byte, error) {
	if store == nil || node == nil || h == nil {
		return nil, nil
	}

	// 递归保存节点数据到 KVStore
	err := saveNodeData(store, node, h)
	if err != nil {
		return nil, err
	}

	// 获取根节点的哈希值作为 Merkle Root
	rootHash, err := store.Get([]byte("root"))
	if err != nil {
		return nil, err
	}

	return rootHash, nil
}

func saveNodeData(store KVStore, node Node, h hash.Hash) error {
	if node.Type() == FILE {
		file, ok := node.(File)
		if !ok {
			return nil // 或者返回错误，表示节点类型不符合预期
		}

		// 获取文件内容并计算哈希值
		fileContent := file.Bytes()
		h.Write(fileContent)
		hashValue := h.Sum(nil)

		// 将哈希值作为文件内容保存到 KVStore
		err := store.Put(hashValue, fileContent)
		if err != nil {
			return err
		}

		return nil
	} else if node.Type() == DIR {
		dir, ok := node.(Dir)
		if !ok {
			return nil // 或者返回错误，表示节点类型不符合预期
		}

		// 获取文件夹迭代器，遍历子节点并递归保存节点数据
		iterator := dir.It()
		for iterator.Next() {
			child := iterator.Node()
			err := saveNodeData(store, child, h)
			if err != nil {
				return err
			}
		}

		return nil
	}

	return nil // 或者返回错误，表示不支持的节点类型
}
