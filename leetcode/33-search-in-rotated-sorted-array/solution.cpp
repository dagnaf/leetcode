class Solution {
public:
    int search(vector<int>& nums, int target) {
        int lo = 0;
        int hi = nums.size() - 1;
        bool tar_in_front = target >= nums[0];
        while (lo < hi) {
            int mid = (lo + hi) / 2;
            bool mid_in_front = nums[mid] >= nums[0];
            if ((tar_in_front && mid_in_front) ||
                (!tar_in_front && !mid_in_front)) {
                if (target > nums[mid]) {
                    lo = mid + 1;
                } else {
                    hi = mid;
                }
            } else {
                if (tar_in_front) hi = mid - 1;
                else lo = mid + 1;
            }
            // cout << "lo " << lo << " hi " << hi << endl;
        }
        return nums[lo] == target ? lo : -1;
    }

    int main(int argc, const char *argv[]) {
        int target, n;
        while (cin >> target >> n) {
            vector<int> nums(n);
            for (int i = 0; i < n; ++i) cin >> nums[i];
            cout << search(nums, target) << endl;
        }
        return 0;
    }
};