class Solution {
public:
    int firstMissingPositive(vector<int>& nums) {
        int n = nums.size();
        nums.push_back(0);
        for (int i = n; i >= 1; --i) {
            nums[i] = nums[i - 1];
        }
        for (int i = 1; i <= n; ++i) {
            while (/*nums[i] != i && redundant*/ 1 <= nums[i] && nums[i] <= n && nums[i] != nums[nums[i]]) {
                swap(nums[i], nums[nums[i]]);
            }
        }
        // int ans = 1;
        // for (int i = 1; i <= n; ++i) {
        //     if (1 <= nums[i] && nums[i] <= n) {
        //         if (nums[i] == ans) ++ans;
        //         else break;
        //     }
        // }
        // return ans;
        for (int i = 1; i <= n; ++i) {
            if (nums[i] != i) return i;
        }
        return n + 1;
    }

    int main(int argc, const char *argv[]) {
        return 0;
    }
};