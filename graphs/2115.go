package graphs

import "container/list"

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	ingredientsToRecipes := buildIngredientsToRecipesMap(recipes, ingredients)
	visited := make(map[string]bool)
	queue := list.New()
	for i := range supplies {
		queue.PushBack(supplies[i])
	}

	for queue.Len() > 0 {
		e := queue.Front()
		queue.Remove(e)
		ingredient := e.Value.(string)
		if visited[ingredient] {
			continue
		}
		visited[ingredient] = true
		for _, recipe := range ingredientsToRecipes[ingredient] {
			for i := range ingredients[recipe] {
				if ingredients[recipe][i] == ingredient {
					ingredients[recipe] = append(ingredients[recipe][:i], ingredients[recipe][i+1:]...)
					if len(ingredients[recipe]) == 0 {
						queue.PushBack(recipes[recipe])
					}
					break
				}
			}
		}
	}
	answ := make([]string, 0)
	for _, r := range recipes {
		if visited[r] {
			answ = append(answ, r)
		}
	}
	return answ
}

func buildIngredientsToRecipesMap(recipes []string, ingredients [][]string) map[string][]int {
	m := make(map[string][]int)
	for i := range ingredients {
		for j := range ingredients[i] {
			if m[ingredients[i][j]] == nil {
				m[ingredients[i][j]] = make([]int, 0)
			}
			m[ingredients[i][j]] = append(m[ingredients[i][j]], i)
		}
	}
	return m
}
