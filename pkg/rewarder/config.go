package rewarder

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// Config struct
type Config struct {
	Name         string           `json:"name"`
	Round        string           `json:"round"`
	StartBlock   uint64           `json:"startBlock"`
	EndBlock     uint64           `json:"endBlock"`
	CubeNames    []string         `json:"cubes"`
	Pool         PoolConfig       `json:"pool"`
	Pools        []PoolConfig     `json:"pools"`
	RewardMap    RewardMap        `json:"rewards"`
	RewardAmount decimal.Decimal  `json:"reward_amount"`
	MaxGasUsed   decimal.Decimal  `json:"max_gas_used"`
	Nfts         []common.Address `json:"nfts"`
	NftBoost     float64          `json:"nft_boost"`
	NftMaxBoost  float64          `json:"nft_max_boost"`

	rewardDir      string
	roundDir       string
	startTimestamp uint64
	endTimestamp   uint64
	blocks         uint64
	cubeFinders    CubeFinders
	poolPrices     PoolPrices
}

// NewConfig new config
func NewConfig(filepath string) (*Config, error) {
	if _, err := os.Stat(filepath); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := jsonex.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// LogToFile log to file
func (c *Config) LogToFile() error {
	round := strings.Split(c.Round, "/")
	logFilepath := path.Join("logs", fmt.Sprintf("%s_%s_%d.log", c.Name, round[0], time.Now().Unix()))
	logFile, err := os.OpenFile(logFilepath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	return nil
}

// RewardDir get reward dir
func (c *Config) RewardDir() string {
	if len(c.rewardDir) == 0 {
		c.rewardDir = path.Join("rewards", c.Name)
	}

	return c.rewardDir
}

// RoundDir get round dir
func (c *Config) RoundDir() string {
	if len(c.roundDir) == 0 {
		c.roundDir = path.Join(c.RewardDir(), c.Round)
	}

	return c.roundDir
}

// MakeRoundDir make round dir
func (c *Config) MakeRoundDir() error {
	log.Printf("making round dir: ./%s/", c.RoundDir())

	if err := os.MkdirAll(c.RoundDir(), os.ModePerm); err != nil {
		return err
	}

	return nil
}

// LoadCubeFinders load cube finders
func (c *Config) LoadCubeFinders() error {
	log.Printf("loading cube finders: %s", strings.Join(c.CubeNames, ","))

	findCubes, err := GetCubeFinders(c.CubeNames)
	if err != nil {
		return err
	}

	c.cubeFinders = findCubes

	return nil
}
