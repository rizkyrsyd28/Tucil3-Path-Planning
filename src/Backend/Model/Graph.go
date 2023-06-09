package DataType

import (
	"encoding/json"
	"fmt"
)

type Link struct {
	To   *Node
	Dist float64
}

type Graph struct {
	Nodes []*Node
	Links map[*Node][]*Link
}

// READ FILE
func Json2Nodes(jsonData []byte, nodes *[]Node) ([][]int, int, int, string) {
	var parse struct {
		Nodes []*struct {
			ID   int        `json::"id"`
			Name string     `json:"name"`
			Coor [2]float64 `json:"coor"`
		} `json:"nodes"`
		AdjMat [][]int `json:"mat"`
		From   int     `json:"from"`
		To     int     `json:"to"`
		Algo   string  `json:"algo"`
	}

	err := json.Unmarshal(jsonData, &parse)
	if err != nil {
		panic(err)
	}

	*nodes = make([]Node, len(parse.Nodes))
	for i := 0; i < len(*nodes); i++ {
		(*nodes)[i] = Node{Id: parse.Nodes[i].ID, Name: parse.Nodes[i].Name, Coor: struct {
			X float64
			Y float64
		}{Y: parse.Nodes[i].Coor[0], X: parse.Nodes[i].Coor[1]}}
	}
	return parse.AdjMat, parse.From, parse.To, parse.Algo
}

func (g *Graph) CreateGraph(nodes []Node, mat [][]int) {
	g.Nodes = make([]*Node, len(nodes))
	for i := 0; i < len(nodes); i++ {
		g.Nodes[i] = &nodes[i]
	}

	g.Links = make(map[*Node][]*Link)

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				g.Links[g.Nodes[i]] = append(g.Links[g.Nodes[i]], &Link{To: g.Nodes[j], Dist: nodes[i].DistanceHaversine(nodes[j])})
			}
		}
	}
}

func (l Link) PrintLink() {
	l.To.PrintNode()
	fmt.Printf("Dist : %f\n", l.Dist)
}

func (g Graph) PrintGraph() {
	for i := 0; i < len(g.Nodes); i++ {
		g.Nodes[i].PrintNode()
	}

	for key, link := range g.Links {
		fmt.Print("Parent : ")
		key.PrintNode()
		for _, l := range link {
			l.PrintLink()
		}
	}
}
