package main

func combinationSum(candidates []int, target int) [][]int {
	list := make([][]int, 0)
	comb := make([]int, 0)

	combinationSumRec(candidates, 0, comb, target, &list)
	return list
}

func combinationSumRec(candidates []int, fromIndex int, comb []int, target int, list *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		final := make([]int, len(comb))
		copy(final, comb)
		*list = append(*list, final)
		return
	}
	for i := fromIndex; i <= len(candidates); i++ {
		comb = append(comb, candidates[i])
		combinationSumRec(candidates, i, comb, target-candidates[i], list)
		comb = comb[:len(comb)-1]
	}
}

/*
class Solution {
	public void rec(List<List<Integer>> res, int[] candidates, int start, int target, List<Integer> comb) {
		if (target < 0) {
			return;
		}
		if (target == 0) {
			res.add(new ArrayList<Integer>(comb));
			return;
		}

		for (int i=start; i<candidates.length; i++) {
			comb.add(candidates[i]);
			rec(res, candidates, i, target - candidates[i], comb);
			comb.remove(comb.size() - 1);
		}
	}

    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        List<List<Integer>> res = new ArrayList<List<Integer>>();
        rec(res, candidates, 0, target, new ArrayList<Integer>());
        return res;
    }
}

*/
