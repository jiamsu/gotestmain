package main

import  "fmt"
import  "crypto/sha256"
import  "time"
import  "bytes"
// import  "Buffer/binary"
import  "log"


    // 1.定义结构
type Block struct{
	// 1.1版本号
	Version uint64
	// 1.2前区块哈希
	PrevHash  []byte
	// 1.3梅克尔根
	MerkelRoot []byte
	// 1.4 时间戳
	TimeStamp uint64
	// 1.5难度值
	Difficulty uint64
	// 1.6随机数，也就是挖矿要找的数据
	Nonce uint64
	// 当前区块哈希 ，正常比特币区块中没有当前区块哈希
	Hash  []byte
	// 1.3数据
	Data  []byte
}

// 实现一个辅助函数，功能是将uint64转成[]byte

func Uint64ToByte(num uint64)[]byte  {
	var buffer bytes.Buffer
	err:=binary.Write(&buffer,binary.BigEndian,num)
	if err!=nil{
		log.Panic(err)

	}
	return buffer.Bytes()
	
	
}


// 2.创建区块
func NewBlock(data string ,prevBlockHash []byte) *Block{

	block :=Block{
		Version :00,
		PrevHash:prevBlockHash,
		MerkelRoot:[]byte{},
		TimeStamp:uint64(time.Now().Unix()),
		Difficulty:0,
		Nonce:0,


		Hash:[]byte{},//先填空，后面再计算//TODO
		Data:[]byte(data),
	}

	return &block
}
// 3.生成哈希
func (block *Block)  SetHash(){
	var blockInfo []byte

	//todo
	// 1.拼装数据
	// blockInfo = append(blockInfo,Uint64ToByte(block.Version)...)

	// blockInfo = append(blockInfo,block.PrevHash...)
	// blockInfo = append(blockInfo,block.MerkelRoot...)
	// blockInfo = append(blockInfo,Uint64ToByte(block.TimeStamp)...)
	// blockInfo = append(blockInfo,Uint64ToByte(block.Difficulty)...)
	// blockInfo = append(blockInfo,Uint64ToByte(block.Nonce)...)
	// blockInfo = append(blockInfo,block.Data ...)

	tmp :=[][]byte{
		Uint64ToByte(block.Version),
		PrevHash,
		block.MerkelRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,

	}

	blockInfo:=bytes.Join(tmp,[]byte{})

	// 2.sha256
hash:=	sha256.Sum256(blockInfo)
block.Hash=hash[:]
	
}

// 4.引入区块链
type BlockChain struct{
	// 定义一个区块链数组
	blocks []*Block


}


// 5.定义一个区块链 

func NewBlockChain() *BlockChain{
	// 创建一个创世块，并添加到区块链中
	genesisBlock:= GenesisBlock()
	return &BlockChain{
		blocks:[]*Block{genesisBlock},
	}
}

// 创世块,hello
func GenesisBlock() *Block{
	return	NewBlock("go一期创世块，老牛拜了！",[]byte{})
}


// 5。添加区块
func (bc *BlockChain) AddBlock(data string)  {
	// 获取前区块哈希


	lastBlock:=bc.blocks[len(bc.blocks)-1]

	prevHash:=lastBlock.Hash
// a.创建新区块
block:=NewBlock(data,prevHash)

// b.添加到区块链数组中
bc.blocks=append(bc.blocks,block)




	
}

// 6.重构代码





func main() {

	
	bc:=NewBlockChain()
	bc.AddBlock("班长向班花转了50枚比特币")
	bc.AddBlock("班长又向班花转了50枚比特币")

	for i, block:= range bc.blocks {
		fmt.Printf("===当前区块高度：%d===\n", i)

	fmt.Printf("前区块哈希：%x\n", block.PrevHash)
	fmt.Printf("当前区块哈希：%x\n", block.Hash)
	fmt.Printf("区块数据：%s\n", block.Data)		
	}	
}