package main

import (
	"encoding/json"
	"net/http"
)

type PairRequest struct {
	Array  []int `json:"array"`
	Target int   `json:"target"`
}

type PairResponse struct {
	Solution [][]int `json:"solution"`
}

func main() {
	http.HandleFunc("/findPair", findPairHandler)
	http.ListenAndServe(":8080",nil)
}

func findPairHandler(w http.ResponseWriter, r *http.Request) {
	var req PairRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w,"Invalid Request", http.StatusBadRequest)
		return
	}

	result := findPairHandlerTarget(req.Array, req.Target)

	response := PairResponse{Solution: result}
	
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(response)
}

func findPairHandlerTarget(arr []int, target int) [][]int {

	seen := map[int]int{}
	solution := make([][]int,1 ) 

	for i, num := range arr {
		if index,found := seen[target-num]; found && target-num+num != 0 {
			solution = append(solution, []int{index, i})
		}
		seen[num] = i
	}
	return solution
}