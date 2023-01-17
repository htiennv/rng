package internal

import "fmt"

type RNGConfig struct {
	clientSeed string
	serverSeed string
	nonce      int64
}

func NewRNGConfig(clientSeed, serverSeed string, nonce int64) *RNGConfig {
	return &RNGConfig{
		clientSeed: clientSeed,
		serverSeed: serverSeed,
		nonce:      nonce,
	}
}

type ProvablyFairRNG struct {
	config *RNGConfig

	currentRound       int64
	currentRoundCursor uint64
	currentRoundMac    []byte
}

func NewProvablyFairRNG(config *RNGConfig) *ProvablyFairRNG {
	return &ProvablyFairRNG{
		config:             config,
		currentRound:       0,
		currentRoundCursor: 0,
		currentRoundMac:    nil,
	}
}

func (p *ProvablyFairRNG) updateCurrentRoundBuffer() {
	key := []byte(p.config.serverSeed)

	input := fmt.Sprintf("%s:%d:%d", p.config.clientSeed, p.config.nonce, p.currentRound)

	mac := HmacSha256(key, []byte(input))

	p.currentRoundMac = mac
}

func (p *ProvablyFairRNG) NextByte() byte {
	if p.currentRoundMac == nil {
		p.updateCurrentRoundBuffer()
		return p.NextByte()
	}

	buf := p.currentRoundMac

	result := buf[p.currentRoundCursor]

	if p.currentRoundCursor == 31 {
		p.currentRoundCursor = 0
		p.currentRound += 1
		p.currentRoundMac = nil
	} else {
		p.currentRoundCursor += 1
	}

	return result
}

func (p *ProvablyFairRNG) NextFloat() float64 {
	bytesPerFloat := 8

	bytes := make([]byte, bytesPerFloat)

	for i := 0; i < bytesPerFloat; i++ {
		b := p.NextByte()
		bytes[i] = b
	}
	fmt.Println(bytes)
	return BytesToFloat(bytes)
}
