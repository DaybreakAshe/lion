// @program: superlion
// @author: yanjl
// @description: superlion
// @create: 2024-08-15 15:06
package util

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func InitSnowflake() {
	node, _ = snowflake.NewNode(1)

}

func GenerateID() string {

	id := node.Generate().String()
	return id
}
