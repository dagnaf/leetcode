class Solution {
private:
    void dfs(vector<vector<int>> &ans, vector<int> &cur, vector<int> &a, int target, int i) {
        if (target < 0) return;
        if (target == 0) {
            ans.push_back(cur);
            return;
        }
        int n = a.size();
        for (int j = i; j < n; ++j) {
            if (j && j != i && a[j] == a[j - 1]) continue;
            cur.push_back(a[j]);
            dfs(ans, cur, a, target - a[j], j + 1);
            cur.pop_back();
        }
    }
public:
    vector<vector<int>> combinationSum2(vector<int>& candidates, int target) {
        vector<vector<int>> ans;
        vector<int> cur;
        sort(candidates.begin(), candidates.end());
        dfs(ans, cur, candidates, target, 0);
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
            auto ans = combinationSum2(candidates, target);
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