package main

import (
	"fmt"
	"math"

	"github.com/tiennv1997/rng/internal"
)

func main() {
	cs := "83e27f682128eb1852b048203dfd6931"
	ss := "e8df2cc3b9ccb583ce5ea92336842387"
	nonce := 1942124

	cfg := internal.NewRNGConfig(cs, ss, int64(nonce))

	rng := internal.NewProvablyFairRNG(cfg)

	for i := 0; i < 10; i++ {
		fmt.Println(rng.NextByte())
	}

	fmt.Println(rng.NextFloat())

	bytes := internal.FloatToBytes(math.Pi)
	fmt.Println(bytes)
	float := internal.BytesToFloat(bytes)
	fmt.Println(float)

}
