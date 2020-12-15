class Solution {
    public int[] runningSum(int[] nums) {
        int[] res = new int[nums.length];
        int counter = 1;
        int total = nums[0];
        res[0] = total;
        for (int i = 1; i < nums.length; i++) {

            res[i] = total+nums[i];
            total = res[i];

        }
        return res;

    }
}