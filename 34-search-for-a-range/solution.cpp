class Solution {
public:
    template<class T>
    int cmp_bound(const vector<int> &v, int target) {
        static T cmp;
        int lo = 0;
        int hi = v.size()/* - 1*/;
        while (lo < hi) {
            int mid = (lo + hi) / 2;
            if (/*v[mid] >= target*/ cmp(v[mid], target) ) {
                hi = mid;
            } else {
                lo = mid + 1;
            }
            // cout << "[ " << lo << " " << hi << "]" << endl;
        }
        return lo;
    }
    int lower_bound(const vector<int> &v, int target) {
        return cmp_bound<greater_equal<int>>(v, target);
    }
    int upper_bound(const vector<int> &v, int target) {
        return cmp_bound<greater<int>>(v, target);
    }
    vector<int> searchRange(vector<int>& nums, int target) {
        // STL
        // auto it_lb = lower_bound(nums.begin(), nums.end(), target);
        // if (it_lb == nums.end() || *it_lb != target) return {-1, -1};
        // auto it_ub = upper_bound(nums.begin(), nums.end(), target);
        // return {it_lb - nums.begin(), it_ub - nums.begin() - 1};
        int n = nums.size();
        if (n == 0) return {-1, -1};
        int lb = lower_bound(nums, target);
        if (lb == n || nums[lb] != target) return {-1, -1};
        int ub = upper_bound(nums, target);
        // return {lb, ub};
        return {lb, ub - 1};
    }

    int main(int argc, const char *argv[]) {
        int target, n;
        while (cin >> target >> n) {
            vector<int> nums(n);
            for (int i = 0; i < n; ++i) cin >> nums[i];
            auto ret = searchRange(nums, target);
            cout << ret[0] << " " << ret[1] << endl;
        }
        return 0;
    }
};