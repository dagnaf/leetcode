class Solution {
public:
    int longestValidParentheses(string s) {
        // TLE
        // int n = s.size();
        // vector<vector<bool>> dp(n, vector<bool>(n, false));
        // int ans = 0;
        // for (int k = 2; k <= n; k += 2) {
        //     for (int i = 0; i + k <= n; ++i) {
        //         int j = i + k - 1;
        //         if ((i + 1 == j || dp[i + 1][j - 1]) && '(' == s[i] && ')' == s[j])
        //             dp[i][j] = true;
        //         for (int m = i + 1; m <= j - 2; m += 2) {
        //             if (dp[i][m] && dp[m + 1][j])
        //                 dp[i][j] = true;
        //         }
        //         if (dp[i][j]) ans = max(ans, j - i + 1);
        //     }
        // }
        // return ans;
        int n = s.size();
        stack<int> stk;
        int ans = 0;
        vector<int> p(n, -1);
        for (int i = 0; i < n; ++i) {
            if (stk.empty() || s[i] == '(') {
                stk.push(i);
            } else if ('(' == s[stk.top()] && ')' == s[i]) {
                p[stk.top()] = i;
                stk.pop();
            }
        }
        // for (int i = 0; i < n; ++i) {
        //     cout << "-" << i << " " << (p[i] == -1 ? "" : s.substr(i, p[i] - i + 1)) << endl;
        // }
        for (int i = 0; i < n;) {
            if (p[i] == -1) {
                ++i;
                continue;
            }
            int j = p[i];
            while (j + 1 < n && p[j + 1] != -1) {
                j = p[j + 1];
            }
            ans = max(ans, j - i + 1);
            i = j + 1;
        }
        return ans;
    }

    int main(int argc, const char *argv[]) {
        string s;
        while (cin >> s) {
            cout << longestValidParentheses(s) << endl;
        }
        return 0;
    }
};