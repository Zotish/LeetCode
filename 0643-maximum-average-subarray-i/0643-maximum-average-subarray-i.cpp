class Solution {
public:
double findMaxAverage(vector<int>& nums, int k) {
    int total = 0;
    for (int i = 0; i < k; i++) {
        total += nums[i];
    }
    int maxTotal = total;

    int lo = 0;
    for (int hi = k; hi < nums.size(); hi++) {
        total = total - nums[lo] + nums[hi];
        lo++;
        maxTotal = max(maxTotal, total);
    }
    return static_cast<double>(maxTotal) / static_cast<double>(k);
}

int max(int a, int b) {
    return (a > b) ? a : b;
}

};