/*
 * @Author: zhanghao
 * @Date: 2018-11-20 15:28:57
 * @Last Modified by: zhanghao
 * @Last Modified time: 2018-11-21 11:15:18
 */

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 打印huffmanTree的top节点
	huffmanTree := createHTreeASICC("./testHuffman.txt")
	fmt.Println("huffmanTree top:", huffmanTree)

	// 打印该文件的huffman编码
	m := make(map[uint32]string)
	hTreeToMap(huffmanTree, "", m)
	fmt.Println("huffmanTree:", m)

	// 用01字符串表示该文件转码后的二进制值
	hc := toHuffmanCode(huffmanTree, "/Users/zhanghao/Documents/algorithmNCEPU/shiyan/go/testHuffman.txt")
	fmt.Println("huffman code:", hc)
}

func createHTreeASICC(filename string) *binaryTree {
	weights := countASICC(filename)
	return createHTree(weights)
}

func hTreeToMap(ht *binaryTree, code string, m map[uint32]string) {
	if ht.left == nil {
		m[ht.c] = code
		return
	}
	hTreeToMap(ht.left, code+"0", m)
	hTreeToMap(ht.right, code+"1", m)
}

func toHuffmanCode(ht *binaryTree, filename string) string {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	huffmanTree := createHTreeASICC(filename)
	m := make(map[uint32]string)
	hTreeToMap(huffmanTree, "", m)
	var huffmanCode string
	for i, _ := range fileBytes {
		huffmanCode = huffmanCode + m[uint32(fileBytes[i])]
	}
	return huffmanCode
}

func countASICC(filename string) (weights map[uint32]uint) {
	weights = make(map[uint32]uint)
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	for i, _ := range fileBytes {
		weights[uint32(fileBytes[i])]++
	}
	return weights
}

type binaryTree struct {
	c      uint32
	weight uint
	left   *binaryTree
	right  *binaryTree
}

// 创建huffman tree. index作为assic码(或者其他码)的值. weights[i]为权值(该码出现的次数)
func createHTree(weights map[uint32]uint) *binaryTree {
	// 创建森林 []*binaryTree
	forest := make([]*binaryTree, len(weights))
	i := 0
	for k, _ := range weights {
		forest[i] = &binaryTree{
			c:      k,
			weight: weights[k],
		}
		i++
	}

	minIndex := [2]uint{}
	minBT := [2]*binaryTree{}
	finish := false
	for !finish {
		j := 0
		for ; j < 2; j++ {
			for i, _ := range forest {
				if forest[i] != nil {
					if minBT[j] == nil {
						minBT[j] = forest[i]
						minIndex[j] = uint(i)
					} else if forest[i].weight < minBT[j].weight {
						minBT[j] = forest[i]
						minIndex[j] = uint(i)
					}
				}
			}
			if j == 0 {
				forest[minIndex[0]] = nil
			} else if minBT[1] == nil {
				finish = true
			} else {
				forest[minIndex[1]] = &binaryTree{
					weight: minBT[0].weight + minBT[1].weight,
					left:   minBT[0],
					right:  minBT[1],
				}
				minBT[0], minBT[1] = nil, nil
			}
		}
	}
	return minBT[0]
}
