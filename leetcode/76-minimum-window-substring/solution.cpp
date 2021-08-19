class Solution {
public:
    string minWindow(string s, string t) {
        map<char, int> tar_map;
        map<char, int> sam_map;
        for (const char c : t) {
            ++tar_map[c];
        }
        queue<int> q;
        int n = s.size();
        int tar_cnt = t.size();
        int sam_cnt = 0;
        int ans_len = n;
        int ans_beg = -1;
        for (int i = 0; i < n; ++i) {
            if (tar_map[s[i]]) {
                if (sam_map[s[i]] < tar_map[s[i]]) {
                    ++sam_cnt;
                }
                ++sam_map[s[i]];
                q.push(i);
                if (sam_cnt == tar_cnt) {
                    int j = q.front();
                    while (sam_map[s[j]] > tar_map[s[j]]) {
                        --sam_map[s[j]];
                        q.pop();
                        j = q.front();
                    }
                    if (i - j + 1 <= ans_len) {
                        ans_len = i - j + 1;
                        ans_beg = j;
                    }
                    --sam_map[s[j]];
                    --sam_cnt;
                    q.pop();
                    if (!q.empty()) {
                        j = q.front();
                        while (sam_map[s[j]] > tar_map[s[j]]) {
                            --sam_map[s[j]];
                            q.pop();
                            j = q.front();
                        }
                    }
                }
            }
        }
        return (ans_beg >= 0) ? s.substr(ans_beg, ans_len) : "";
    }

    int main(int argc, const char *argv[]) {
        string s, t;
        cout << "1 " << minWindow("a", "") << endl;
        cout << "2 " << minWindow("", "a") << endl;
        cout << "3 " << minWindow("", "") << endl;
        while (cin >> s >> t) {
            cout << minWindow(s, t) << endl;
        }
        return 0;
    }
};