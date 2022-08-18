package stl

import (
	"strconv"
	"strings"
)

type Codec struct {
	serializeString string
	deserializeList []string
	index           int
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.serializeString = ""
	return this.serializeDFS(root)
}

func (this *Codec) serializeDFS(root *TreeNode) string {
	if root == nil {
		this.serializeString += "nil,"
		return this.serializeString
	}
	this.serializeString += strconv.Itoa(root.Val) + ","
	this.serialize(root.Left)
	this.serialize(root.Right)
	return this.serializeString
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.deserializeList = strings.Split(data, ",")
	this.index = 0
	return this.deserializeDFS()
}

func (this *Codec) deserializeDFS() *TreeNode {
	if this.deserializeList[this.index] == "nil" {
		this.index++
		return nil
	}
	val, _ := strconv.Atoi(this.deserializeList[this.index])
	this.index++
	return &TreeNode{
		Val:   val,
		Left:  this.deserializeDFS(),
		Right: this.deserializeDFS(),
	}
}

