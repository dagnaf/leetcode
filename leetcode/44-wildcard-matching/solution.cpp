class Solution {
public:
    bool isMatch(string s, string p) {
        bool dp[s.size()+1][p.size()+1] = {false};
        dp[0][0] = true;
        for (size_t j = 1; j <= p.size(); ++j) {
            dp[0][j] = (dp[0][j-1] && '*' == p[j-1]);
            for (size_t i = 1; i <= s.size(); ++i) {
                if ('*' != p[j-1] && '?' != p[j-1]) {
                    dp[i][j] |= (dp[i-1][j-1] && s[i-1] == p[j-1]);
                } else if ('?'== p[j-1]) {
                    dp[i][j] |= dp[i-1][j-1];
                } else {
                    dp[i][j] |= (dp[i-1][j] || dp[i][j-1]);
                }
                // cout << "= s" << i << " [" << s.substr(0, i) << "] p" << j << " [" << p.substr(0, j) << "] " << (dp[i][j] ? "yes" : "no") << endl;
            }
        }
        cout << "s [" << s << "] p [" << p << "] " << (dp[s.size()][p.size()] ? "yes" : "no") << endl;
        return dp[s.size()][p.size()];
    }

    int main(int argc, const char *argv[]) {
        string s, p;
        while (cin >> s >> p) isMatch(s, p);
    }
};