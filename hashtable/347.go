package hashtable

func topKFrequent(nums []int, k int) []int {
	valueCounts := make(map[int]int)
	for _, val := range nums {
		valueCounts[val]++
	}
	countToVal := make(map[int]map[int]bool)
	for key, val := range valueCounts {
		if _, ok := countToVal[val]; !ok {
			countToVal[val] = make(map[int]bool)
		}
		countToVal[val][key] = true
	}
	kFrequentValues := make([]int, 0)
	for i := len(nums); i > 0; i-- {
		if countToVal[i] == nil {
			continue
		}
		values := countToVal[i]
		for val := range values {
			kFrequentValues = append(kFrequentValues, val)
			if len(kFrequentValues) == k {
				return kFrequentValues
			}
		}
	}
	return kFrequentValues
}

/*

    public static List<Integer> topKFrequent(int[] nums, int k) {
    	List<Integer> res = new ArrayList<Integer>();
    	Map<Integer, Integer> map = new HashMap<Integer, Integer>(nums.length);
    	for (int i = 0; i < nums.length; i++) {
			map.putIfAbsent(nums[i], 0);
			map.put(nums[i], map.get(nums[i]) + 1);
		}

    	Map<Integer, Set<Integer>> numbers = new HashMap<Integer, Set<Integer>>();
    	for (int key : map.keySet()) {
    		int val = map.get(key);
			numbers.putIfAbsent(val, new HashSet<Integer>());
			numbers.get(val).add(key);
		}

    	for (int i=nums.length; i>0; i--) {
    		if (!numbers.containsKey(i)) {
    			continue;
    		}
    		Set<Integer> s = numbers.get(i);
    		for (Integer integer : s) {
    			res.add(integer);
    			if (res.size() == k) {
    				return res;
    			}
    		}
    	}

    	return res;
    }


*/
