package main

import (
	"bytes"
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/gin-gonic/gin"
	"sort"
	"strconv"
	"strings"
)

type compareParams struct {
	Lista string `json:"lista"`
	Listb string `json:"listb"`
}

type Result struct {
	OnlyInA      []string
	OnlyInB      []string
	Intersection []string
}

func main() {
	gin.SetMode(gin.DebugMode) //todo, prod mode
	router := gin.Default()
	// ping
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/comma-to-newline", func(c *gin.Context) {
		foo := c.Query("foo")
		bar := strings.ReplaceAll(foo, ",", "\n")
		c.JSON(200, bar)
	})
	router.GET("/newline-to-comma", func(c *gin.Context) {
		bar := c.Query("bar")
		addQuotation, _ := strconv.ParseBool(c.Query("addsinglequotationmarks"))
		arr := strings.Split(bar, "\n")

		var foo bytes.Buffer
		for _, v := range arr {
			if addQuotation {
				foo.WriteString("'")
			}
			foo.WriteString(strings.TrimSpace(v))
			if addQuotation {
				foo.WriteString("'")
			}
			foo.WriteString(",")
		}
		out := foo.String()
		c.JSON(200, out[:len(out)-1])
	})
	router.POST("/compare", func(c *gin.Context) {
		params := &compareParams{}
		_ = c.BindJSON(params)
		fmt.Println("list a: " + params.Lista)
		arrA := strings.Split(params.Lista, "\n")
		arrB := strings.Split(params.Listb, "\n")

		setA := hashset.New()
		for _, v := range arrA {
			if v != "" {
				setA.Add(strings.TrimSpace(v))
			}
		}
		setB := hashset.New()
		for _, v := range arrB {
			if v != "" {
				setB.Add(strings.TrimSpace(v))
			}
		}

		var resultOnlyInA []string
		var resultOnlyInB []string
		var resultIntersection []string
		for _, v := range arrA {
			if setB.Contains(v) {
				resultIntersection = append(resultIntersection, v)
				setB.Remove(v)
				fmt.Println("string exist in list b: " + v)
			} else {
				resultOnlyInA = append(resultOnlyInA, v)
				fmt.Println("string not exist in list b: " + v)
			}
		}
		for _, v := range setB.Values() {
			resultOnlyInB = append(resultOnlyInB, v.(string))
			fmt.Println("add remaining to result b " + v.(string))
		}

		sort.Strings(resultOnlyInA)
		sort.Strings(resultOnlyInB)
		sort.Strings(resultIntersection)

		result := Result{
			OnlyInA:      resultOnlyInA,
			OnlyInB:      resultOnlyInB,
			Intersection: resultIntersection,
		}
		c.JSON(200, result)
	})
	_ = router.Run(":8080")
}
