class Solution {
private:
    vector<vector<int>> ans;
    vector<int> acc;
    void combinationSum(vector<int>::iterator beg, vector<int>::iterator end, int target) {
// copy(acc.begin(), acc.end(), ostream_iterator<int>(cout, ","));
// cout << "|" << (end - beg) << endl;
        if (target == 0) {
            ans.push_back(acc);
            return;
        }
        if (beg == end) return;
        combinationSum(beg + 1, end, target);
        int k = 0;
        while (*beg <= target) {
            target -= *beg;
            acc.push_back(*beg);
            combinationSum(beg + 1, end, target);
            ++k;
        }
        while (k--) { acc.pop_back(); }
    }
public:
    vector<vector<int>> combinationSum(vector<int>& candidates, int target) {
        sort(candidates.begin(), candidates.end());
        ans.clear();
        acc.clear();
        combinationSum(candidates.begin(), candidates.end(), target);
        return ans;
    }

    int main(int argc, const char *argv[]) {
        int n;
        while (cin >> n) {
            int target;
            vector<int> candidates(n);
            for (int i = 0; i < n; ++i) {
                cin >> candidates[i];
            }
            cin >> target;
            auto ans = combinationSum(candidates, target);
            cout << "[" << endl;
            for (auto &v : ans) {
                copy(v.begin(), v.end(), ostream_iterator<int>(cout, " "));
                cout << endl;
            }
            cout << "]" << endl;
        }
        return 0;
    }
};