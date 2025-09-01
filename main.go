package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("must've atleast two args")
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	var testCase map[string]interface{}
	if err := json.Unmarshal(data, &testCase); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	keysData := testCase["keys"].(map[string]interface{})
	k := int(keysData["k"].(float64))

	var points []Point
	for key, value := range testCase {
		if key == "keys" {
			continue
		}

		rootData := value.(map[string]interface{})
		baseStr := rootData["base"].(string)
		valueStr := rootData["value"].(string)

		base, _ := strconv.Atoi(baseStr)
		x, _ := strconv.Atoi(key)
		y, err := convertFromBase(valueStr, base)
		if err != nil {
			continue
		}

		points = append(points, Point{big.NewInt(int64(x)), y})
	}

	if len(points) < k {
		fmt.Println("insufficient points")
		os.Exit(1)
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].X.Cmp(points[j].X) < 0
	})

	secret := lagrangeInterpolation(points[:k])
	fmt.Println(secret.String())
}
