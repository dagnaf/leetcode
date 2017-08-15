class Solution {
public:
    vector<int> findSubstring(string s, vector<string>& words) {
        vector<int> ans;
        int n = s.size();
        int m = words.size();
        if (n == 0 || m == 0) return ans;
        int l = words[0].size();
        unordered_map<string, int> dict;
        for (string &word: words) {
            ++dict[word];
        }
        for (int i = 0; i < l; ++i) {
            unordered_map<string, int> strs;
            int cnt = 0;
            int beg = i;
            for (int j = i; j + l <= n; j += l) {
                string seg = s.substr(j, l);
                int tar_seg_cnt = dict[seg];
                if (tar_seg_cnt == 0) { // not consecutive, reset
                    strs.clear();
                    cnt = 0;
                    beg = j + l;
                    continue;
                }
                int &cur_seg_cnt = strs[seg];
                if (cur_seg_cnt + 1 <= tar_seg_cnt) {
                    ++cnt;
                    ++cur_seg_cnt;
                } else {
                    while (true) {
                        string seg_first = s.substr(beg, l);
                        beg += l;
                        int &cur_seg_first_cnt = strs[seg_first];
                        --cnt;
                        --cur_seg_first_cnt;
                        if (seg == seg_first) {
                            ++cnt;
                            ++cur_seg_first_cnt;
                            break;
                        }
                    }
                }
                if (cnt == m) {
                    ans.push_back(beg);
                }
            }
        }
        return ans;
    }

    int main(int argc, const char *argv[]) {
        string s;
        int n;
        while (cin >> s >> n) {
            vector<string> words(n);
            for (int i = 0; i < n; ++i) {
                cin >> words[i];
            }
            auto ans = findSubstring(s, words);
            for (int i : ans) {
                cout << i << " " << s.substr(i, n * words[0].size()) << endl;
            }
        }
        return 0;
    }
};